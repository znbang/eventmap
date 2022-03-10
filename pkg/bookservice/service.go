package bookservice

import (
	"errors"
	"github.com/microcosm-cc/bluemonday"
	"github.com/znbang/eventmap/pkg/uuid"
)

type BookService struct {
	BookRepository
	ChapterRepository
	BookJobRepository
}

func NewBookService(bookRepository BookRepository, bookJobRepository BookJobRepository, chapterRepository ChapterRepository) *BookService {
	return &BookService{
		BookRepository:    bookRepository,
		BookJobRepository: bookJobRepository,
		ChapterRepository: chapterRepository,
	}
}

func (s *BookService) GetBooksByUser(books *BookList, userId string, q string, page int, size int) error {
	var (
		total int
		items []Book
	)

	if err := s.BookRepository.CountByUserId(&total, userId, q); err != nil {
		return err
	}

	if err := s.BookRepository.FindByUserId(&items, userId, q, page, size); err != nil {
		return err
	}

	books.Page = page
	books.Size = size
	books.Total = (total + size - 1) / size
	books.Items = items

	return nil
}

func (s *BookService) GetChapterTitles(chapters *ChapterList, bookId string, page int, size int) error {
	var (
		total int
		items []Chapter
	)

	if err := s.ChapterRepository.CountByBookId(&total, bookId); err != nil {
		return err
	}

	if err := s.ChapterRepository.FindByBookId(&items, bookId, page, size); err != nil {
		return err
	}

	chapters.Page = page
	chapters.Size = size
	chapters.Total = (total + size - 1) / size
	chapters.Items = items

	return nil
}

func (s *BookService) GetAllChapters(chapters *[]Chapter, bookId string) error {
	return s.ChapterRepository.FindAllByBookId(chapters, bookId)
}

func (s *BookService) CreateBook(userID string, input *Book) error {
	p := bluemonday.UGCPolicy()
	input.Title = p.Sanitize(input.Title)
	input.Author = p.Sanitize(input.Author)
	input.ID = uuid.ULID()
	input.UserID = userID

	return s.BookRepository.Create(input)
}

func (s *BookService) UpdateBook(userId string, input *Book) error {
	var book Book

	if err := s.BookRepository.FindByIdAndUserId(&book, input.ID, userId); err != nil {
		return err
	}

	p := bluemonday.UGCPolicy()
	book.Title = p.Sanitize(input.Title)
	book.Author = p.Sanitize(input.Author)
	book.URL = input.URL

	return s.BookRepository.Save(&book)
}

func (s *BookService) DeleteBook(userId, bookId string) error {
	var book Book

	if err := s.BookRepository.FindByIdAndUserId(&book, bookId, userId); errors.Is(err, ErrNoSuchBook) {
		return errors.New("book: current user is not the owner of the book")
	}

	return s.BookRepository.Delete(&book)
}

func (s *BookService) SyncBook(book Book) error {
	job := BookJob{}
	if err := s.BookJobRepository.FindByBookId(&job, book.ID); errors.Is(err, ErrNoSuchBookJob) {
		job = BookJob{
			ID:     uuid.ULID(),
			BookID: book.ID,
			Status: StatusPending,
		}

		return s.BookJobRepository.Create(&job)
	} else if err != nil {
		return err
	}

	if job.IsDone() {
		job.Status = StatusPending
		job.Message = ""
		return s.BookJobRepository.Save(&job)
	}

	return nil
}

func (s *BookService) StopSyncBook(bookId string) error {
	var bookJob BookJob

	if err := s.BookJobRepository.FindByBookId(&bookJob, bookId); err != nil {
		return err
	}

	bookJob.Status = StatusDone

	return s.BookJobRepository.Save(&bookJob)
}

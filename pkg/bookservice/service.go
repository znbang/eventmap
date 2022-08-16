package bookservice

import (
	"errors"

	"github.com/microcosm-cc/bluemonday"
	"github.com/znbang/eventmap/pkg/uuid"
)

type BookService struct {
	bookRepository BookRepository
	chapterRepository ChapterRepository
	bookJobRepository BookJobRepository
}

func NewBookService(bookRepository BookRepository, bookJobRepository BookJobRepository, chapterRepository ChapterRepository) *BookService {
	return &BookService{
		bookRepository:    bookRepository,
		bookJobRepository: bookJobRepository,
		chapterRepository: chapterRepository,
	}
}

func (s *BookService) GetBooksByUser(books *BookList, userId string, q string, page int, size int) error {
	var (
		total int
		items []Book
	)

	if err := s.bookRepository.CountByUserId(&total, userId, q); err != nil {
		return err
	}

	if err := s.bookRepository.FindByUserId(&items, userId, q, page, size); err != nil {
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

	if err := s.chapterRepository.CountByBookId(&total, bookId); err != nil {
		return err
	}

	if err := s.chapterRepository.FindByBookId(&items, bookId, page, size); err != nil {
		return err
	}

	chapters.Page = page
	chapters.Size = size
	chapters.Total = (total + size - 1) / size
	chapters.Items = items

	return nil
}

func (s *BookService) GetAllChapters(chapters *[]Chapter, bookId string) error {
	return s.chapterRepository.FindAllByBookId(chapters, bookId)
}

func (s *BookService) CreateBook(userID string, input *Book) error {
	p := bluemonday.UGCPolicy()
	input.Title = p.Sanitize(input.Title)
	input.Author = p.Sanitize(input.Author)
	input.ID = uuid.ULID()
	input.UserID = userID

	return s.bookRepository.Create(input)
}

func (s *BookService) UpdateBook(userId string, input *Book) error {
	var book Book

	if err := s.bookRepository.FindByIdAndUserId(&book, input.ID, userId); err != nil {
		return err
	}

	p := bluemonday.UGCPolicy()
	book.Title = p.Sanitize(input.Title)
	book.Author = p.Sanitize(input.Author)
	book.URL = input.URL

	return s.bookRepository.Save(&book)
}

func (s *BookService) DeleteBook(userId, bookId string) error {
	var book Book

	if err := s.bookRepository.FindByIdAndUserId(&book, bookId, userId); errors.Is(err, ErrNoSuchBook) {
		return errors.New("book: current user is not the owner of the book")
	}

	return s.bookRepository.Delete(&book)
}

func (s *BookService) SyncBook(book Book) error {
	job := BookJob{}
	if err := s.bookJobRepository.FindByBookId(&job, book.ID); errors.Is(err, ErrNoSuchBookJob) {
		job = BookJob{
			ID:     uuid.ULID(),
			BookID: book.ID,
			Status: StatusPending,
		}

		return s.bookJobRepository.Create(&job)
	} else if err != nil {
		return err
	}

	if job.IsDone() {
		job.Status = StatusPending
		job.Message = ""
		return s.bookJobRepository.Save(&job)
	}

	return nil
}

func (s *BookService) StopSyncBook(bookId string) error {
	var bookJob BookJob

	if err := s.bookJobRepository.FindByBookId(&bookJob, bookId); err != nil {
		return err
	}

	bookJob.Status = StatusDone

	return s.bookJobRepository.Save(&bookJob)
}

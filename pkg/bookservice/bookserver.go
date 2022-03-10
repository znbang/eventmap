package bookservice

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/bufbuild/connect-go"
	v1 "github.com/znbang/eventmap/gen/book/v1"
	"github.com/znbang/eventmap/internal/mvc"
	"github.com/znbang/eventmap/pkg/auth"
	"github.com/znbang/eventmap/pkg/crawlerservice"
	"github.com/znbang/eventmap/pkg/userservice"
	"github.com/znbang/validation"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookServer struct {
	*BookService
}

func (s *BookServer) CreateBook(ctx context.Context, r *connect.Request[v1.CreateBookRequest]) (*connect.Response[v1.CreateBookResponse], error) {
	var (
		user     userservice.User
		validate = validation.New()
	)

	if err := auth.CurrentUser(ctx, &user); err != nil {
		return nil, err
	}

	book := convertPBToBook(r.Msg.Book)

	if err := validateBook(s.BookRepository, validate, user, &book); err != nil {
		return nil, err
	}

	if validate.HasErrors() {
		return nil, errors.New(fmt.Sprintf("%v", validate.Errors))
	}

	if err := s.BookService.CreateBook(user.ID, &book); err != nil {
		return nil, err
	}

	if err := s.BookService.SyncBook(book); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.CreateBookResponse{
		Id: book.ID,
	}), nil
}

func (s *BookServer) UpdateBook(ctx context.Context, r *connect.Request[v1.UpdateBookRequest]) (*connect.Response[v1.UpdateBookResponse], error) {
	var (
		user     userservice.User
		validate = validation.New()
	)

	if err := auth.CurrentUser(ctx, &user); err != nil {
		return nil, err
	}

	book := convertPBToBook(r.Msg.Book)

	if err := validateBook(s.BookRepository, validate, user, &book); err != nil {
		return nil, err
	}

	if validate.HasErrors() {
		return nil, errors.New(fmt.Sprintf("%v", validate.Errors))
	}

	if err := s.BookService.UpdateBook(user.ID, &book); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.UpdateBookResponse{
		Id: book.ID,
	}), nil
}

func (s *BookServer) DeleteBook(ctx context.Context, r *connect.Request[v1.DeleteBookRequest]) (*connect.Response[v1.DeleteBookResponse], error) {
	var user userservice.User

	if err := auth.CurrentUser(ctx, &user); err != nil {
		return nil, err
	}

	if err := s.BookService.DeleteBook(user.ID, r.Msg.Id); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.DeleteBookResponse{}), nil
}

func (s *BookServer) DeleteChapter(ctx context.Context, r *connect.Request[v1.DeleteChapterRequest]) (*connect.Response[v1.DeleteChapterResponse], error) {
	var (
		book Book
		user userservice.User
	)

	if err := auth.CurrentUser(ctx, &user); err != nil {
		return nil, err
	}

	if err := s.BookRepository.FindById(&book, r.Msg.Id); err != nil {
		return nil, err
	}

	if user.ID != book.UserID {
		return nil, errors.New("user is not book's owner")
	}

	if err := s.ChapterRepository.DeleteByPage(r.Msg.Id, int(r.Msg.Page)); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.DeleteChapterResponse{}), nil
}

func (s *BookServer) GetBook(ctx context.Context, r *connect.Request[v1.GetBookRequest]) (*connect.Response[v1.GetBookResponse], error) {
	var book Book

	if err := s.BookRepository.FindById(&book, r.Msg.Id); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.GetBookResponse{
		Book: convertBookToPB(book),
	}), nil
}

func (s *BookServer) GetToc(ctx context.Context, r *connect.Request[v1.GetTocRequest]) (*connect.Response[v1.GetTocResponse], error) {
	var (
		book       Book
		chapters   ChapterList
		bookId     = r.Msg.Id
		page, size = mvc.PageSize(r.Msg.Page, r.Msg.Size)
	)

	if err := s.BookRepository.FindById(&book, bookId); err != nil {
		return nil, err
	}

	if err := s.BookService.GetChapterTitles(&chapters, bookId, page, size); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.GetTocResponse{
		Book:  convertBookToPB(book),
		Page:  int32(chapters.Page),
		Size:  int32(chapters.Size),
		Total: int32(chapters.Total),
		Items: convertChapterItems(chapters.Items),
	}), nil
}

func (s *BookServer) GetChapter(ctx context.Context, r *connect.Request[v1.GetChapterRequest]) (*connect.Response[v1.GetChapterResponse], error) {
	var (
		book    Book
		chapter Chapter
		total   int
		page, _ = mvc.PageSize(r.Msg.Page, 0)
		bookId  = r.Msg.Id
	)

	if err := s.BookRepository.FindById(&book, bookId); err != nil {
		return nil, err
	}

	if err := s.ChapterRepository.FindByBookIdAndPage(&chapter, bookId, page); err != nil {
		return nil, err
	}

	if err := s.ChapterRepository.CountByBookId(&total, bookId); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.GetChapterResponse{
		Book:    convertBookToPB(book),
		Chapter: convertChapterToPB(chapter),
		Page:    int32(page),
		Total:   int32(total),
	}), nil
}

func (s *BookServer) ListBook(ctx context.Context, r *connect.Request[v1.ListBookRequest]) (*connect.Response[v1.ListBookResponse], error) {
	var (
		user       userservice.User
		books      BookList
		page, size = mvc.PageSize(r.Msg.Page, r.Msg.Size)
	)

	if err := auth.CurrentUser(ctx, &user); err != nil {
		return nil, err
	}

	if err := s.BookService.GetBooksByUser(&books, user.ID, r.Msg.Filter, page, size); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.ListBookResponse{
		Page:  int32(books.Page),
		Size:  int32(books.Size),
		Total: int32(books.Total),
		Items: convertBookItems(books.Items),
	}), nil
}

func (s *BookServer) SyncNew(ctx context.Context, r *connect.Request[v1.SyncNewRequest]) (*connect.Response[v1.SyncNewResponse], error) {
	var (
		user   userservice.User
		book   Book
		bookId = r.Msg.Id
	)

	if err := auth.CurrentUser(ctx, &user); err != nil {
		return nil, err
	}

	if err := s.BookRepository.FindByIdAndUserId(&book, bookId, user.ID); err != nil {
		return nil, err
	}

	if err := s.BookService.SyncBook(book); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.SyncNewResponse{}), nil
}

func (s *BookServer) SyncAll(ctx context.Context, r *connect.Request[v1.SyncAllRequest]) (*connect.Response[v1.SyncAllResponse], error) {
	var (
		user   userservice.User
		book   Book
		bookId = r.Msg.Id
	)

	if err := auth.CurrentUser(ctx, &user); err != nil {
		return nil, err
	}

	if err := s.BookRepository.FindByIdAndUserId(&book, bookId, user.ID); err != nil {
		return nil, err
	}

	if err := s.ChapterRepository.DeleteByBookId(bookId); err != nil {
		return nil, err
	}

	if err := s.BookService.SyncBook(book); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.SyncAllResponse{}), nil
}

func (s *BookServer) StopSync(ctx context.Context, r *connect.Request[v1.StopSyncRequest]) (*connect.Response[v1.StopSyncResponse], error) {
	var (
		user   userservice.User
		bookId = r.Msg.Id
	)

	if err := auth.CurrentUser(ctx, &user); err != nil {
		return nil, err
	}

	if err := s.BookService.StopSyncBook(bookId); err != nil {
		return nil, err
	}

	return connect.NewResponse(&v1.StopSyncResponse{}), nil
}

func (s *BookServer) SyncStatus(ctx context.Context, r *connect.Request[v1.SyncStatusRequest], stream *connect.ServerStream[v1.SyncStatusResponse]) error {
	ch := EventBus.Register()
	defer EventBus.Unregister(ch)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case event := <-ch:
			if err := stream.Send(&v1.SyncStatusResponse{
				Id:      event.ID,
				Message: event.Message,
				Status:  v1.BookJobStatus(event.Status),
			}); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *BookServer) DownloadBook(ctx context.Context, r *connect.Request[v1.DownloadBookRequest]) (*connect.Response[v1.DownloadBookResponse], error) {
	var (
		book     Book
		chapters []Chapter
		bookId   = r.Msg.Id
		sb       strings.Builder
	)

	if err := s.BookRepository.FindById(&book, bookId); err != nil {
		return nil, err
	}

	if err := s.BookService.GetAllChapters(&chapters, bookId); err != nil {
		return nil, err
	}

	for _, chapter := range chapters {
		texts := [...]string{chapter.Title, "\n\n", chapter.Body, "\n\n"}
		for _, text := range texts {
			sb.WriteString(text)
		}
	}

	return connect.NewResponse(&v1.DownloadBookResponse{
		Title:   book.Title,
		Content: sb.String(),
	}), nil
}

func convertBookItems(src []Book) []*v1.Book {
	items := make([]*v1.Book, 0, len(src))
	for _, item := range src {
		items = append(items, convertBookToPB(item))
	}
	return items
}

func convertBookToPB(src Book) *v1.Book {
	return &v1.Book{
		Id:        src.ID,
		UserId:    src.UserID,
		Title:     src.Title,
		Author:    src.Author,
		Url:       src.URL,
		Job:       convertBookJobToPB(src.Job),
		CreatedAt: timestamppb.New(src.CreatedAt),
		UpdatedAt: timestamppb.New(src.UpdatedAt),
	}
}

func convertBookJobToPB(src *BookJob) *v1.BookJob {
	if src == nil {
		return nil
	}

	return &v1.BookJob{
		Id:        src.ID,
		Status:    v1.BookJobStatus(src.Status),
		BookId:    src.BookID,
		Message:   src.Message,
		CreatedAt: timestamppb.New(src.CreatedAt),
	}
}

func convertPBToBook(src *v1.Book) Book {
	return Book{
		ID:     src.Id,
		Title:  src.Title,
		Author: src.Author,
		URL:    src.Url,
	}
}

func convertChapterItems(src []Chapter) []*v1.Chapter {
	items := make([]*v1.Chapter, 0, len(src))
	for _, item := range src {
		items = append(items, convertChapterToPB(item))
	}
	return items
}

func convertChapterToPB(src Chapter) *v1.Chapter {
	return &v1.Chapter{
		Id:        src.ID,
		BookId:    src.BookID,
		Page:      int32(src.Page),
		Url:       src.URL,
		Title:     src.Title,
		Body:      src.Body,
		CreatedAt: timestamppb.New(src.CreatedAt),
		UpdatedAt: timestamppb.New(src.UpdatedAt),
	}
}

func validateBook(bookRepository BookRepository, validate *validation.Validation, user userservice.User, input *Book) error {
	input.Title = strings.TrimSpace(input.Title)
	input.Author = strings.TrimSpace(input.Author)
	input.URL = strings.TrimSpace(input.URL)

	validate.Required(input.Title, "title", "books.required.title")
	validate.Required(input.Author, "author", "books.required.author")
	if validate.Required(input.URL, "url", "books.required.url") {
		validate.URL(input.URL, "url", "books.invalid.url")
	}

	errs := validate.Errors

	if errs["title"] == "" && errs["author"] == "" {
		var book Book
		if err := bookRepository.FindByUserIdAndTitleAndAuthor(&book, user.ID, input.Title, input.Author); errors.Is(err, ErrNoSuchBook) {
		} else if err != nil {
			return err
		} else if book.ID != input.ID {
			errs["title"] = "books.exists.title"
			errs["author"] = "books.exists.author"
		}
	}

	if errs["url"] == "" {
		var book Book
		if err := bookRepository.FindByUserIdAndUrl(&book, user.ID, input.URL); errors.Is(err, ErrNoSuchBook) {
		} else if err != nil {
			return err
		} else if book.ID != input.ID {
			errs["url"] = "books.exists.url"
		}
	}

	if errs["url"] == "" {
		crawler := crawlerservice.New()
		if !crawler.Supports(input.URL) {
			errs["url"] = "books.invalid.url"
		}
	}

	return nil
}

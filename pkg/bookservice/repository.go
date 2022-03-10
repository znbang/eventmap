package bookservice

type BookRepository interface {
	CountByUserId(total *int, userId string, q string) error
	FindByUserId(books *[]Book, userId string, q string, page int, size int) error
	FindByUserIdAndTitleAndAuthor(book *Book, userId string, title string, author string) error
	FindByUserIdAndUrl(book *Book, userId string, url string) error
	FindById(book *Book, id string) error
	FindByIdAndUserId(book *Book, id string, userId string) error
	Create(book *Book) error
	Save(book *Book) error
	Delete(book *Book) error
}

type BookJobRepository interface {
	FindById(bookJob *BookJob, id string) error
	FindByBookId(bookJob *BookJob, bookId string) error
	Find100Pending(bookJobs *[]BookJob) error
	Create(bookJob *BookJob) error
	Save(bookJob *BookJob) error
}

type ChapterRepository interface {
	CountByBookId(total *int, bookId string) error
	FindByBookId(chapters *[]Chapter, bookId string, page int, size int) error
	FindByBookIdAndPage(chapter *Chapter, bookId string, page int) error
	FindAllByBookId(chapters *[]Chapter, bookId string) error
	FindLastChapterByBookId(chapter *Chapter, bookId string) error
	Create(chapter *Chapter) error
	DeleteByBookId(bookId string) error
	DeleteByPage(bookId string, page int) error
}

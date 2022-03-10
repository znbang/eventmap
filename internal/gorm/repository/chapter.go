package repository

import (
	"errors"
	"github.com/znbang/eventmap/internal/dbx"
	"github.com/znbang/eventmap/pkg/bookservice"
	"gorm.io/gorm"
)

type chapterRepository struct {
	*dbx.DB
}

func NewChapterRepository(db *dbx.DB) bookservice.ChapterRepository {
	return &chapterRepository{
		DB: db,
	}
}

func (r chapterRepository) CountByBookId(total *int, bookId string) error {
	var total64 int64

	if err := r.DB.Model(&bookservice.Chapter{}).Where("book_id=?", bookId).Count(&total64).Error; err != nil {
		return err
	}

	*total = int(total64)

	return nil
}

func (r chapterRepository) FindByBookId(chapters *[]bookservice.Chapter, bookId string, page int, size int) error {
	offset := size * (page - 1)
	return r.DB.Select("id", "title", "page").Where("book_id=?", bookId).Order("page").Limit(size).Offset(offset).Find(chapters).Error
}

func (r chapterRepository) FindByBookIdAndPage(chapter *bookservice.Chapter, bookId string, page int) error {
	if err := r.DB.Where("book_id=? and page=?", bookId, page).First(chapter).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return bookservice.ErrNoSuchChapter
	} else {
		return err
	}
}

func (r chapterRepository) FindAllByBookId(chapters *[]bookservice.Chapter, bookId string) error {
	return r.DB.Where("book_id=?", bookId).Order("page").Find(chapters).Error
}

func (r chapterRepository) FindLastChapterByBookId(chapter *bookservice.Chapter, bookId string) error {
	if err := r.DB.Where("book_id=?", bookId).Order("page desc").First(chapter).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return bookservice.ErrNoSuchChapter
	} else {
		return err
	}
}

func (r chapterRepository) Create(chapter *bookservice.Chapter) error {
	return r.DB.Create(chapter).Error
}

func (r chapterRepository) DeleteByBookId(bookId string) error {
	return r.DB.Where("book_id=?", bookId).Delete(&bookservice.Chapter{}).Error
}

func (r chapterRepository) DeleteByPage(bookId string, page int) error {
	return r.DB.Where("book_id=? and page>=?", bookId, page).Delete(&bookservice.Chapter{}).Error
}

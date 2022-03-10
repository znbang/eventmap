package repository

import (
	"errors"
	"github.com/znbang/eventmap/internal/dbx"
	"github.com/znbang/eventmap/pkg/bookservice"
	"gorm.io/gorm"
)

type bookRepository struct {
	*dbx.DB
}

func NewBookRepository(db *dbx.DB) bookservice.BookRepository {
	return &bookRepository{
		DB: db,
	}
}

func (r bookRepository) CountByUserId(total *int, userId string, q string) error {
	var total64 int64

	if q == "" {
		if err := r.DB.Model(&bookservice.Book{}).Where("user_id=?", userId).Count(&total64).Error; err != nil {
			return err
		}
	} else {
		term := "%" + q + "%"
		if err := r.DB.Model(&bookservice.Book{}).Where("user_id=? and (author like ? or title like ?)", userId, term, term).Count(&total64).Error; err != nil {
			return err
		}
	}

	*total = int(total64)

	return nil
}

func (r bookRepository) FindByUserId(books *[]bookservice.Book, userId string, q string, page int, size int) error {
	offset := size * (page - 1)

	if q == "" {
		return r.DB.Joins("Job").Where("user_id=?", userId).Order("updated_at desc").Offset(offset).Limit(size).Find(books).Error
	} else {
		term := "%" + q + "%"
		return r.DB.Joins("Job").Where("user_id=? and (author like ? or title like ?)", userId, term, term).Order("updated_at desc").Offset(offset).Limit(size).Find(books).Error
	}
}

func (r bookRepository) FindByUserIdAndTitleAndAuthor(book *bookservice.Book, userId string, title string, author string) error {
	if err := r.DB.Where("user_id=? and title=? and author=?", userId, title, author).First(book).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return bookservice.ErrNoSuchBook
	} else {
		return err
	}
}

func (r bookRepository) FindByUserIdAndUrl(book *bookservice.Book, userId string, url string) error {
	if err := r.DB.Where("user_id=? and url=?", userId, url).First(book).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return bookservice.ErrNoSuchBook
	} else {
		return err
	}
}

func (r bookRepository) FindById(book *bookservice.Book, id string) error {
	if err := r.DB.Where("id=?", id).First(book).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return bookservice.ErrNoSuchBook
	} else {
		return err
	}
}

func (r bookRepository) FindByIdAndUserId(book *bookservice.Book, id string, userId string) error {
	if err := r.DB.Where("id=? and user_id=?", id, userId).First(book).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return bookservice.ErrNoSuchBook
	} else {
		return err
	}
}

func (r bookRepository) Create(book *bookservice.Book) error {
	return r.DB.Create(book).Error
}

func (r bookRepository) Save(book *bookservice.Book) error {
	return r.DB.Save(book).Error
}

func (r bookRepository) Delete(book *bookservice.Book) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := r.DB.Where("book_id=?", book.ID).Delete(&bookservice.BookJob{}).Error; err != nil {
			return err
		}
		if err := r.DB.Where("book_id=?", book.ID).Delete(&bookservice.Chapter{}).Error; err != nil {
			return err
		}
		return tx.Delete(book).Error
	})
}

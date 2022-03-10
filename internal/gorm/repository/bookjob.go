package repository

import (
	"errors"
	"github.com/znbang/eventmap/internal/dbx"
	"github.com/znbang/eventmap/pkg/bookservice"
	"gorm.io/gorm"
)

type bookJobRepository struct {
	*dbx.DB
}

func NewBookJobRepository(db *dbx.DB) bookservice.BookJobRepository {
	return &bookJobRepository{
		DB: db,
	}
}

func (r bookJobRepository) FindById(bookJob *bookservice.BookJob, id string) error {
	if err := r.DB.Where("id=?", id).First(bookJob).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return bookservice.ErrNoSuchBookJob
	} else {
		return err
	}
}

func (r bookJobRepository) FindByBookId(bookJob *bookservice.BookJob, bookId string) error {
	if err := r.DB.Where("book_id=?", bookId).First(bookJob).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return bookservice.ErrNoSuchBookJob
	} else {
		return err
	}
}

func (r bookJobRepository) Find100Pending(bookJobs *[]bookservice.BookJob) error {
	return r.DB.Where("status<>?", bookservice.StatusDone).Limit(100).Find(bookJobs).Error
}

func (r bookJobRepository) Create(bookJob *bookservice.BookJob) error {
	return r.DB.Create(bookJob).Error
}

func (r bookJobRepository) Save(bookJob *bookservice.BookJob) error {
	return r.DB.Save(bookJob).Error
}

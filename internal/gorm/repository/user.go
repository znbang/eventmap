package repository

import (
	"errors"
	"github.com/znbang/eventmap/internal/dbx"
	"github.com/znbang/eventmap/pkg/userservice"
	"gorm.io/gorm"
)

type userRepository struct {
	*dbx.DB
}

func NewUserRepository(db *dbx.DB) userservice.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r userRepository) FindById(user *userservice.User, id string) error {
	if err := r.DB.Where("id=?", id).First(user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return userservice.ErrNoSuchUser
	} else {
		return err
	}
}

func (r userRepository) FindByUid(user *userservice.User, uid string, provider int) error {
	if err := r.DB.Where("uid=? and provider=?", uid, provider).First(user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return userservice.ErrNoSuchUser
	} else {
		return err
	}
}

func (r userRepository) Create(user *userservice.User) error {
	return r.DB.Create(user).Error
}

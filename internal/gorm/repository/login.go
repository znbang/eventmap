package repository

import (
	"github.com/znbang/eventmap/internal/dbx"
	"github.com/znbang/eventmap/pkg/login"
)

type userSessionRepository struct {
	*dbx.DB
}

func NewUserSessionRepository(db *dbx.DB) login.Repository {
	return &userSessionRepository{
		DB: db,
	}
}

func (r userSessionRepository) FindByID(userSession *login.UserSession, id string) error {
	return r.DB.Where("id=?", id).First(userSession).Error
}

func (r userSessionRepository) DeleteByID(id string) error {
	return r.DB.Where("id=?", id).Delete(&login.UserSession{}).Error
}

func (r userSessionRepository) Create(userSession *login.UserSession) error {
	return r.DB.Create(userSession).Error
}

func (r userSessionRepository) Save(userSession *login.UserSession) error {
	return r.DB.Save(userSession).Error
}

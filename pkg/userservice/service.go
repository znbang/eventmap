package userservice

import (
	"github.com/znbang/eventmap/pkg/uuid"
)

type UserService struct {
	UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{userRepository}
}

func (s *UserService) FindById(user *User, id string) error {
	return s.UserRepository.FindById(user, id)
}

func (s *UserService) FindByUid(user *User, uid string, provider int) error {
	return s.UserRepository.FindByUid(user, uid, provider)
}

func (s *UserService) Register(user *User) error {
	user.ID = uuid.RandomID()
	return s.UserRepository.Create(user)
}

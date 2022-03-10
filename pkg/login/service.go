package login

import (
	"github.com/znbang/eventmap/pkg/uuid"
)

type Service struct {
	Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		Repository: repository,
	}
}

func (s *Service) CreateUserSession(userSession *UserSession, userId string) error {
	*userSession = UserSession{
		ID:     uuid.RandomID(),
		UserID: userId,
	}

	return s.Repository.Create(userSession)
}

package login

import (
	"github.com/znbang/eventmap/pkg/uuid"
)

type Service struct {
	userSessionRepository UserSessionRepository
}

func NewService(userSessionRepository UserSessionRepository) *Service {
	return &Service{
		userSessionRepository: userSessionRepository,
	}
}

func (s *Service) CreateSession(userSession *UserSession, userId string) error {
	*userSession = UserSession{
		ID:     uuid.RandomID(),
		UserID: userId,
	}

	return s.userSessionRepository.Create(userSession)
}

func (s *Service) FindSessionById(userSession *UserSession, id string) error {
	return s.userSessionRepository.FindByID(userSession, id)
}

func (s *Service) SaveSession(userSession *UserSession) error {
	return s.userSessionRepository.Save(userSession)
}

func (s *Service) DeleteSessionById(id string) error {
	return s.userSessionRepository.DeleteByID(id)
}
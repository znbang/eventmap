package auth

import (
	"errors"
	"github.com/znbang/eventmap/pkg/userservice"
)

type Service struct {
	*userservice.UserService
}

func NewService(userService *userservice.UserService) *Service {
	return &Service{
		UserService: userService,
	}
}

func (s *Service) Login(user *userservice.User, provider string, accessToken string) error {
	var userInfo UserInfo

	if err := getUserInfo(&userInfo, provider, accessToken); err != nil {
		return err
	}

	*user = userInfo.ToUser()

	if err := s.UserService.FindByUid(user, user.Uid, user.Provider); errors.Is(err, userservice.ErrNoSuchUser) {
		return s.UserService.Register(user)
	} else {
		return err
	}
}

package auth

import (
	"errors"

	"github.com/znbang/eventmap/pkg/userservice"
)

type Service struct {
	userService *userservice.UserService
}

func NewService(userService *userservice.UserService) *Service {
	return &Service{
		userService: userService,
	}
}

func (s *Service) Login(user *userservice.User, provider string, accessToken string) error {
	var userInfo UserInfo

	if err := getUserInfo(&userInfo, provider, accessToken); err != nil {
		return err
	}

	*user = userInfo.ToUser()

	if err := s.userService.FindByUid(user, user.Uid, user.Provider); errors.Is(err, userservice.ErrNoSuchUser) {
		return s.userService.Register(user)
	} else {
		return err
	}
}

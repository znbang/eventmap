package userservice

import "errors"

var (
	ErrNoSuchUser = errors.New("no such user")
)

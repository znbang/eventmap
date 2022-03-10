package userservice

import "time"

const (
	LOCAL = iota
	GOOGLE
	FACEBOOK
	GITHUB
	LINE
)

type User struct {
	ID        string
	CreatedAt time.Time
	Provider  int
	Uid       string
	Email     string
	Name      string
}

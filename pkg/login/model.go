package login

import "time"

type UserSession struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    string
}

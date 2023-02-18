package auth

type contextKey int

const (
	sessionIdKey contextKey = iota + 1
	currentUserKey
)

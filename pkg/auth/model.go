package auth

type Authentication struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

package login

type UserSessionRepository interface {
	FindByID(userSession *UserSession, id string) error
	DeleteByID(id string) error
	Create(userSession *UserSession) error
	Save(userSession *UserSession) error
}

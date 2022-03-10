package login

type Repository interface {
	FindByID(userSession *UserSession, id string) error
	DeleteByID(id string) error
	Create(userSession *UserSession) error
	Save(userSession *UserSession) error
}

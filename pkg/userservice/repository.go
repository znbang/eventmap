package userservice

type UserRepository interface {
	FindById(user *User, id string) error
	FindByUid(user *User, uid string, provider int) error
	Create(user *User) error
}

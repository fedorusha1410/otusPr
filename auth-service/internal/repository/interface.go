package repository

import "auth-service/internal/model/user"

type UserRepository interface {
	GetUsers() []*user.User
	GetUserById(id int) *user.User
	UpdateUser(id int, data *user.User)
	DeleteUser(id int)
	Save(user user.User)
	SaveUserInFile()
	Restore()
}
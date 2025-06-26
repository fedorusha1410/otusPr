package repository

import "auth-service/internal/model/user"

type UserRepository interface {
	GetUsers() ([]*user.User, error)
	GetUserById(id int)  (*user.User, error)
	UpdateUser(id int, data *user.User) error
	DeleteUser(id int) error
	Save(user *user.User) error
	GetByUsername(username string) (*user.User, error)
}

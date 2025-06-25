package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Role int

const (
	Creator Role = iota
	Manager
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Role     Role   `json:"role"`
	Password string `json:"password"`
}

func NewObject() User {
	return User{}
}

func NewUser(name string, role Role, id int) User {
	return User{Name: name, Role: role, Id: id}
}

func (user *User) SetRole(role Role) {
	user.Role = role
}

func (user *User) GetRole() Role {
	return user.Role
}

func (user *User) GetId() int {
	return user.Id
}

func (user *User) SetId(id int) int {
	user.Id = id
	return user.Id
}

func (user *User) SetPassword(password string) error {
	passByBytes := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passByBytes, bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error of generate password hash")
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.New("password is wrong")
	}
	return nil

}

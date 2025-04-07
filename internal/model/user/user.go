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
	id       int
	Name     string
	role     Role
	password string
}

func NewUser(name string, role Role, id int) User {
	return User{Name: name, role: role, id: id}
}

func (user *User) SetRole(role Role) {
	user.role = role
}

func (user *User) GetRole() Role {
	return user.role
}

func (user *User) GetId() int {
	return user.id
}

func (user *User) SetId(id int) int {
	user.id = id
	return user.id
}

func (user *User) SetPassword(password string) error {
	passByBytes := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passByBytes, bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error of generate password hash")
	}
	user.password = string(hashedPassword)
	return nil
}

func (user *User) ComaprePassword(password string) (string, error) {
	passByBytes := []byte(password)
	err := bcrypt.CompareHashAndPassword([]byte(user.password), passByBytes)
	if err != nil {
		return "", errors.New("password is wrong")
	} else {
		return "password is correct", nil
	}

}

package user

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
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

func (user *User) SetPassword(password string) {
	passByBytes := []byte(password)
	hasher := sha1.New()
	hasher.Write(passByBytes)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	user.password = sha
}

func (user *User) ComaprePassword(password string) (string, error) {
	passByBytes := []byte(password)
	hasher := sha1.New()
	hasher.Write(passByBytes)
	encodedPass := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	if encodedPass == user.password {
		return "password is correct", nil
	} else {
		return "", errors.New("password is wrong")
	}
}

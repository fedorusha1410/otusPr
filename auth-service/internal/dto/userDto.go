package dto

import "auth-service/internal/model/user"

type CreateUserDto struct {
	Id       int    `json:"id"`
	Name     string `json:"Name"`
	Role     int    `json:"Role"`
	Password string `json:"password"`
}

type UpdateUserDto struct {
	Name     string `json:"Name"`
	Password string `json:"password"`
}

func MapToUserModel(dto CreateUserDto) user.User {

	role := user.Creator
	if dto.Role == 1 {
		role = user.Manager
	}

	return user.NewUser(dto.Name, role, dto.Id)
}

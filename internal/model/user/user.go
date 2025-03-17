package user

type User struct {
	id   int
	Name string
	role string
}

func (user *User) GetRole() string {
	return user.role
}

func (user *User) SetRole(role string) string {
	user.role = role
	return user.role
}

func (user *User) GetId() int {
	return user.id
}

func (user *User) SetId(id int) int {
	user.id = id
	return user.id
}

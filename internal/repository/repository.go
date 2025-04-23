package repository

import (
	"otus/internal/model/task"
	"otus/internal/model/user"
)

type Repository struct {
	Tasks []*task.Task
	Users []*user.User
}

func New() Repository {
	return Repository{}
}
func (repository *Repository) Save(param any) {

	switch value := param.(type) {
	case task.Task:
		repository.Tasks = append(repository.Tasks, &value)
	case user.User:
		repository.Users = append(repository.Users, &value)
	}
}

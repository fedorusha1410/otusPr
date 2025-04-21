package repository

import (
	"otus/internal/model/task"
	"otus/internal/model/user"
	"sync"
)

type Repository struct {
	Tasks []*task.Task
	Users []*user.User
}

func New() Repository {
	return Repository{}
}
func (repository *Repository) Save(cwg *sync.WaitGroup, ch <-chan interface{}) {

	defer cwg.Done()
	for val := range ch {
		switch value := val.(type) {
		case task.Task:
			repository.Tasks = append(repository.Tasks, &value)
		case user.User:
			repository.Users = append(repository.Users, &value)
		}

	}
}

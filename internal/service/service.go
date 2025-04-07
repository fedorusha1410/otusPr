package service

import (
	"errors"
	"otus/internal/model/task"
	"otus/internal/model/user"
	"otus/internal/repository"
	"time"
)


func Create(repo *repository.Repository, params ...any) error {

	if len(params) == 3 {
		result := user.NewUser(params[0].(string), params[1].(user.Role), params[2].(int))
		repo.Save(result)
		return nil
	}
	if len(params) == 7 {
		result := task.NewTask(params[0].(int), params[1].(string),
			params[2].(string), params[3].(string), params[4].(time.Time),
			params[5].(string), params[6].(int))
		repo.Save(result)
		return nil
	}

	return errors.New("error of create struct, incorrect number of input parameters")
}

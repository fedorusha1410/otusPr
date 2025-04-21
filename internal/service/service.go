package service

import (
	"errors"
	"otus/internal/model/task"
	"otus/internal/model/user"
	"time"
)

func Create(params ...any) (interface{}, error) {

	if len(params) == 3 {
		result := user.NewUser(params[0].(string), params[1].(user.Role), params[2].(int))
		return result, nil
	}
	if len(params) == 7 {
		result := task.NewTask(params[0].(int), params[1].(string),
			params[2].(string), params[3].(string), params[4].(time.Time),
			params[5].(string), params[6].(int))
		return result, nil

	}

	return nil, errors.New("error of create struct, incorrect number of input parameters")
}

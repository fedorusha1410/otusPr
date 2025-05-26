package dto

import (
	"otus/internal/model/task"
	"time"
)

type CreateTaskDto struct {
	Id       int    `json:"Id"`
	Status   string `json:"Status"`
	Title    string `json:"Title"`
	Note     string `json:"Note"`
	Priority string `json:"Priority"`
	AuthorId int    `json:"authorId"`
}

type UpdateTaskDto struct {
	Status   string `json:"Status"`
	Title    string `json:"Title"`
	Note     string `json:"Note"`
	Priority string `json:"Priority"`
}

func MapToTaskModel(dto CreateTaskDto) task.Task {

	return task.NewTask(
		dto.Id,
		dto.Status,
		dto.Title,
		dto.Note,
		time.Now(),
		dto.Priority,
		dto.AuthorId)
}

package repository

import "task-manager/internal/model/task"

type TaskRepository interface {
	GetTasks() ([]*task.Task, error)
	GetTaskById(id int) (*task.Task, error)
	UpdateTask(id int, newData *task.Task) error
	DeleteTask(id int) error
	Save(task task.Task) error
	SaveTaskInFile()
	Restore()
}



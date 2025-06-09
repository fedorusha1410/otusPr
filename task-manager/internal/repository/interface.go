package repository

import "task-manager/internal/model/task"

type TaskRepository interface {
	GetTasks() []*task.Task
	GetTaskById(id int) *task.Task
	UpdateTask(id int, newData *task.Task)
	DeleteTask(id int)
	Save(task task.Task)
	SaveTaskInFile()
	Restore()
}



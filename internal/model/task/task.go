package task

import "time"

type Task struct {
	id          int
	Status      string
	Title       string
	Note        string
	CreatedTime time.Time
	UpdatedTime time.Time
	Priority    string
	authorId    int
}

func (task *Task) GetAuthorId() int {
	return task.authorId
}

func (task *Task) SetAuthorId(authorId int) int {
	task.authorId = authorId
	return task.authorId
}

func (task *Task) GetId() int {
	return task.id
}

func (task *Task) SetId(id int) int {
	task.id = id
	return task.id
}

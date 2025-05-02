package task

import "time"

type Task struct {
	Id          int       `json:"Id"`
	Status      string    `json:"Status"`
	Title       string    `json:"Title"`
	Note        string    `json:"Note"`
	CreatedTime time.Time `json:"CreatedTime"`
	UpdatedTime time.Time `json:"UpdatedTime"`
	Priority    string    `json:"Priority"`
	AuthorId    int       `json:"authorId"`
}

func NewObject() Task {
	return Task{}
}
func NewTask(id int, status string, title string, note string, createdTime time.Time, priority string, authorId int) Task {
	return Task{Id: id, Status: status, Title: title, Note: note, CreatedTime: createdTime, Priority: priority, AuthorId: authorId}
}

func (task *Task) GetAuthorId() int {
	return task.AuthorId
}

func (task *Task) SetAuthorId(authorId int) int {
	task.AuthorId = authorId
	return task.AuthorId
}

func (task *Task) GetId() int {
	return task.Id
}

func (task *Task) SetId(id int) int {
	task.Id = id
	return task.Id
}

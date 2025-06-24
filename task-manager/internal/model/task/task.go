package task

import "time"

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	Note        string    `json:"note"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
	Priority    string    `json:"priority"`
	AuthorId    int       `json:"author_id"`
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

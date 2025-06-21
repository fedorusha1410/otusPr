package task

import "time"

type Task struct {
	Id          int        `json:"Id" bson:"id"`
	Status      string    `json:"Status" bson:"status"`
	Title       string    `json:"Title" bson:"title"`
	Note        string    `json:"Note" bson:"note"`
	CreatedTime time.Time `json:"CreatedTime" bson:"created_time"`
	UpdatedTime time.Time `json:"UpdatedTime" bson:"updated_time"`
	Priority    string    `json:"Priority" bson:"priority"`
	AuthorId    int       `json:"authorId" bson:"author_id"`
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

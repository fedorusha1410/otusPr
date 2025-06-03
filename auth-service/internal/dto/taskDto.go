package dto

type CreateTaskRequest struct {
	Id       int    `json:"Id"`
	Title    string `json:"Title"`
	Note     string `json:"Note"`
	Priority string `json:"Priority"`
	AuthorId int    `json:"authorId"`
}

type CreateTaskResponse struct {
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

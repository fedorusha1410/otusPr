package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"task-manager/internal/model/task"
)

const taskFile = "tasks.json"

type Repository struct {
	Tasks []*task.Task
}

func New() Repository {
	return Repository{}
}

func (repository *Repository) GetTasks() []*task.Task {
	return repository.Tasks
}

func (repository *Repository) GetTaskById(id int) *task.Task {

	for _, task := range repository.Tasks {
		if task.Id == id {
			return task
		}
	}
	return nil
}

func (repository *Repository) UpdateTask(id int, newData *task.Task) {

	for _, task := range repository.Tasks {
		if task.Id == id {
			task.Note = newData.Note
			task.Priority = newData.Priority
			task.Status = newData.Status
			task.Title = newData.Title
			task.UpdatedTime = newData.UpdatedTime
		}
	}
}

func (repository *Repository) DeleteTask(id int) {

	for i, task := range repository.Tasks {
		if task.Id == id {
			repository.Tasks = append(repository.Tasks[:i], repository.Tasks[i+1:]...)
		}
	}
}

func (repository *Repository) Save(task task.Task) {
	repository.Tasks = append(repository.Tasks, &task)
}

func (repository *Repository) SaveTaskInFile() {
	file, err := os.OpenFile(taskFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error of opening task file")
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(repository.Tasks)
	if err != nil {
		fmt.Println("Error of writing task: ", err)
		return
	}
}

func (repository *Repository) Restore() {
	fileTask, err := os.Open(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File not found, create file")
			fileTask, err = os.Create(taskFile)
			if err != nil {
				fmt.Println("Error of creating 'task' file:", err)
				return
			}
		} else {
			fmt.Println("Error of opening 'task' file:", err)
			return
		}
	}
	fileData, err := os.ReadFile(taskFile)
	if err == nil && len(fileData) > 0 {
		err = json.Unmarshal(fileData, &repository.Tasks)
		if err != nil {
			fmt.Println("error decoding existing tasks: %w", err)
		}
	}

	defer fileTask.Close()

}

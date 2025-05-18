package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"otus/internal/model/task"
	"otus/internal/model/user"
)

const taskFile = "tasks.json"
const userFile = "users.json"

type Repository struct {
	Tasks []*task.Task
	Users []*user.User
}

func New() Repository {
	return Repository{}
}

func (repository *Repository) GetTasks() []*task.Task {
	return repository.Tasks
}

func (repository *Repository) GetUsers() []*user.User {
	return repository.Users
}

func (repository *Repository) GetUserById(id int) *user.User {

	for _, user := range repository.Users {
		if user.Id == id {
			return user
		}
	}
	return nil
}

func (repository *Repository) GetTaskById(id int) *task.Task {

	for _, task := range repository.Tasks {
		if task.Id == id {
			return task
		}
	}
	return nil
}

func (repository *Repository) UpdateUser(id int, newData *user.User) {

	for _, user := range repository.Users {
		if user.Id == id {
			user.Name = newData.Name
			user.Role = newData.Role
		}
	}
}

func (repository *Repository) UpdateTask(id int, newData *task.Task) {

	for _, task := range repository.Tasks {
		if task.Id == id {
			task.Note = newData.Note
			task.Priority = newData.Priority
		}
	}
}

func (repository *Repository) DeleteUser(id int) {

	for i, user := range repository.Users {
		if user.Id == id {
			repository.Users = append(repository.Users[:i], repository.Users[i+1:]...)
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

func (repository *Repository) Save(param any) {

	switch value := param.(type) {
	case task.Task:
		repository.Tasks = append(repository.Tasks, &value)
	case user.User:
		repository.Users = append(repository.Users, &value)
	}
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

func (repository *Repository) SaveUserInFile() {
	file, err := os.OpenFile(userFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error of opening task file")
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(repository.Users)
	if err != nil {
		fmt.Println("Error of writing user:", err)
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

	fileUser, err := os.Open(userFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File not found, create file")
			fileUser, err = os.Create(userFile)
			if err != nil {
				fmt.Println("Error of creating 'user' file:", err)
				return
			}
		} else {
			fmt.Println("Error of opening 'user' file:", err)
			return
		}
	}

	fileData, err = os.ReadFile(userFile)
	if err == nil && len(fileData) > 0 {
		err = json.Unmarshal(fileData, &repository.Users)
		if err != nil {
			fmt.Println("error decoding existing users: %w", err)
		}
	}

	defer fileUser.Close()
}

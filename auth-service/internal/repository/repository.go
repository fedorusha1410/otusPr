package repository

import (
	"auth-service/internal/model/user"
	"encoding/json"
	"fmt"
	"os"
)

const userFile = "users.json"

type Repository struct {
	Users []*user.User
}

func New() *Repository {
	return &Repository{}
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

func (repository *Repository) UpdateUser(id int, newData *user.User) {

	for _, user := range repository.Users {
		if user.Id == id {
			user.Name = newData.Name
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

func (repository *Repository) Save(newUser user.User) {
	repository.Users = append(repository.Users, &newUser)

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

	fileData, err := os.ReadFile(userFile)
	if err == nil && len(fileData) > 0 {
		err = json.Unmarshal(fileData, &repository.Users)
		if err != nil {
			fmt.Println("error decoding existing users:", err)
		}
	}

	defer fileUser.Close()
}

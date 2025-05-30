package userHandler

import (
	"encoding/json"
	"net/http"
	"auth-service/internal/dto"
	"auth-service/internal/repository"
	"strconv"
	"strings"
)

// Get User by id dosc
// @Summary      Get User
// @Description  Get User by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  user.User
// @Router       /users/{id} [get]
func GetById(w http.ResponseWriter, request *http.Request, repo *repository.Repository) {
	path := request.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	idStr := parts[2]
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	user := repo.GetUserById(userID)

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// @Summary      Get Users
// @Description  Get All Users from file
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}  user.User
// @Router       /users/ [get]
func GetAll(w http.ResponseWriter, request *http.Request, repo *repository.Repository) {

	users := repo.GetUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// @Summary      Insert user
// @Description  Insert user into slice and file
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      dto.CreateUserDto  true  "User to create"
// @Success      200   {object}  dto.CreateUserDto
// @Failure      400   {string}  string  "Invalid input"
// @Router       /users/ [post]
// @Security BearerAuth
func Insert(w http.ResponseWriter, request *http.Request, repo *repository.Repository) {

	var newUser dto.CreateUserDto
	err := json.NewDecoder(request.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	userOld := repo.GetUserById(newUser.Id)

	if userOld != nil {
		http.Error(w, "User with that ID already exists, try other", http.StatusBadRequest)
		return
	}

	if newUser.Name == "" {
		http.Error(w, "User name is required parametr", http.StatusBadRequest)
		return
	}

	if newUser.Role > 1 || newUser.Role < 0 {
		http.Error(w, "Ivalid user role", http.StatusBadRequest)
		return
	}
	user := dto.MapToUserModel(newUser)
	user.SetPassword(newUser.Password)
	repo.Save(user)
	repo.SaveUserInFile()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

// @Summary      Update user by ID
// @Description  Update user in slice and file
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int        true  "User ID"
// @Param        user  body      dto.UpdateUserDto  true  "User to update"
// @Success      200
// @Failure      400   {string}  string  "Invalid input"
// @Router       /users/{id} [put]
// @Security BearerAuth
func Update(w http.ResponseWriter, request *http.Request, repo *repository.Repository) {
	path := request.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	idStr := parts[2]
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	var newUser dto.UpdateUserDto
	err = json.NewDecoder(request.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	user := repo.GetUserById(userID)

	if user != nil {
		user.Name = newUser.Name
		user.SetPassword(newUser.Password)
		repo.UpdateUser(userID, user)
		repo.SaveUserInFile()
	} else {
		http.Error(w, "User with this ID doesnt exist", http.StatusBadRequest)
	}

}

// @Summary      Delete user by ID
// @Description  Delete user in slice and file
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200
// @Failure      400  {string}  string  "Invalid input"
// @Router       /users/{id} [delete]
// @Security BearerAuth
func Delete(w http.ResponseWriter, request *http.Request, repo *repository.Repository) {
	path := request.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	idStr := parts[2]
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user := repo.GetUserById(userID)

	if user != nil {
		repo.DeleteUser(userID)
		repo.SaveUserInFile()
	} else {
		http.Error(w, "User with this ID doesnt exist", http.StatusBadRequest)
	}
}

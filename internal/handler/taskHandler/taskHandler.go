package taskHandler

import (
	"encoding/json"
	"net/http"
	"otus/internal/model/task"
	"otus/internal/repository"
	"strconv"
	"strings"
)

func GetById(w http.ResponseWriter, request *http.Request, repo *repository.Repository) {
	path := request.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	idStr := parts[2]
	taskID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	task := repo.GetTaskById(taskID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func GetAll(w http.ResponseWriter, request *http.Request, repo *repository.Repository) {

	tasks := repo.GetTasks()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func Insert(w http.ResponseWriter, request *http.Request, repo *repository.Repository) {

	var newTask task.Task
	err := json.NewDecoder(request.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	taskOld := repo.GetTaskById(newTask.Id)

	if taskOld != nil {
		http.Error(w, "Task with that ID already exists, try other", http.StatusBadRequest)
		return
	}

	if newTask.Title == "" {
		http.Error(w, "Task title is required parametr", http.StatusBadRequest)
		return
	}

	if newTask.Note == "" {
		http.Error(w, "Task note is required", http.StatusBadRequest)
		return
	}

	if newTask.AuthorId == 0 {
		http.Error(w, "Task Author Id is required parametr", http.StatusBadRequest)
		return
	}

	repo.Save(newTask)
	repo.SaveTaskInFile()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}

func Update(w http.ResponseWriter, request *http.Request, repo *repository.Repository) {
	path := request.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	idStr := parts[2]
	taskID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	var newTask task.Task
	err = json.NewDecoder(request.Body).Decode(&newTask)

	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	task := repo.GetTaskById(taskID)

	if task != nil {
		repo.UpdateTask(taskID, &newTask)
		repo.SaveTaskInFile()
	} else {
		http.Error(w, "task with this ID doesnt exist", http.StatusBadRequest)
	}

}

func Delete(w http.ResponseWriter, request *http.Request, repo *repository.Repository) {
	path := request.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	idStr := parts[2]
	taskID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task := repo.GetTaskById(taskID)

	if task != nil {
		repo.DeleteTask(taskID)
		repo.SaveTaskInFile()
	} else {
		http.Error(w, "task with this ID doesnt exist", http.StatusBadRequest)
	}
}

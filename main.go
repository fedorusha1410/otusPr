package main

import (
	"net/http"
	"otus/internal/handler/taskHandler"
	"otus/internal/handler/userHandler"
	"otus/internal/repository"
	"strings"
)

func main() {

	repository := repository.New()
	repository.Restore()

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

			if len(pathParts) == 1 {
				userHandler.GetAll(w, r, &repository)

			} else if len(pathParts) == 2 {
				userHandler.GetById(w, r, &repository)
			} else {
				http.Error(w, "Invalid URL", http.StatusBadRequest)
			}
		case http.MethodPost:
			userHandler.Insert(w, r, &repository)
		case http.MethodPut:
			userHandler.Update(w, r, &repository)
		case http.MethodDelete:
			userHandler.Delete(w, r, &repository)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

			if len(pathParts) == 1 {
				taskHandler.GetAll(w, r, &repository)

			} else if len(pathParts) == 2 {
				taskHandler.GetById(w, r, &repository)
			} else {
				http.Error(w, "Invalid URL", http.StatusBadRequest)
			}
		case http.MethodPost:
			taskHandler.Insert(w, r, &repository)
		case http.MethodPut:
			taskHandler.Update(w, r, &repository)
		case http.MethodDelete:
			taskHandler.Delete(w, r, &repository)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8090", nil)
}

package main

import (
	"log"
	"net/http"
	_ "otus/docs"
	"otus/internal/handler/authHandler"
	"otus/internal/handler/taskHandler"
	"otus/internal/handler/userHandler"
	"otus/internal/middleware"
	"otus/internal/repository"
	"strings"

	"github.com/joho/godotenv"
)

// @title           Task Manager API
// @description     This is a sample server for managing tasks.
// @version         1.0
// @host      		localhost:8090
// @securityDefinitions.apikey BearerAuth
// @tokenUrl  /login
// @in header
// @name Authorization
func main() {

	repository := repository.New()
	repository.Restore()

	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	http.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./docs"))))
	http.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger-ui"))))
	http.HandleFunc("/login", authHandler.Login)

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
			middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
				userHandler.Insert(w, r, &repository)
			})(w, r)
		case http.MethodPut:
			middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
				userHandler.Update(w, r, &repository)
			})(w, r)
		case http.MethodDelete:
			middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
				userHandler.Delete(w, r, &repository)
			})(w, r)
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
			middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
				taskHandler.Insert(w, r, &repository)
			})(w, r)
		case http.MethodPut:
			middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
				taskHandler.Update(w, r, &repository)
			})(w, r)
		case http.MethodDelete:
			middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
				taskHandler.Delete(w, r, &repository)
			})(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8090", nil)
}

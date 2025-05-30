package taskHandler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"task-manager/pb"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"
)

var taskClient pb.TaskServiceClient

func Init(client pb.TaskServiceClient) {
	taskClient = client
}

// @Summary      Get Tasks
// @Description  Get All Tasks from file
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Success      200  {array}  dto.CreateTaskDto
// @Router       /tasks/ [get]
func GetAll(w http.ResponseWriter, r *http.Request) {

	resp, err := taskClient.GetAllTasks(context.Background(), &emptypb.Empty{})
	if err != nil {
		handleGrpcError(w, err)
		return
	}

	marshaller := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}

	data, err := marshaller.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// Get Task by id dosc
// @Summary      Get Task
// @Description  Get Task by ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Success      200  {object}  dto.CreateTaskDto
// @Router       /tasks/{id} [get]
func GetById(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
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

	resp, err := taskClient.GetTaskById(context.Background(), &pb.GetTaskRequest{Id: int32(taskID)})
	if err != nil {
		handleGrpcError(w, err)
		return
	}

	marshaller := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}

	data, err := marshaller.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// @Summary      Insert task
// @Description  Insert task into slice and file
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        task  body      dto.CreateTaskRequest  true  "Task to create"
// @Success      200   {object}  dto.CreateTaskResponse
// @Failure      400   {string}  string  "Invalid input"
// @Router       /tasks/ [post]
// @Security BearerAuth
func Insert(w http.ResponseWriter, r *http.Request) {

	newTask := &pb.CreateTaskRequest{}
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	resp, err := taskClient.CreateTask(context.Background(), newTask)
	if err != nil {
		handleGrpcError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// @Summary      Update task by ID
// @Description  Update task in slice and file
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id    path      int        true  "Task ID"
// @Param        task  body      dto.UpdateTaskDto  true  "Task to update"
// @Success      200
// @Failure      400   {string}  string  "Invalid input"
// @Router       /tasks/{id} [put]
// @Security BearerAuth
func Update(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
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

	updatedTask := &pb.UpdateTaskRequest{}
	err = json.NewDecoder(r.Body).Decode(&updatedTask)

	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	updatedTask.Id = int32(taskID)

	resp, err := taskClient.UpdateTask(context.Background(), updatedTask)
	if err != nil {
		handleGrpcError(w, err)
		return
	}

	log.Printf("response from server: %s", resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// @Summary      Delete task by ID
// @Description  Delete task in slice and file
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Success      200
// @Failure      400  {string}  string  "Invalid input"
// @Router       /tasks/{id} [delete]
// @Security BearerAuth
func Delete(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
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

	resp, err := taskClient.DeleteTask(context.Background(), &pb.DeleteTaskRequest{Id: int32(taskID)})
	if err != nil {
		handleGrpcError(w, err)
		return
	}

	fmt.Printf("response from server: %s", resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func handleGrpcError(w http.ResponseWriter, err error) {
	st, ok := status.FromError(err)
	if !ok {
		http.Error(w, "Unknown error", http.StatusInternalServerError)
		return
	}

	code := http.StatusInternalServerError
	if st.Code() == 5 {
		code = http.StatusNotFound
	}

	http.Error(w, st.Message(), code)
}

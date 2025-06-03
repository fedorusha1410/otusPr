package grpc

import (
	"context"
	"log"
	"task-manager/internal/model/task"
	"task-manager/internal/repository"
	"task-manager/pb"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskServer struct {
	pb.UnimplementedTaskServiceServer
	repo *repository.Repository
}

func NewTaskServer(repository *repository.Repository) *TaskServer {
	return &TaskServer{
		repo: repository,
	}
}

func (s *TaskServer) GetTaskById(ctx context.Context, req *pb.GetTaskRequest) (*pb.TaskResponse, error) {

	task := s.repo.GetTaskById(int(req.Id))

	if task == nil {
		return nil, status.Errorf(codes.NotFound, "task with ID %d not found", req.Id)
	}

	return &pb.TaskResponse{
		Id:          int32(task.Id),
		AuthorId:    int32(task.AuthorId),
		Note:        task.Note,
		Status:      task.Status,
		Title:       task.Title,
		Priority:    task.Priority,
		CreatedTime: timestamppb.New(task.CreatedTime),
		UpdatedTime: timestamppb.New(task.UpdatedTime),
	}, nil

}

func (s *TaskServer) GetAllTasks(ctx context.Context, req *emptypb.Empty) (*pb.TaskListResponse, error) {

	tasks := s.repo.GetTasks()

	var pbTasks []*pb.TaskResponse
	for _, t := range tasks {
		pbTask := &pb.TaskResponse{
			Id:          int32(t.Id),
			Title:       t.Title,
			Note:        t.Note,
			Status:      t.Status,
			Priority:    t.Priority,
			CreatedTime: timestamppb.New(t.CreatedTime),
			UpdatedTime: timestamppb.New(t.UpdatedTime),
		}
		pbTasks = append(pbTasks, pbTask)
	}

	return &pb.TaskListResponse{
		Tasks: pbTasks,
	}, nil

}

func (s *TaskServer) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	log.Printf("CreateTask, Request: %+v\n", req)

	newTask := task.Task{
		Title:       req.Title,
		Note:        req.Note,
		Priority:    req.Priority,
		AuthorId:    int(req.AuthorId),
		Status:      "new",
		CreatedTime: time.Now(),
		Id:          int(req.Id),
	}

	s.repo.Save(newTask)
	s.repo.SaveTaskInFile()

	return &pb.CreateTaskResponse{
		Id:       int32(newTask.Id),
		Status:   newTask.Status,
		Title:    newTask.Title,
		Note:     newTask.Note,
		Priority: newTask.Priority,
		AuthorId: int32(newTask.AuthorId),
	}, nil

}

func (s *TaskServer) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*emptypb.Empty, error) {
	log.Printf("UpdateTask, Request: %+v\n", req)
	task := s.repo.GetTaskById(int(req.Id))

	if task == nil {
		return nil, status.Errorf(codes.NotFound, "task with ID %d not found", req.Id)
	}

	s.repo.UpdateTask(int(req.Id), task)
	s.repo.SaveTaskInFile()

	return &emptypb.Empty{}, nil
}

func (s *TaskServer) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*emptypb.Empty, error) {

	task := s.repo.GetTaskById(int(req.Id))

	if task == nil {
		return nil, status.Errorf(codes.NotFound, "task with ID %d not found", req.Id)
	}

	s.repo.DeleteTask(int(req.Id))
	s.repo.SaveTaskInFile()

	return &emptypb.Empty{}, nil
}

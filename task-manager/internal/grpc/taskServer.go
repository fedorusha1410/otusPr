package grpc

import (
	"context"
	"log"
	"task-manager/internal/model/task"
	"task-manager/internal/service/task"

	"task-manager/pb"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskServer struct {
	pb.UnimplementedTaskServiceServer
	service *service.Service
}

func NewTaskServer(s *service.Service) *TaskServer {
	return &TaskServer{
		service: s,
	}
}

func (s *TaskServer) GetTaskById(ctx context.Context, req *pb.GetTaskRequest) (*pb.TaskResponse, error) {

	task, err := s.service.GetTaskByID(ctx, int(req.Id))

	if err != nil {
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

	tasks, err := s.service.GetTasks(ctx)

	if err != nil {
		return nil, err
	}

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

	task, err := s.service.CreateTask(ctx, &newTask)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Error: = %s ", err)
	}
	s.service.SaveInFile()

	return &pb.CreateTaskResponse{
		Id:       int32(newTask.Id),
		Status:   task.Status,
		Title:    task.Title,
		Note:     task.Note,
		Priority: task.Priority,
		AuthorId: int32(task.AuthorId),
	}, nil

}

func (s *TaskServer) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*emptypb.Empty, error) {
	log.Printf("UpdateTask, Request: %+v\n", req)

	newTask := task.Task{
		Title:    req.Title,
		Note:     req.Note,
		Priority: req.Priority,
		Status:   req.Status,
		Id:       int(req.Id),
	}

	err := s.service.UpdateTask(ctx, int(req.Id), &newTask)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Error: = %s ", err)
	}

	s.service.SaveInFile()

	return &emptypb.Empty{}, nil
}

func (s *TaskServer) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*emptypb.Empty, error) {

	err := s.service.DeleteTask(ctx, int(req.Id))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Error: =  %d ", err)
	}
	s.service.SaveInFile()

	return &emptypb.Empty{}, nil
}

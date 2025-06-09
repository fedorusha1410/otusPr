package main

import (
	"log"
	"net"
	mygrpc "task-manager/internal/grpc"
	"task-manager/internal/repository"
	"task-manager/internal/service/task"
	"task-manager/pb"

	"google.golang.org/grpc"
)

func main() {
	repo := repository.New()
	service := service.New(&repo)
	service.Restore()

	grpcServer := grpc.NewServer()

	taskServer := mygrpc.NewTaskServer(service)
	pb.RegisterTaskServiceServer(grpcServer, taskServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}
	log.Println("gRPC server is running on port 50051...")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

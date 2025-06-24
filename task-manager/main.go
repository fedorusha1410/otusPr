package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	mygrpc "task-manager/internal/grpc"
	redislog "task-manager/internal/logger"
	"task-manager/internal/repository"
	service "task-manager/internal/service/task"
	"task-manager/pb"
	"time"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db := connectPostgres()
	defer db.Close()
	repository := repository.New(db)
	logger := redislog.NewRedisLogger("localhost:6379")

	service := service.New(*repository, *logger)
	logger.LogPrinter(ctx, 20*time.Second)
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

func connectPostgres() *sql.DB {

	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/taskdb?sslmode=disable")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	return db
}

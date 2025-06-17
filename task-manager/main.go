package main

import (
	"context"
	"fmt"
	"log"
	"net"
	mygrpc "task-manager/internal/grpc"
	redislog "task-manager/internal/logger"
	"task-manager/internal/repository"
	service "task-manager/internal/service/task"
	"task-manager/pb"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client := connectMongo()
	collection := client.Database("taskdb").Collection("tasks")
	mongoRepo := repository.NewMongoRepository(collection)
	logger := redislog.NewRedisLogger("localhost:6379")

	service := service.New(*mongoRepo, *logger)
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

func connectMongo() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Mongo connect error: %v", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Mongo ping error: %v", err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}

package repository

import (
	"context"
	"fmt"
	"task-manager/internal/model/task"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(collection *mongo.Collection) *MongoRepository {
	return &MongoRepository{
		collection: collection,
	}
}

func (r *MongoRepository) GetTasks() ([]*task.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("GetTasks error:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []*task.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		fmt.Println("cursor.All error:", err)
		return nil, err
	}
	return tasks, nil
}

func (r *MongoRepository) GetTaskById(id int) (*task.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var t task.Task
	err := r.collection.FindOne(ctx, bson.M{"id": id}).Decode(&t)
	if err == mongo.ErrNoDocuments {
		return nil, err
	} else if err != nil {
		fmt.Println("GetTaskById error:", err)
		return nil, err
	}
	return &t, nil
}

func (r *MongoRepository) UpdateTask(id int, newData *task.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"note":         newData.Note,
			"priority":     newData.Priority,
			"status":       newData.Status,
			"title":        newData.Title,
			"updated_time": time.Now(),
		},
	}

	_, err := r.collection.UpdateOne(ctx, bson.M{"id": id}, update)
	if err != nil {
		fmt.Println("UpdateTask error:", err)
		return err
	}
	return nil
}

func (r *MongoRepository) DeleteTask(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		fmt.Println("DeleteTask error:", err)
		return err
	}
	return nil
}

func (r *MongoRepository) Save(t task.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, t)
	if err != nil {
		fmt.Println("Save error:", err)
		return err
	}

	return nil
}

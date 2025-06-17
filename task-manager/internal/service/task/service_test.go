package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"task-manager/internal/model/task"
	"task-manager/internal/service/task"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) GetTasks() []*task.Task {
	args := m.Called()
	return args.Get(0).([]*task.Task)
}

func (m *MockRepo) GetTaskById(id int) *task.Task {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*task.Task)
}

func (m *MockRepo) UpdateTask(id int, newData *task.Task) {
	m.Called(id, newData)
}

func (m *MockRepo) DeleteTask(id int) {
	m.Called(id)
}

func (m *MockRepo) Save(t task.Task) {
	m.Called(t)
}

func (m *MockRepo) SaveTaskInFile() {}
func (m *MockRepo) Restore()        {}

func TestGetTaskByID(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := service.New(mockRepo)

	t1 := &task.Task{Id: 1, Title: "Test Task"}
	mockRepo.On("GetTaskById", 1).Return(t1)
	mockRepo.On("GetTaskById", 42).Return(nil)

	tests := []struct {
		name    string
		id      int
		want    *task.Task
		wantErr error
	}{
		{"task found", 1, t1, nil},
		{"task not found", 42, nil, service.ErrNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := svc.GetTaskByID(context.Background(), tt.id)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestCreateTask(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := service.New(mockRepo)

	valid := &task.Task{Id: 10, Title: "Do this"}
	mockRepo.On("Save", mock.AnythingOfType("task.Task")).Return()

	t.Run("valid create", func(t *testing.T) {
		res, err := svc.CreateTask(context.Background(), valid)
		assert.NoError(t, err)
		assert.Equal(t, "Do this", res.Title)
		assert.WithinDuration(t, time.Now(), res.CreatedTime, time.Second)
	})

	t.Run("invalid (empty title)", func(t *testing.T) {
		_, err := svc.CreateTask(context.Background(), &task.Task{})
		assert.True(t, errors.Is(err, service.ErrInvalid))
	})
}

func TestDeleteTask(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := service.New(mockRepo)

	mockRepo.On("GetTaskById", 1).Return(&task.Task{Id: 1})
	mockRepo.On("GetTaskById", 99).Return(nil)
	mockRepo.On("DeleteTask", 1).Return()

	t.Run("delete ok", func(t *testing.T) {
		err := svc.DeleteTask(context.Background(), 1)
		assert.NoError(t, err)
	})

	t.Run("delete not found", func(t *testing.T) {
		err := svc.DeleteTask(context.Background(), 99)
		assert.ErrorIs(t, err, service.ErrNotFound)
	})
}

func TestUpdateTask(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := service.New(mockRepo)

	existing := &task.Task{Id: 1, Title: "Old"}
	updated := &task.Task{Title: "New", Note: "Updated note", Priority: "High", Status: "Open"}

	mockRepo.On("GetTaskById", 1).Return(existing)
	mockRepo.On("GetTaskById", 99).Return(nil)
	mockRepo.On("UpdateTask", 1, updated).Return()

	t.Run("update ok", func(t *testing.T) {
		err := svc.UpdateTask(context.Background(), 1, updated)
		assert.NoError(t, err)
	})

	t.Run("update not found", func(t *testing.T) {
		err := svc.UpdateTask(context.Background(), 99, updated)
		assert.ErrorIs(t, err, service.ErrNotFound)
	})
}

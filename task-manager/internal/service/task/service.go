package service

import (
	"context"
	"errors"
	redislog "task-manager/internal/logger"
	"task-manager/internal/model/task"
	"task-manager/internal/repository"
	"time"
)

type Service struct {
	repo   repository.Repository
	logger redislog.RedisLogger
}

var (
	ErrNotFound = errors.New("task not found")
	ErrInvalid  = errors.New("invalid task input")
)

func New(repo repository.Repository, logger redislog.RedisLogger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) CreateTask(ctx context.Context, t *task.Task) (*task.Task, error) {
	if t == nil || t.Title == "" {
		return nil, ErrInvalid
	}

	t.CreatedTime = time.Now()
	t.UpdatedTime = t.CreatedTime
	_ = s.logger.LogAction(ctx, "create", "task", map[string]interface{}{
		"id":    t.Id,
		"title": t.Title,
		"time":  t.CreatedTime.Format(time.RFC3339),
	})

	s.repo.Save(t)

	return t, nil
}

func (s *Service) GetTaskByID(ctx context.Context, id int) (*task.Task, error) {
	t, err := s.repo.GetTaskById(id)

	if err != nil {
		return nil, err
	}
	if t == nil {
		return nil, ErrNotFound
	}

	return t, nil
}

func (s *Service) GetTasks(ctx context.Context, userId int, userRole string, filter *task.TaskFilter) ([]*task.Task, error) {
	return s.repo.GetTasks(userId, userRole, filter)
}

func (s *Service) UpdateTask(ctx context.Context, id int, newData *task.Task) error {
	old, err := s.repo.GetTaskById(id)

	if err != nil{
		return err
	}

	if old == nil {
		return ErrNotFound
	}

	newData.UpdatedTime = time.Now()
	s.repo.UpdateTask(id, newData)

	_ = s.logger.LogAction(ctx, "update", "task", map[string]interface{}{
		"id":    id,
		"title": newData.Title,
		"time":  newData.UpdatedTime.Format(time.RFC3339),
	})

	return nil
}

func (s *Service) DeleteTask(ctx context.Context, id int) error {
	t, err := s.repo.GetTaskById(id)

	if err != nil {
		return err
	}
	if t == nil {
		return ErrNotFound
	}
	err = s.repo.DeleteTask(id)

	if err != nil {
		return err
	}
	_ = s.logger.LogAction(ctx, "delete", "task", map[string]interface{}{
		"id":    id,
		"title": t.Title,
	})
	return nil
}


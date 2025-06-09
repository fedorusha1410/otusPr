package service

import (
	"context"
	"errors"
	"task-manager/internal/model/task"
	"task-manager/internal/repository"
	"time"
)

type Service struct {
	repo repository.TaskRepository
}

var (
	ErrNotFound = errors.New("task not found")
	ErrInvalid  = errors.New("invalid task input")
)

func New(repo repository.TaskRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTask(ctx context.Context, t *task.Task) (*task.Task, error) {
	if t == nil || t.Title == "" {
		return nil, ErrInvalid
	}

	t.CreatedTime = time.Now()
	t.UpdatedTime = t.CreatedTime

	s.repo.Save(*t)
	return t, nil
}

func (s *Service) GetTaskByID(ctx context.Context, id int) (*task.Task, error) {
	t := s.repo.GetTaskById(id)
	if t == nil {
		return nil, ErrNotFound
	}
	return t, nil
}

func (s *Service) GetTasks(ctx context.Context) ([]*task.Task, error) {
	return s.repo.GetTasks(), nil
}

func (s *Service) UpdateTask(ctx context.Context, id int, newData *task.Task) error {
	old := s.repo.GetTaskById(id)
	if old == nil {
		return ErrNotFound
	}

	newData.UpdatedTime = time.Now()
	s.repo.UpdateTask(id, newData)
	return nil
}

func (s *Service) DeleteTask(ctx context.Context, id int) error {
	t := s.repo.GetTaskById(id)
	if t == nil {
		return ErrNotFound
	}
	s.repo.DeleteTask(id)
	return nil
}

func (s *Service) SaveInFile() {
	s.repo.SaveTaskInFile()
}

func (s *Service) Restore() {
	s.repo.Restore()
}

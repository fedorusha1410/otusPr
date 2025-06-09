package userservice

import (
	"auth-service/internal/model/user"
	"auth-service/internal/repository"
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidUser  = errors.New("invalid user input")
)

type Service struct {
	repo repository.UserRepository
}

func New(repo repository.UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(u *user.User) (*user.User, error) {
	if u == nil || u.Name == "" {
		return nil, ErrInvalidUser
	}

	s.repo.Save(*u)
	return u, nil
}

func (s *Service) GetUserByID(id int) (*user.User, error) {
	u := s.repo.GetUserById(id)
	if u == nil {
		return nil, ErrUserNotFound
	}
	return u, nil
}

func (s *Service) GetUsers() ([]*user.User, error) {
	return s.repo.GetUsers(), nil
}

func (s *Service) UpdateUser(id int, data *user.User) error {
	existing := s.repo.GetUserById(id)
	if existing == nil {
		return ErrUserNotFound
	}
	s.repo.UpdateUser(id, data)
	return nil
}

func (s *Service) DeleteUser(id int) error {
	existing := s.repo.GetUserById(id)
	if existing == nil {
		return ErrUserNotFound
	}
	s.repo.DeleteUser(id)
	return nil
}

func (s *Service) SaveInFile() {
	s.repo.SaveUserInFile()
}

func (s *Service) Restore() {
	s.repo.Restore()
}

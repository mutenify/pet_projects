package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mutenify/taskflow/internal/repository"
	"github.com/mutenify/taskflow/internal/usecase"
)

type TaskRepository interface {
	Add(t repository.Task)
	Get(id string) (repository.Task, error)
}

type Service struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTask() string {
	id := uuid.New().String()
	s.repo.Add(repository.Task{ID: id, Status: usecase.StatusReady, Result: usecase.ResultDone})
	return id
}

func (s *Service) GetTask(id string) (repository.Task, error) {
	task, err := s.repo.Get(id)
	if err != nil {
		return repository.Task{}, fmt.Errorf("get task %s: %w", id, err)
	}
	return task, nil
}

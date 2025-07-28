package service

import (
	"task-api/internal/domain"
	"task-api/internal/port/inbound"
	"task-api/internal/port/outbound"
)

type taskService struct {
	repo outbound.TaskRepository
}

func NewTaskService(repo outbound.TaskRepository) inbound.TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) Create(task *domain.Task) error {
	return s.repo.Create(task)
}

func (s *taskService) GetByID(id int64) (*domain.Task, error) {
	return s.repo.GetByID(id)
}

func (s *taskService) GetAll() ([]*domain.Task, error) {
	return s.repo.GetAll()
}

func (s *taskService) Update(task *domain.Task) error {
	return s.repo.Update(task)
}

func (s *taskService) Delete(id int64) error {
	return s.repo.Delete(id)
}


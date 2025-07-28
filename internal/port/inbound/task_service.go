package inbound

import "task-api/internal/domain"

type TaskService interface {
	Create(task *domain.Task) error
	GetByID(id int64) (*domain.Task, error)
	GetAll() ([]*domain.Task, error)
	Update(task *domain.Task) error
	Delete(id int64) error
}


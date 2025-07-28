package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"task-api/internal/application"
	"task-api/internal/domain"
)

type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) Create(task *domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *mockRepo) GetByID(id int64) (*domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *mockRepo) GetAll() ([]*domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Task), args.Error(1)
}

func (m *mockRepo) Update(task *domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *mockRepo) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateTask(t *testing.T) {
	mockRepo := new(mockRepo)
	s := service.NewTaskService(mockRepo)

	task := &domain.Task{Title: "Test", Description: "Test Desc"}
	mockRepo.On("Create", task).Return(nil)

	err := s.Create(task)

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Create", task)
}

func TestGetByID(t *testing.T) {
	mockRepo := new(mockRepo)
	s := service.NewTaskService(mockRepo)

	expected := &domain.Task{ID: 1, Title: "Title", Description: "Desc"}
	mockRepo.On("GetByID", int64(1)).Return(expected, nil)

	result, err := s.GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestDeleteFailure(t *testing.T) {
	mockRepo := new(mockRepo)
	s := service.NewTaskService(mockRepo)

	mockRepo.On("Delete", int64(99)).Return(errors.New("not found"))

	err := s.Delete(99)

	assert.Error(t, err)
}


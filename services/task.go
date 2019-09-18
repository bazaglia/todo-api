package services

import (
	"todo/config"
	"todo/models"
)

type TaskService struct {
	config     *config.Config
	repository *models.TaskRepository
}

func (service *TaskService) FindAll(filter *models.TaskFilter) (*[]*models.Task, error) {
	return service.repository.FindAll(filter)
}

func (service *TaskService) Get(id string) (*models.Task, error) {
	return service.repository.Get(id)
}

func (service *TaskService) Create(data *models.Task) (*models.Task, error) {
	return service.repository.Create(data)
}

func (service *TaskService) Update(data *models.Task) (*models.Task, error) {
	return service.repository.Update(data)
}

func NewTaskService(config *config.Config, repository *models.TaskRepository) *TaskService {
	return &TaskService{config: config, repository: repository}
}

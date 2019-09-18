package services

import (
	"todo/config"
	"todo/models"
)

// UserService layer to user methods
type UserService struct {
	config     *config.Config
	repository *models.UserRepository
}

// Create Thin layer to create user's repository method
func (service *UserService) Create(data *models.User) (*models.User, error) {
	return service.repository.Create(data)
}

// Login Thin layer to login user's repository method
func (service *UserService) Login(email, password string) (*models.User, error) {
	return service.repository.Login(email, password)
}

// NewUserService initialize user service
func NewUserService(config *config.Config, repository *models.UserRepository) *UserService {
	return &UserService{config: config, repository: repository}
}

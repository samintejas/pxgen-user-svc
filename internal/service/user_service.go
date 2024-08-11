package service

import (
	"pxgen.io/user/internal/models"
	"pxgen.io/user/internal/repo"
)

type UserServiceInterface interface {
	ListUsers() ([]models.User, error)
	GetUserByID(id string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(id string, user *models.User) error
	DeleteUser(id string) error
}

type UserService struct {
	repo repo.UserRepositoryInterface
}

func NewUserService(repo repo.UserRepositoryInterface) UserServiceInterface {
	return &UserService{repo: repo}
}

func (s *UserService) ListUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) UpdateUser(id string, user *models.User) error {
	return s.repo.UpdateUser(id, user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}

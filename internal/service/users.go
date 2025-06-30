package service

import (
	"bus_depot/internal/models"
	"bus_depot/internal/repository"
	"bus_depot/utils"
	"errors"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
	// Хэшируем пароль с помощью SHA-256
	user.Password = utils.GenerateHash(user.Password)
	return s.repo.CreateUser(user)
}


func (s *UserService) AuthenticateUser(email, password string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("пользователь не найден")
	}

	hashedInput := utils.GenerateHash(password)
	if user.Password != hashedInput {
		return nil, errors.New("неверный пароль")
	}

	return user, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}

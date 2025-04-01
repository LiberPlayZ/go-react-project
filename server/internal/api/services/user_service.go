package services

import (
	"server/internal/db/repositories"
	"server/internal/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.UserRepo.GetAllUsers()
}

func (s *UserService) CreateUser(user models.User) error {
	user.ID = uuid.New()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashPassword)

	return s.UserRepo.CreateUser(user)

}

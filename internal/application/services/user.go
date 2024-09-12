package services

import (
	entities "gowhisper/internal/domain/entitites"
	"gowhisper/internal/domain/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(username string) (*entities.User, error) {
	newUser, _ := entities.NewUser(username)
	return newUser, s.userRepo.Save(newUser)
}

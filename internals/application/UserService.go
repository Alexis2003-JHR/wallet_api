package application

import (
	"wallet/internals/core/domain"
	"wallet/internals/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(user domain.User) error {
	return s.userRepo.Create(&user)
}

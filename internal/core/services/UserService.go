package services

import (
	"wallet/internal/core/domain"
	repository "wallet/internal/repository/user"
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

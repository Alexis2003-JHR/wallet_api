package user

import "wallet/internal/core/domain"

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id uint) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id uint) error
	GetAll() ([]*domain.User, error)
}

package postgres

import (
	"wallet/internals/core/domain"
	"wallet/internals/repository"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func NewGormRepository(db *gorm.DB) repository.UserRepository {
	return &GormUserRepository{DB: db}
}

func (r *GormUserRepository) Create(user *domain.User) error {
	return r.DB.Create(user).Error
}

func (r *GormUserRepository) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	result := r.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *GormUserRepository) Update(user *domain.User) error {
	return r.DB.Save(user).Error
}

func (r *GormUserRepository) Delete(id uint) error {
	return r.DB.Delete(&domain.User{}, id).Error
}

func (r *GormUserRepository) GetAll() ([]*domain.User, error) {
	var users []*domain.User
	result := r.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

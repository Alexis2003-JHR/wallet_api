package users

import (
	"wallet/internals/repository"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func NewGormRepository(db *gorm.DB) repository.UserRepository {
	return &GormUserRepository{DB: db}
}

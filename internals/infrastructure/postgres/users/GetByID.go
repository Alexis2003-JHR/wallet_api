package users

import "wallet/internals/core/domain"

func (r *GormUserRepository) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	result := r.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

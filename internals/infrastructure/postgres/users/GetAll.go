package users

import "wallet/internals/core/domain"

func (r *GormUserRepository) GetAll() ([]*domain.User, error) {
	var users []*domain.User
	result := r.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

package users

import "wallet/internals/core/domain"

func (r *GormUserRepository) Create(user *domain.User) error {
	return r.DB.Create(user).Error
}

package users

import "wallet/internals/core/domain"

func (r *GormUserRepository) Update(user *domain.User) error {
	return r.DB.Save(user).Error
}

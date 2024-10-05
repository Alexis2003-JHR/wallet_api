package users

import "wallet/internals/core/domain"

func (r *GormUserRepository) Delete(id uint) error {
	return r.DB.Delete(&domain.User{}, id).Error
}

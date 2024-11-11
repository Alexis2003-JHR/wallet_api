package services

import (
	"wallet/internal/core/domain"

	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) EncryptPassword(user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

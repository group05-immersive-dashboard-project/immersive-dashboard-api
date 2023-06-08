package repository

import (
	models "alta-immersive-dashboard/features"
)

type UserRepository interface {
	Insert(user models.UserEntity) (uint, error)
	Select(userID uint) (models.UserEntity, error)
	SelectAll() ([]models.UserEntity, error)
	Update(userID uint, updatedUser models.UserEntity) error
	Delete(userID uint) error
	Login(email, password string) (models.UserEntity, string, error)
}

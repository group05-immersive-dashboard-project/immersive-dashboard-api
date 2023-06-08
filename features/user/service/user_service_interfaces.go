package service

import (
	models "alta-immersive-dashboard/features"
)

type UserService interface {
	CreateUser(user models.UserEntity) (uint, error)
	GetUser(userID uint) (models.UserEntity, error)
	GetAllUser() ([]models.UserEntity, error)
	UpdateUser(userID uint, updatedUser models.UserEntity) error
	DeleteUser(userID uint) error
	Login(email, password string) (models.UserEntity, string, error)
}

package service

import (
	userEntity "alta-immersive-dashboard/features/user/repository"
)

type UserService interface {
	CreateUser(user userEntity.UserEntity) (uint, error)
	GetUser(userID uint) (userEntity.UserEntity, error)
	GetAllUser() ([]userEntity.UserEntity, error)
	UpdateUser(userID uint, updatedUser userEntity.UserEntity) error
	DeleteUser(userID uint) error
	Login(email, password string) (userEntity.UserEntity, string, error)
}

package service

import (
	userRepo "alta-immersive-dashboard/features/user/repository"
)

type UserService interface {
	CreateUser(user userRepo.UserEntity) (uint, error)
	GetUser(userID uint) (userRepo.UserEntity, error)
	GetAllUser() ([]userRepo.UserEntity, error)
	UpdateUser(userID uint, updatedUser userRepo.UserEntity) error
	DeleteUser(userID uint) error
	Login(email, password string) (userRepo.UserEntity, string, error)
}

package service

import (
	"alta-immersive-dashboard/features"
	models "alta-immersive-dashboard/features"
	userRepo "alta-immersive-dashboard/features/user/repository"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userRepository userRepo.UserRepository
	validate       *validator.Validate
}

// CreateUser implements UserService.
func (us *userService) CreateUser(user features.UserEntity) (uint, error) {
	err := us.validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, e := range validationErrors {
			switch e.Field() {
			case "TeamID":
				return 0, errors.New("error, team id is required")
			case "FullName":
				return 0, errors.New("error, name is required")
			case "Email":
				return 0, errors.New("error, invalid email format")
			case "Password":
				return 0, errors.New("error, password is required")
			}
		}
	}

	userID, err := us.userRepository.Insert(user)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	return userID, nil
}

// DeleteUser implements UserService.
func (us *userService) DeleteUser(userID uint) error {
	panic("unimplemented")
}

// GetAllUser implements UserService.
func (us *userService) GetAllUser() ([]features.UserEntity, error) {
	panic("unimplemented")
}

// GetUser implements UserService.
func (us *userService) GetUser(userID uint) (features.UserEntity, error) {
	userEntity, err := us.userRepository.Select(userID)
	if err != nil {
		return models.UserEntity{}, fmt.Errorf("error: %v", err)
	}
	return userEntity, nil
}

// Login implements UserService.
func (us *userService) Login(email string, password string) (features.UserEntity, string, error) {
	panic("unimplemented")
}

// UpdateUser implements UserService.
func (us *userService) UpdateUser(userID uint, updatedUser features.UserEntity) error {
	panic("unimplemented")
}

func New(repo userRepo.UserRepository) UserService {
	return &userService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

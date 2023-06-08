package service

import (
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
func (us *userService) CreateUser(user userRepo.UserEntity) (uint, error) {
	err := us.validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, e := range validationErrors {
			switch e.Field() {
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
	err := us.userRepository.Delete(userID)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

// GetAllUser implements UserService.
func (us *userService) GetAllUser() ([]userRepo.UserEntity, error) {
	userEntities, err := us.userRepository.SelectAll()
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	return userEntities, nil
}

// GetUser implements UserService.
func (us *userService) GetUser(userID uint) (userRepo.UserEntity, error) {
	userEntity, err := us.userRepository.Select(userID)
	if err != nil {
		return userRepo.UserEntity{}, fmt.Errorf("error: %v", err)
	}
	return userEntity, nil
}

// Login implements UserService.
func (us *userService) Login(email string, password string) (userRepo.UserEntity, string, error) {
	if email == "" {
		return userRepo.UserEntity{}, "", errors.New("email is required")
	} else if password == "" {
		return userRepo.UserEntity{}, "", errors.New("password is required")
	}

	loggedInUser, accessToken, err := us.userRepository.Login(email, password)
	if err != nil {
		return userRepo.UserEntity{}, "", err
	}

	return loggedInUser, accessToken, nil
}

// UpdateUser implements UserService.
func (us *userService) UpdateUser(userID uint, updatedUser userRepo.UserEntity) error {
	err := us.userRepository.Update(userID, updatedUser)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func New(repo userRepo.UserRepository) UserService {
	return &userService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

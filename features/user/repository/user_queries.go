package repository

import (
	"alta-immersive-dashboard/app/middlewares"
	"alta-immersive-dashboard/utils"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

// Delete implements UserRepository.
func (uq *userQuery) Delete(userID uint) error {
	panic("unimplemented")
}

// Insert implements UserRepository.
func (uq *userQuery) Insert(user UserEntity) (uint, error) {
	userModel := EntityToModel(user)
	userModel.Password = utils.HashPass(userModel.Password)

	createOpr := uq.db.Create(&userModel)
	if createOpr.Error != nil {
		if strings.Contains(createOpr.Error.Error(), "email") {
			return 0, errors.New("email already in use")
		} else if strings.Contains(createOpr.Error.Error(), "phone") {
			return 0, errors.New("phone already in use")
		}
		return 0, createOpr.Error
	}

	if createOpr.RowsAffected == 0 {
		return 0, errors.New("failed to insert, row affected is 0")
	}

	return userModel.ID, nil
}

// Login implements UserRepository.
func (uq *userQuery) Login(email string, password string) (UserEntity, string, error) {
	var user User

	queryResult := uq.db.Where("email = ?", email).First(&user)
	if queryResult.Error != nil {
		return UserEntity{}, "", errors.New(queryResult.Error.Error() + ", invalid email")
	}
	if queryResult.RowsAffected == 0 {
		return UserEntity{}, "", errors.New("login failed, invalid email")
	}

	err := utils.ComparePass([]byte(user.Password), []byte(password))
	if err != nil {
		return UserEntity{}, "", errors.New("login failed, invalid password")
	}

	accessToken, err := middlewares.GenerateToken(int(user.ID), user.Role)
	if err != nil {
		return UserEntity{}, "", err
	}
	userEntity := ModelToEntity(user)

	return userEntity, accessToken, nil
}

// Select implements UserRepository.
func (uq *userQuery) Select(userID uint) (UserEntity, error) {
	var user User

	queryResult := uq.db.Preload("Classes").Preload("Feedbacks").First(&user, userID)
	if queryResult.Error != nil {
		return UserEntity{}, queryResult.Error
	}

	userEntity := ModelToEntity(user)

	return userEntity, nil
}

// SelectAll implements UserRepository.
func (uq *userQuery) SelectAll(userID uint) ([]UserEntity, error) {
	panic("unimplemented")
}

// Update implements UserRepository.
func (uq *userQuery) Update(userID uint, updatedUser UserEntity) error {
	panic("unimplemented")
}

func New(db *gorm.DB) UserRepository {
	return &userQuery{
		db: db,
	}
}

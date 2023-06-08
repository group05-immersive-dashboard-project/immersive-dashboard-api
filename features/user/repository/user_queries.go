package repository

import (
	"alta-immersive-dashboard/app/middlewares"
	convert "alta-immersive-dashboard/features"
	models "alta-immersive-dashboard/features"
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
func (uq *userQuery) Insert(user models.UserEntity) (uint, error) {
	userModel := convert.UserEntityToModel(user)
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
func (uq *userQuery) Login(email string, password string) (models.UserEntity, string, error) {
	var user models.User

	queryResult := uq.db.Where("email = ?", email).First(&user)
	if queryResult.Error != nil {
		return models.UserEntity{}, "", errors.New(queryResult.Error.Error() + ", invalid email")
	}
	if queryResult.RowsAffected == 0 {
		return models.UserEntity{}, "", errors.New("login failed, invalid email")
	}

	err := utils.ComparePass([]byte(user.Password), []byte(password))
	if err != nil {
		return models.UserEntity{}, "", errors.New("login failed, invalid password")
	}

	accessToken, err := middlewares.GenerateToken(int(user.ID), user.Role)
	if err != nil {
		return models.UserEntity{}, "", err
	}
	userEntity := convert.UserModelToEntity(user)

	return userEntity, accessToken, nil
}

// Select implements UserRepository.
func (uq *userQuery) Select(userID uint) (models.UserEntity, error) {
	var user models.User

	queryResult := uq.db.Preload("Classes").Preload("Feedbacks").First(&user, userID)
	if queryResult.Error != nil {
		return models.UserEntity{}, queryResult.Error
	}

	userEntity := convert.UserModelToEntity(user)

	return userEntity, nil
}

// SelectAll implements UserRepository.
func (uq *userQuery) SelectAll() ([]models.UserEntity, error) {
	panic("unimplemented")
}

// Update implements UserRepository.
func (uq *userQuery) Update(userID uint, updatedUser models.UserEntity) error {
	panic("unimplemented")
}

func New(db *gorm.DB) UserRepository {
	return &userQuery{
		db: db,
	}
}

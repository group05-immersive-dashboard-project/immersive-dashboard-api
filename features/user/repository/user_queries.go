package repository

import (
	"alta-immersive-dashboard/app/middlewares"
	"alta-immersive-dashboard/utils"
	"errors"

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
func (uq *userQuery) Insert(user UserEntity) error {
	panic("unimplemented")
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
	panic("unimplemented")
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

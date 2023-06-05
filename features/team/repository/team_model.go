package repository

import (
	"gorm.io/gorm"

	userModel "alta-immersive-dashboard/features/user/repository"
)

type Team struct {
	gorm.Model
	TeamName string           `gorm:"team_name;not null"`
	Users    []userModel.User `gorm:"foreignKey:TeamID"`
}

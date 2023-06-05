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

func EntityToModel(team TeamEntity) Team {
	// Convert user entities to user models
	var userModels []userModel.User
	for _, user := range team.Users {
		userModels = append(userModels, userModel.EntityToModel(user))
	}

	return Team{
		TeamName: team.TeamName,
		Users:    userModels,
	}
}

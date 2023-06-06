package repository

import (
	classModel "alta-immersive-dashboard/features/class/repository"
	feedbackModel "alta-immersive-dashboard/features/feedback/repository"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TeamID    uint                     `gorm:"column:team_id;not null"`
	FullName  string                   `gorm:"column:full_name;not null"`
	Email     string                   `gorm:"column:email;not null"`
	Password  string                   `gorm:"column:password;not null"`
	Status    string                   `gorm:"type:enum('active','inactive', 'deleted');default:'active';column:status;not null"`
	Role      string                   `gorm:"type:enum('admin','user');default:'user';column:role;not null"`
	Classes   []classModel.Class       `gorm:"foreignKey:UserID"`
	Feedbacks []feedbackModel.Feedback `gorm:"foreignKey:UserID"`
}

func EntityToModel(user UserEntity) User {
	// Convert class entities to class models
	var classModels []classModel.Class
	for _, class := range user.Classes {
		classModels = append(classModels, classModel.EntityToModel(class))
	}

	// Convert feedback entities to feedback models
	var feedbackModels []feedbackModel.Feedback
	for _, feedback := range user.Feedbacks {
		feedbackModels = append(feedbackModels, feedbackModel.EntityToModel(feedback))
	}

	return User{
		TeamID:    user.TeamID,
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  user.Password,
		Status:    user.Status,
		Role:      user.Role,
		Classes:   classModels,
		Feedbacks: feedbackModels,
	}
}

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
	Status    string                   `gorm:"type:enum('active','inactive');default:'active';column:status;not null"`
	Role      string                   `gorm:"type:enum('admin','user');default:'user';column:role;not null"`
	Classes   []classModel.Class       `gorm:"foreignKey:UserID;column:classes"`
	Feedbacks []feedbackModel.Feedback `gorm:"foreignKey:UserID;column:feedbacks"`
}

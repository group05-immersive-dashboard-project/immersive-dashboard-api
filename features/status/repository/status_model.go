package repository

import (
	feedbackModel "alta-immersive-dashboard/features/feedback/repository"
	menteeModel "alta-immersive-dashboard/features/mentee/repository"

	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	StatusName string                   `gorm:"status_name;not null"`
	Mentees    []menteeModel.Mentee     `gorm:"foreignKey:StatusID"`
	Feedbacks  []feedbackModel.Feedback `gorm:"foreignKey:StatusID"`
}

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

func EntityToModel(status StatusEntity) Status {
	// Convert mentee entities to mentee models
	var menteeModels []menteeModel.Mentee
	for _, mentee := range status.Mentees {
		menteeModels = append(menteeModels, menteeModel.EntityToModel(mentee))
	}

	// Convert feedback entities to feedback models
	var feedbackModels []feedbackModel.Feedback
	for _, feedback := range status.Feedbacks {
		feedbackModels = append(feedbackModels, feedbackModel.EntityToModel(feedback))
	}

	return Status{
		StatusName: status.StatusName,
		Mentees:    menteeModels,
		Feedbacks:  feedbackModels,
	}
}

package repository

import (
	feedbackEntity "alta-immersive-dashboard/features/feedback/repository"
	menteeEntity "alta-immersive-dashboard/features/mentee/repository"
	"time"
)

type StatusEntity struct {
	ID         uint                            `json:"status_id,omitempty" form:"status_id"`
	StatusName string                          `gorm:"status_name;not null"`
	CreatedAt  time.Time                       `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt  time.Time                       `json:"updated_at,omitempty" form:"updated_at"`
	DeletedAt  time.Time                       `json:"deleted_at,omitempty" form:"deleted_at"`
	Mentees    []menteeEntity.MenteeEntity     `json:"mentees,omitempty"`
	Feedbacks  []feedbackEntity.FeedbackEntity `json:"feedbacks,omitempty"`
}

func ModelToEntity(status Status) StatusEntity {
	// Convert mentee models to mentee entities
	var menteeEntities []menteeEntity.MenteeEntity
	for _, mentee := range status.Mentees {
		menteeEntities = append(menteeEntities, menteeEntity.ModelToEntity(mentee))
	}

	// Convert feedback models to Feedback entities
	var feedbackEntities []feedbackEntity.FeedbackEntity
	for _, feedback := range status.Feedbacks {
		feedbackEntities = append(feedbackEntities, feedbackEntity.ModelToEntity(feedback))
	}

	return StatusEntity{
		ID:         status.ID,
		StatusName: status.StatusName,
		CreatedAt:  status.CreatedAt,
		UpdatedAt:  status.UpdatedAt,
		DeletedAt:  status.DeletedAt.Time,
		Mentees:    menteeEntities,
		Feedbacks:  feedbackEntities,
	}
}

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

package repository

import (
	feedbackEntity "alta-immersive-dashboard/features/feedback/repository"
	"time"
)

type MenteeEntity struct {
	ID              uint                            `json:"mentee_id,omitempty" form:"mentee_id"`
	StatusID        uint                            `json:"status_id,omitempty" form:"status_id"`
	ClassID         uint                            `json:"class_id,omitempty" form:"class_id"`
	FullName        string                          `json:"full_name,omitempty" form:"full_name"`
	NickName        string                          `json:"nick_name,omitempty" form:"nick_name"`
	Email           string                          `json:"email,omitempty" form:"email"`
	Phone           string                          `json:"phone,omitempty" form:"phone"`
	CurrentAddress  string                          `json:"current_address,omitempty" form:"current_address"`
	HomeAddress     string                          `json:"home_address,omitempty" form:"home_address"`
	Telegram        string                          `json:"telegram,omitempty" form:"telegram"`
	Gender          string                          `json:"gender,omitempty" form:"gender"`
	EducationType   string                          `json:"education_type,omitempty" form:"education_type"`
	Major           string                          `json:"major,omitempty" form:"major"`
	Graduate        string                          `json:"graduate,omitempty" form:"graduate"`
	Institution     string                          `json:"institution,omitempty" form:"institution"`
	EmergencyName   string                          `json:"emergency_name,omitempty" form:"emergency_name"`
	EmergencyPhone  string                          `json:"emergency_phone,omitempty" form:"emergency_phone"`
	EmergencyStatus string                          `json:"emergency_status,omitempty" form:"emergency_status"`
	CreatedAt       time.Time                       `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt       time.Time                       `json:"updated_at,omitempty" form:"updated_at"`
	DeletedAt       time.Time                       `json:"deleted_at,omitempty" form:"deleted_at"`
	Feedbacks       []feedbackEntity.FeedbackEntity `json:"feedbacks,omitempty"`
}

func ModelToEntity(mentee Mentee) MenteeEntity {
	// Convert feedback models to Feedback entities
	var feedbackEntities []feedbackEntity.FeedbackEntity
	for _, feedback := range mentee.Feedbacks {
		feedbackEntities = append(feedbackEntities, feedbackEntity.ModelToEntity(feedback))
	}

	return MenteeEntity{
		ID:              mentee.ID,
		StatusID:        mentee.StatusID,
		ClassID:         mentee.ClassID,
		FullName:        mentee.FullName,
		NickName:        mentee.NickName,
		Email:           mentee.Email,
		Phone:           mentee.Phone,
		CurrentAddress:  mentee.CurrentAddress,
		HomeAddress:     mentee.HomeAddress,
		Telegram:        mentee.Telegram,
		Gender:          mentee.Gender,
		EducationType:   mentee.EducationType,
		Major:           mentee.Major,
		Graduate:        mentee.Graduate,
		Institution:     mentee.Institution,
		EmergencyName:   mentee.EmergencyName,
		EmergencyPhone:  mentee.EmergencyPhone,
		EmergencyStatus: mentee.EmergencyStatus,
		CreatedAt:       mentee.CreatedAt,
		UpdatedAt:       mentee.UpdatedAt,
		DeletedAt:       mentee.DeletedAt.Time,
		Feedbacks:       feedbackEntities,
	}
}

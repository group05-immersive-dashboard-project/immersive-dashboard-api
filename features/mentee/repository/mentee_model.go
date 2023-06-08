package repository

import (
	feedbackModel "alta-immersive-dashboard/features/feedback/repository"

	"gorm.io/gorm"
)

type Mentee struct {
	gorm.Model
	StatusID        uint                     `gorm:"column:status_id;not null"`
	ClassID         uint                     `gorm:"column:class_id;not null"`
	FullName        string                   `gorm:"column:full_name;not null"`
	NickName        string                   `gorm:"column:nick_name"`
	Email           string                   `gorm:"column:email;unique"`
	Phone           string                   `gorm:"column:phone;unique"`
	CurrentAddress  string                   `gorm:"column:current_address"`
	HomeAddress     string                   `gorm:"column:home_address"`
	Telegram        string                   `gorm:"column:telegram;unique"`
	Gender          string                   `gorm:"type:enum('male','female');default:'male';column:gender;not null"`
	EducationType   string                   `gorm:"type:enum('informatics','non-informatics');default:'informatics';column:education_type;not null"`
	Major           string                   `gorm:"column:major"`
	Graduate        string                   `gorm:"column:graduate"`
	Institution     string                   `gorm:"column:institution"`
	EmergencyName   string                   `gorm:"column:emergency_name"`
	EmergencyPhone  string                   `gorm:"column:emergency_phone;unique"`
	EmergencyStatus string                   `gorm:"type:enum('parent','grandparents', 'siblings');default:'parent';column:emergency_status;not null"`
	Feedbacks       []feedbackModel.Feedback `gorm:"foreignKey:MenteeID"`
}

func EntityToModel(mentee MenteeEntity) Mentee {
	// Convert feedback entities to feedback models
	var feedbackModels []feedbackModel.Feedback
	for _, feedback := range mentee.Feedbacks {
		feedbackModels = append(feedbackModels, feedbackModel.EntityToModel(feedback))
	}

	return Mentee{
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
		Feedbacks:       feedbackModels,
	}
}

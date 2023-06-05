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
	Telegram        string                   `gorm:"column:telegram"`
	Discord         string                   `gorm:"column:discord"`
	Gender          string                   `gorm:"type:enum('male','female');default:'male';column:gender;not null"`
	EducationType   string                   `gorm:"type:enum('informatics','non-informatics');default:'informatics';column:education_type;not null"`
	Major           string                   `gorm:"column:major"`
	Graduate        string                   `gorm:"column:graduate"`
	Institution     string                   `gorm:"column:institution"`
	EmergencyName   string                   `gorm:"column:emergency_name"`
	EmergencyPhone  string                   `gorm:"column:emergency_phone;unique"`
	EmergencyStatus string                   `gorm:"type:enum('parent','grandparents', 'parents brother');default:'parent';column:emergency_status;not null"`
	Feedbacks       []feedbackModel.Feedback `gorm:"foreignKey:MenteeID"`
}

package controllers

import (
	menteeRepo "alta-immersive-dashboard/features/mentee/repository"
)

type CreateDeleteMenteeResponse struct {
	ID            uint   `json:"mentee_id" form:"mentee_id"`
	ClassID       uint   `json:"class_id" form:"class_id"`
	StatusID      uint   `json:"status_id" form:"status_id"`
	FullName      string `json:"full_name" form:"full_name"`
	Gender        string `json:"gender" form:"gender"`
	EducationType string `json:"education_type" form:"education_type"`
}

type ReadUpdateMenteeResponse struct {
	ID              uint   `json:"mentee_id,omitempty" form:"mentee_id"`
	StatusID        uint   `json:"status_id,omitempty" form:"status_id"`
	ClassID         uint   `json:"class_id,omitempty" form:"class_id"`
	FullName        string `json:"full_name,omitempty" form:"full_name"`
	NickName        string `json:"nick_name,omitempty" form:"nick_name"`
	Email           string `json:"email,omitempty" form:"email"`
	Phone           string `json:"phone,omitempty" form:"phone"`
	CurrentAddress  string `json:"current_address,omitempty" form:"current_address"`
	HomeAddress     string `json:"home_address,omitempty" form:"home_address"`
	Telegram        string `json:"telegram,omitempty" form:"telegram"`
	Gender          string `json:"gender,omitempty" form:"gender"`
	EducationType   string `json:"education_type,omitempty" form:"education_type"`
	Major           string `json:"major,omitempty" form:"major"`
	Graduate        string `json:"graduate,omitempty" form:"graduate"`
	Institution     string `json:"institution,omitempty" form:"institution"`
	EmergencyName   string `json:"emergency_name,omitempty" form:"emergency_name"`
	EmergencyPhone  string `json:"emergency_phone,omitempty" form:"emergency_phone"`
	EmergencyStatus string `json:"emergency_status,omitempty" form:"emergency_status"`
}

func EntityToCreateDeleteMenteeResponse(mentee menteeRepo.MenteeEntity) CreateDeleteMenteeResponse {
	return CreateDeleteMenteeResponse{
		ID:            mentee.ID,
		ClassID:       mentee.ClassID,
		StatusID:      mentee.StatusID,
		FullName:      mentee.FullName,
		Gender:        mentee.Gender,
		EducationType: mentee.EducationType,
	}
}

func EntityToReadUpdateMenteeResponse(mentee menteeRepo.MenteeEntity) ReadUpdateMenteeResponse {
	return ReadUpdateMenteeResponse{
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
	}
}

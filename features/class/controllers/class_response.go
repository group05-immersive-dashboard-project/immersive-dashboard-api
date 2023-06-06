package controllers

import (
	classRepo "alta-immersive-dashboard/features/class/repository"
)

// Response for POST and DELETE class
type CreateDeleteClassResponse struct {
	ID     uint   `json:"class_id" form:"class_id"`
	UserID uint   `json:"user_id" form:"user_id"`
	Name   string `json:"class_name" form:"class_name"`
}

// Response for GET and PUT class
type GetUpdateClassResponse struct {
	ID           uint   `json:"class_id" form:"class_id"`
	UserID       uint   `json:"user_id" form:"user_id"`
	Name         string `json:"class_name" form:"class_name"`
	StartDate    string `json:"start_date" form:"graduate_date"`
	GraduateDate string `json:"graduate_date" form:"graduate_date"`
}

func EntityToCreateDeleteClassResponse(class classRepo.ClassEntity) CreateDeleteClassResponse {
	return CreateDeleteClassResponse{
		ID:     class.ID,
		UserID: class.UserID,
		Name:   class.Name,
	}
}

func EntityToGetUpdateClassResponse(class classRepo.ClassEntity) GetUpdateClassResponse {
	return GetUpdateClassResponse{
		ID:           class.ID,
		UserID:       class.UserID,
		Name:         class.Name,
		StartDate:    class.StartDate,
		GraduateDate: class.GraduateDate,
	}
}

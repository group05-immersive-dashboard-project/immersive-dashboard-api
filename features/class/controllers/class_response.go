package controllers

import (
	classRepo "alta-immersive-dashboard/features/class/repository"
)

type ClassResponse struct {
	ID     uint   `json:"class_id" form:"class_id"`
	UserID uint   `json:"user_id" form:"user_id"`
	Name   string `json:"class_name" form:"class_name"`
	// StartDate    string `json:"start_date" form:"graduate_date"`
	// GraduateDate string `json:"graduate_date" form:"graduate_date"`
}

func EntityToClassResponse(class classRepo.ClassEntity) ClassResponse {
	return ClassResponse{
		ID:     class.ID,
		UserID: class.UserID,
		Name:   class.Name,
	}
}

package controllers

import (
	userRepo "alta-immersive-dashboard/features/user/repository"
)

type ReadUserResponse struct {
	ID       uint   `json:"user_id" form:"user_id"`
	TeamID   uint   `json:"team_id" form:"team_id"`
	FullName string `json:"full_name" form:"full_name"`
	Email    string `json:"email" form:"email"`
	Role     string `json:"role" form:"role"`
	Status   string `json:"status" form:"status"`
}

func EntityToReadUserResponse(user userRepo.UserEntity) ReadUserResponse {
	return ReadUserResponse{
		ID:       user.ID,
		TeamID:   user.TeamID,
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
		Status:   user.Status,
	}
}

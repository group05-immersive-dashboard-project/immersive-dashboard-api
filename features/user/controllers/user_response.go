package controllers

import (
	models "alta-immersive-dashboard/features"
)

type UserResponse struct {
	ID       uint `json:"user_id" form:"user_id"`
	TeamID   uint `json:"team_id" form:"team_id"`
	Team     models.Team
	FullName string `json:"full_name" form:"full_name"`
	Email    string `json:"email" form:"email"`
	Role     string `json:"role" form:"role"`
	Status   string `json:"status" form:"status"`
}

func EntityToUserResponse(user models.UserEntity) UserResponse {
	return UserResponse{
		ID:       user.ID,
		TeamID:   user.TeamID,
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
		Status:   user.Status,
	}
}

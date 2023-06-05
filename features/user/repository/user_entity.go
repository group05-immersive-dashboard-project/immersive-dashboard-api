package repository

import (
	classEntity "alta-immersive-dashboard/features/class/repository"
	feedbackEntity "alta-immersive-dashboard/features/feedback/repository"
	"time"
)

type UserEntity struct {
	ID        uint                            `json:"user_id,omitempty" form:"user_id"`
	TeamID    uint                            `json:"team_id,omitempty" form:"team_id"`
	FullName  string                          `json:"full_name,omitempty" form:"full_name"`
	Email     string                          `json:"email,omitempty" form:"email"`
	Password  string                          `json:"password,omitempty" form:"password"`
	Status    string                          `json:"status,omitempty" form:"status"`
	Role      string                          `json:"role,omitempty" form:"role"`
	CreatedAt time.Time                       `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt time.Time                       `json:"updated_at,omitempty" form:"updated_at"`
	DeletedAt time.Time                       `json:"deleted_at,omitempty" form:"deleted_at"`
	Classes   []classEntity.ClassEntity       `json:"classes,omitempty"`
	Feedbacks []feedbackEntity.FeedbackEntity `json:"feedbacks,omitempty"`
}

func ModelToEntity(user User) UserEntity {
	// Convert Classes model to ClassEntities
	var classEntities []classEntity.ClassEntity
	for _, class := range user.Classes {
		classEntities = append(classEntities, classEntity.ModelToEntity(class))
	}

	// Convert feedback models to Feedback entities
	var feedbackEntities []feedbackEntity.FeedbackEntity
	for _, feedback := range user.Feedbacks {
		feedbackEntities = append(feedbackEntities, feedbackEntity.ModelToEntity(feedback))
	}

	return UserEntity{
		ID:        user.ID,
		TeamID:    user.TeamID,
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  user.Password,
		Status:    user.Status,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt.Time,
		Classes:   classEntities,
		Feedbacks: feedbackEntities,
	}
}

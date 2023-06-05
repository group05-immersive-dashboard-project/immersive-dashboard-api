package repository

import (
	menteeEntity "alta-immersive-dashboard/features/mentee/repository"
	"time"
)

type ClassEntity struct {
	ID           uint                        `json:"class_id,omitempty" form:"class_id"`
	UserID       uint                        `json:"user_id,omitempty" form:"user_id"`
	Name         string                      `json:"class_name,omitempty" form:"class_name"`
	StartDate    string                      `json:"start_date,omitempty" form:"start_date"`
	GraduateDate string                      `json:"graduate_date,omitempty" form:"graduate_date"`
	CreatedAt    time.Time                   `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt    time.Time                   `json:"updated_at,omitempty" form:"updated_at"`
	DeletedAt    time.Time                   `json:"deleted_at,omitempty" form:"deleted_at"`
	Mentees      []menteeEntity.MenteeEntity `json:"mentees,omitempty"`
}

func ModelToEntity(class Class) ClassEntity {
	// Convert mentee models to mentee entities
	var menteeEntities []menteeEntity.MenteeEntity
	for _, mentee := range class.Mentees {
		menteeEntities = append(menteeEntities, menteeEntity.ModelToEntity(mentee))
	}

	return ClassEntity{
		ID:           class.ID,
		UserID:       class.UserID,
		Name:         class.Name,
		StartDate:    class.StartDate,
		GraduateDate: class.GraduateDate,
		CreatedAt:    class.CreatedAt,
		UpdatedAt:    class.UpdatedAt,
		DeletedAt:    class.DeletedAt.Time,
		Mentees:      menteeEntities,
	}
}

package repository

import (
	menteeModel "alta-immersive-dashboard/features/mentee/repository"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	UserID       uint                 `gorm:"column:user_id;not null"`
	Name         string               `gorm:"column:class_name;not null"`
	StartDate    string               `gorm:"column:start_date;not null"`
	GraduateDate string               `gorm:"column:graduate_date;not null"`
	Mentees      []menteeModel.Mentee `gorm:"foreignKey:ClassID"`
}

func EntityToModel(class ClassEntity) Class {
	// Convert mentee entities to mentee models
	var menteeModels []menteeModel.Mentee
	for _, mentee := range class.Mentees {
		menteeModels = append(menteeModels, menteeModel.EntityToModel(mentee))
	}

	return Class{
		UserID:       class.UserID,
		Name:         class.Name,
		StartDate:    class.StartDate,
		GraduateDate: class.GraduateDate,
		Mentees:      menteeModels,
	}
}

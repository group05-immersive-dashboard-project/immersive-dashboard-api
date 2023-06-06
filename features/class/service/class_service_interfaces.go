package service

import (
	classEntity "alta-immersive-dashboard/features/class/repository"
)

type ClassService interface {
	CreateClass(class classEntity.ClassEntity) (uint, error)
	GetClass(classID uint) (classEntity.ClassEntity, error)
	GetAllClass() ([]classEntity.ClassEntity, error)
	UpdateClass(classID uint, updatedClass classEntity.ClassEntity) error
	DeleteClass(classID uint) error
}

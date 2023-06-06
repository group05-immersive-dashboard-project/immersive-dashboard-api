package service

import (
	classRepo "alta-immersive-dashboard/features/class/repository"
	"errors"
	"fmt"
)

type classService struct {
	classRepository classRepo.ClassRepository
}

// CreateClass implements ClassService.
func (cs *classService) CreateClass(class classRepo.ClassEntity) (uint, error) {
	if class.UserID == 0 {
		return 0, errors.New("error, user id is required")
	}
	if class.Name == "" {
		return 0, errors.New("error, class name is required")
	}
	if class.StartDate == "" {
		return 0, errors.New("error, start date is required")
	}
	if class.GraduateDate == "" {
		return 0, errors.New("error, graduate date is required")
	}

	classID, err := cs.classRepository.Insert(class)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	return classID, nil
}

// DeleteClass implements ClassService.
func (cs *classService) DeleteClass(classID uint) error {
	err := cs.classRepository.Delete(classID)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

// GetAllClass implements ClassService.
func (cs *classService) GetAllClass() ([]classRepo.ClassEntity, error) {
	panic("unimplemented")
}

// GetClass implements ClassService.
func (cs *classService) GetClass(classID uint) (classRepo.ClassEntity, error) {
	panic("unimplemented")
}

// UpdateClass implements ClassService.
func (cs *classService) UpdateClass(classID uint, updatedClass classRepo.ClassEntity) error {
	panic("unimplemented")
}

func New(repo classRepo.ClassRepository) ClassService {
	return &classService{
		classRepository: repo,
	}
}

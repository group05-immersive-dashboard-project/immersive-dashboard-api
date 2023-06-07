package service

import (
	menteeRepo "alta-immersive-dashboard/features/mentee/repository"
	"errors"
	"fmt"
)

type menteeService struct {
	menteeRepository menteeRepo.MenteeRepository
}

// CreateMentee implements MenteeService.
func (ms *menteeService) CreateMentee(mentee menteeRepo.MenteeEntity) (uint, error) {
	switch {
	case mentee.StatusID == 0:
		return 0, errors.New("error, status is required")
	case mentee.ClassID == 0:
		return 0, errors.New("error, class is required")
	case mentee.FullName == "":
		return 0, errors.New("error, name is required")
	case mentee.NickName == "":
		return 0, errors.New("error, nick name is required")
	case mentee.Email == "":
		return 0, errors.New("error, email is required")
	case mentee.Phone == "":
		return 0, errors.New("error, phone is required")
	case mentee.CurrentAddress == "":
		return 0, errors.New("error, current address is required")
	case mentee.HomeAddress == "":
		return 0, errors.New("error, home address is required")
	case mentee.Telegram == "":
		return 0, errors.New("error, telegram is required")
	case mentee.Gender == "":
		return 0, errors.New("error, gender is required")
	case mentee.EducationType == "":
		return 0, errors.New("error, education type is required")
	case mentee.Major == "":
		return 0, errors.New("error, major is required")
	case mentee.Graduate == "":
		return 0, errors.New("error, graduate is required")
	case mentee.Institution == "":
		return 0, errors.New("error, institution is required")
	case mentee.EmergencyName == "":
		return 0, errors.New("error, emergency name is required")
	case mentee.EmergencyPhone == "":
		return 0, errors.New("error, emergency phone is required")
	case mentee.EmergencyStatus == "":
		return 0, errors.New("error, emergency status is required")
	}

	menteeID, err := ms.menteeRepository.Insert(mentee)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	return menteeID, nil
}

// DeleteMentee implements MenteeService.
func (ms *menteeService) DeleteMentee(menteeID uint) error {
	err := ms.menteeRepository.Delete(menteeID)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

// GetAllMentee implements MenteeService.
func (ms *menteeService) GetAllMentee() ([]menteeRepo.MenteeEntity, error) {
	menteeEntities, err := ms.menteeRepository.SelectAll()
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	return menteeEntities, nil
}

// GetAllMenteeByFilters implements MenteeService.
func (ms *menteeService) GetAllMenteeByFilters(filters menteeRepo.MenteeFilter) ([]menteeRepo.MenteeEntity, error) {
	menteeEntities, err := ms.menteeRepository.SelectAllByFilters(filters)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	return menteeEntities, nil
}

// GetMentee implements MenteeService.
func (ms *menteeService) GetMentee(menteeID uint) (menteeRepo.MenteeEntity, error) {
	menteeEntity, err := ms.menteeRepository.Select(menteeID)
	if err != nil {
		return menteeRepo.MenteeEntity{}, fmt.Errorf("error: %v", err)
	}

	return menteeEntity, nil
}

// UpdateMentee implements MenteeService.
func (ms *menteeService) UpdateMentee(menteeID uint, updatedMentee menteeRepo.MenteeEntity) error {
	err := ms.menteeRepository.Update(menteeID, updatedMentee)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func New(repo menteeRepo.MenteeRepository) MenteeService {
	return &menteeService{
		menteeRepository: repo,
	}
}

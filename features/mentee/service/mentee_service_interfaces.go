package service

import (
	menteeRepo "alta-immersive-dashboard/features/mentee/repository"
)

type MenteeService interface {
	CreateMentee(mentee menteeRepo.MenteeEntity) (uint, error)
	GetMentee(menteeID uint) (menteeRepo.MenteeEntity, error)
	GetAllMentee() ([]menteeRepo.MenteeEntity, error)
	UpdateMentee(menteeID uint, updatedMentee menteeRepo.MenteeEntity) error
	DeleteMentee(menteeID uint) error
}

package repository

import (
	"errors"

	"gorm.io/gorm"
)

type menteeQuery struct {
	db *gorm.DB
}

// Delete implements MenteeRepository.
func (mq *menteeQuery) Delete(menteeID uint) error {
	panic("unimplemented")
}

// Insert implements MenteeRepository.
func (mq *menteeQuery) Insert(mentee MenteeEntity) (uint, error) {
	menteeModel := EntityToModel(mentee)

	createOpr := mq.db.Create(&menteeModel)
	if createOpr.Error != nil {
		return 0, createOpr.Error
	}

	if createOpr.RowsAffected == 0 {
		return 0, errors.New("failed to insert, row affected is 0")
	}

	return menteeModel.ID, nil
}

// Select implements MenteeRepository.
func (mq *menteeQuery) Select(menteeID uint) (MenteeEntity, error) {
	var mentee Mentee

	queryResult := mq.db.Preload("Feedbacks").First(&mentee, menteeID)
	if queryResult.Error != nil {
		return MenteeEntity{}, queryResult.Error
	}

	menteeEntity := ModelToEntity(mentee)

	return menteeEntity, nil
}

// SelectAll implements MenteeRepository.
func (mq *menteeQuery) SelectAll() ([]MenteeEntity, error) {
	panic("unimplemented")
}

// Update implements MenteeRepository.
func (mq *menteeQuery) Update(menteeID uint, updatedMentee MenteeEntity) error {
	panic("unimplemented")
}

func New(db *gorm.DB) MenteeRepository {
	return &menteeQuery{
		db: db,
	}
}

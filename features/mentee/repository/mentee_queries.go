package repository

import (
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type menteeQuery struct {
	db *gorm.DB
}

// Delete implements MenteeRepository.
func (mq *menteeQuery) Delete(menteeID uint) error {
	deleteOpr := mq.db.Delete(&Mentee{}, menteeID)
	if deleteOpr.Error != nil {
		return errors.New(deleteOpr.Error.Error() + ", failed to delete mentee")
	}

	return nil
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
	var mentees []Mentee

	queryResult := mq.db.Preload("Feedbacks").Find(&mentees)
	if queryResult.Error != nil {
		return []MenteeEntity{}, queryResult.Error
	}

	var menteeEntities []MenteeEntity
	for _, mentee := range mentees {
		claasEntity := ModelToEntity(mentee)
		menteeEntities = append(menteeEntities, claasEntity)
	}

	return menteeEntities, nil
}

func convertClassID(ID string) (uint, error) {
	classID, err := strconv.Atoi(ID)
	if err != nil {
		return 0, errors.New("invalid class id")
	}
	return uint(classID), nil
}

func convertStatusID(ID string) (uint, error) {
	statusID, err := strconv.Atoi(ID)
	if err != nil {
		return 0, err
	}
	return uint(statusID), errors.New("invalid status id")
}

func (mq *menteeQuery) SelectAllByFilters(filters MenteeFilter) ([]MenteeEntity, error) {
	var mentees []Mentee

	query := mq.db.Preload("Feedbacks")

	switch {
	case filters.ClassID == "" && filters.StatusID == "":
		query = query.Where("education_type = ?", filters.Category)
	case filters.ClassID == "" && filters.Category == "":
		statusID, _ := convertStatusID(filters.StatusID)
		query = query.Where("status_id = ?", statusID)
	case filters.StatusID == "" && filters.Category == "":
		classID, _ := convertClassID(filters.ClassID)
		query = query.Where("class_id = ?", classID)
	case filters.ClassID == "":
		statusID, _ := convertStatusID(filters.StatusID)
		query = query.Where("status_id = ? AND education_type = ?", statusID, filters.Category)
	case filters.StatusID == "":
		classID, _ := convertClassID(filters.ClassID)
		query = query.Where("class_id = ? AND education_type = ?", classID, filters.Category)
	case filters.Category == "":
		statusID, _ := convertStatusID(filters.StatusID)
		classID, _ := convertClassID(filters.ClassID)
		query = query.Where("class_id = ? AND status_id = ?", statusID, classID)
	}

	queryResult := query.Find(&mentees)
	if queryResult.Error != nil {
		return []MenteeEntity{}, queryResult.Error
	}

	var menteeEntities []MenteeEntity
	for _, mentee := range mentees {
		claasEntity := ModelToEntity(mentee)
		menteeEntities = append(menteeEntities, claasEntity)
	}

	return menteeEntities, nil
}

// Update implements MenteeRepository.
func (mq *menteeQuery) Update(menteeID uint, updatedMentee MenteeEntity) error {
	var mentee Mentee

	queryResult := mq.db.First(&mentee, menteeID)
	if queryResult.Error != nil {
		return errors.New(queryResult.Error.Error() + ", failed to get mentee")
	}

	menteeToUpdate := EntityToModel(updatedMentee)
	updateOpr := mq.db.Model(&mentee).Updates(menteeToUpdate)
	if updateOpr.Error != nil {
		return errors.New(updateOpr.Error.Error() + ", failed to update mentee")
	}

	return nil
}

func New(db *gorm.DB) MenteeRepository {
	return &menteeQuery{
		db: db,
	}
}

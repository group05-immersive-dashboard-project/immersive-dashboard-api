package repository

import (
	"errors"

	"gorm.io/gorm"
)

type feedbackQuery struct {
	db *gorm.DB
}

// Select implements FeedbackRepository.
func (fq *feedbackQuery) Select(feedbackID uint) (FeedbackEntity, error) {
	var feedback Feedback

	queryResult := fq.db.First(&feedback, feedbackID)
	if queryResult.Error != nil {
		return FeedbackEntity{}, queryResult.Error
	}

	feedbackEntity := ModelToEntity(feedback)

	return feedbackEntity, nil
}

// Delete implements FeedbackRepository.
func (fq *feedbackQuery) Delete(feedbackID uint) error {
	deleteOpr := fq.db.Delete(&Feedback{}, feedbackID)
	if deleteOpr.Error != nil {
		return errors.New(deleteOpr.Error.Error() + ", failed to delete feedback")
	}

	return nil
}

// Insert implements FeedbackRepository.
func (fq *feedbackQuery) Insert(feedback FeedbackEntity) (uint, error) {
	feedbackModel := EntityToModel(feedback)

	createOpr := fq.db.Create(&feedbackModel)
	if createOpr.Error != nil {
		return 0, createOpr.Error
	}

	if createOpr.RowsAffected == 0 {
		return 0, errors.New("failed to insert, row affected is 0")
	}

	return feedbackModel.ID, nil
}

// Update implements FeedbackRepository.
func (fq *feedbackQuery) Update(feedbackID uint, updatedFeedback FeedbackEntity) error {
	var feedback Feedback

	queryResult := fq.db.First(&feedback, feedbackID)
	if queryResult.Error != nil {
		return errors.New(queryResult.Error.Error() + ", failed to get feedback")
	}

	feedbackToUpdate := EntityToModel(updatedFeedback)
	updateOpr := fq.db.Model(&feedback).Updates(feedbackToUpdate)
	if updateOpr.Error != nil {
		return errors.New(updateOpr.Error.Error() + ", failed to update feedback")
	}

	return nil
}

func New(db *gorm.DB) FeedbackRepository {
	return &feedbackQuery{
		db: db,
	}
}

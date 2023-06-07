package service

import (
	feedbackRepo "alta-immersive-dashboard/features/feedback/repository"
)

type FeedbackService interface {
	CreateFeedback(feedback feedbackRepo.FeedbackEntity) (uint, error)
	UpdateFeedback(feedbackID uint, updatedFeedback feedbackRepo.FeedbackEntity) error
	DeleteFeedback(feedbackID uint) error
}

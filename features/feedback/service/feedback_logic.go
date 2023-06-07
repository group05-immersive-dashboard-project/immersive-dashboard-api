package service

import (
	feedbackRepo "alta-immersive-dashboard/features/feedback/repository"
	"errors"
	"fmt"
)

type feedbackService struct {
	feedbackRepository feedbackRepo.FeedbackRepository
}

// GetFeedback implements FeedbackService.
func (fs *feedbackService) GetFeedback(feedbackID uint) (feedbackRepo.FeedbackEntity, error) {
	feedbackEntity, err := fs.feedbackRepository.Select(feedbackID)
	if err != nil {
		return feedbackRepo.FeedbackEntity{}, fmt.Errorf("error: %v", err)
	}

	return feedbackEntity, nil
}

// CreateFeedback implements FeedbackService.
func (fs *feedbackService) CreateFeedback(feedback feedbackRepo.FeedbackEntity) (uint, error) {
	switch {
	case feedback.MenteeID == 0:
		return 0, errors.New("error, mentee ID is required")
	case feedback.StatusID == 0:
		return 0, errors.New("error, status ID is required")
	case feedback.UserID == 0:
		return 0, errors.New("error, user ID is required")
	case feedback.Notes == "":
		return 0, errors.New("error, notes is required")
	case feedback.Proof == "":
		return 0, errors.New("error, proof is required")
	}

	feedbackID, err := fs.feedbackRepository.Insert(feedback)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	return feedbackID, nil
}

// DeleteFeedback implements FeedbackService.
func (fs *feedbackService) DeleteFeedback(feedbackID uint) error {
	err := fs.feedbackRepository.Delete(feedbackID)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

// UpdateFeedback implements FeedbackService.
func (fs *feedbackService) UpdateFeedback(feedbackID uint, updatedFeedback feedbackRepo.FeedbackEntity) error {
	err := fs.feedbackRepository.Update(feedbackID, updatedFeedback)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func New(repo feedbackRepo.FeedbackRepository) FeedbackService {
	return &feedbackService{
		feedbackRepository: repo,
	}
}

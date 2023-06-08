package controllers

import (
	feedbackRepo "alta-immersive-dashboard/features/feedback/repository"
)

type FeedbackResponse struct {
	ID       uint   `json:"feedback_id" form:"feedback_id"`
	MenteeID uint   `json:"mentee_id" form:"mentee_id"`
	StatusID uint   `json:"status_id" form:"status_id"`
	UserID   uint   `json:"user_id" form:"user_id"`
	Notes    string `json:"notes" form:"notes"`
	Proof    string `json:"proof" form:"proof"`
}

func EntityToFeedbackResponse(feedback feedbackRepo.FeedbackEntity) FeedbackResponse {
	return FeedbackResponse{
		ID:       feedback.ID,
		MenteeID: feedback.MenteeID,
		StatusID: feedback.StatusID,
		UserID:   feedback.UserID,
		Notes:    feedback.Notes,
		Proof:    feedback.Proof,
	}
}

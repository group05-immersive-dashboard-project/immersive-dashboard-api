package repository

import "time"

type FeedbackEntity struct {
	ID        uint      `json:"feedback_id,omitempty" form:"feedback_id"`
	MenteeID  uint      `json:"mentee_id,omitempty" form:"mentee_id"`
	StatusID  uint      `json:"status_id,omitempty" form:"status_id"`
	UserID    uint      `json:"user_id,omitempty" form:"user_id"`
	Notes     string    `json:"notes,omitempty" form:"notes"`
	Proof     string    `json:"proof,omitempty" form:"proof"`
	CreatedAt time.Time `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" form:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty" form:"deleted_at"`
}

func ModelToEntity(feedback Feedback) FeedbackEntity {
	return FeedbackEntity{
		ID:        feedback.ID,
		MenteeID:  feedback.MenteeID,
		StatusID:  feedback.StatusID,
		UserID:    feedback.UserID,
		Notes:     feedback.Notes,
		Proof:     feedback.Proof,
		CreatedAt: feedback.CreatedAt,
		UpdatedAt: feedback.UpdatedAt,
		DeletedAt: feedback.DeletedAt.Time,
	}
}

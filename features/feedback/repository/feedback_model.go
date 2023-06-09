package repository

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	MenteeID uint   `gorm:"column:mentee_id;not null"`
	StatusID uint   `gorm:"column:status_id;not null"`
	UserID   uint   `gorm:"column:user_id;not null"`
	Notes    string `gorm:"column:notes;not null"`
	Proof    string `gorm:"column:proofs;not null"`
}

func EntityToModel(feedback FeedbackEntity) Feedback {
	return Feedback{
		MenteeID: feedback.MenteeID,
		StatusID: feedback.StatusID,
		UserID:   feedback.UserID,
		Notes:    feedback.Notes,
		Proof:    feedback.Proof,
	}
}

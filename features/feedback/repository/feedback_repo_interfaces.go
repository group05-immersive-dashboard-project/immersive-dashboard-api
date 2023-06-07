package repository

type FeedbackRepository interface {
	Insert(feedback FeedbackEntity) (uint, error)
	Update(feedbackID uint, updatedFeedback FeedbackEntity) error
	Delete(feedbackID uint) error
}

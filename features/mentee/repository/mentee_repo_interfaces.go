package repository

type MenteeRepository interface {
	Insert(mentee MenteeEntity) (uint, error)
	Select(menteeID uint) (MenteeEntity, error)
	SelectAll() ([]MenteeEntity, error)
	Update(menteeID uint, updatedMentee MenteeEntity) error
	Delete(menteeID uint) error
}

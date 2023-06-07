package repository

type MenteeFilter struct {
	ClassID  string
	StatusID string
	Category string
}

func (mf MenteeFilter) IsEmpty() bool {
	return mf.ClassID == "" && mf.StatusID == "" && mf.Category == ""
}

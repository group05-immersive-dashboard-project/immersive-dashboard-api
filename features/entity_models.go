package features

import "time"

type UserEntity struct {
	ID        uint             `json:"user_id,omitempty" form:"user_id"`
	TeamID    uint             `json:"team_id,omitempty" form:"team_id"`
	Team      TeamEntity       `json:"team,omitempty"`
	FullName  string           `json:"full_name,omitempty" form:"full_name" validate:"required"`
	Email     string           `json:"email,omitempty" form:"email" validate:"required,email"`
	Password  string           `json:"password,omitempty" form:"password" validate:"required,min=8"`
	Status    string           `json:"status,omitempty" form:"status"`
	Role      string           `json:"role,omitempty" form:"role"`
	CreatedAt time.Time        `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt time.Time        `json:"updated_at,omitempty" form:"updated_at"`
	DeletedAt time.Time        `json:"deleted_at,omitempty" form:"deleted_at"`
	Classes   []ClassEntity    `json:"classes,omitempty"`
	Feedbacks []FeedbackEntity `json:"feedbacks,omitempty"`
}

type ClassEntity struct {
	ID           uint           `json:"class_id,omitempty" form:"class_id"`
	UserID       uint           `json:"user_id,omitempty" form:"user_id"`
	User         UserEntity     `json:"users,omitempty"`
	Name         string         `json:"class_name,omitempty" form:"class_name"`
	StartDate    string         `json:"start_date,omitempty" form:"start_date"`
	GraduateDate string         `json:"graduate_date,omitempty" form:"graduate_date"`
	CreatedAt    time.Time      `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at,omitempty" form:"updated_at"`
	DeletedAt    time.Time      `json:"deleted_at,omitempty" form:"deleted_at"`
	Mentees      []MenteeEntity `json:"mentees,omitempty"`
}

type TeamEntity struct {
	ID        uint         `json:"team_id,omitempty" form:"team_id"`
	TeamName  string       `json:"team_name,omitempty" form:"team_name"`
	CreatedAt time.Time    `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt time.Time    `json:"updated_at,omitempty" form:"updated_at"`
	DeletedAt time.Time    `json:"deleted_at,omitempty" form:"deleted_at"`
	Users     []UserEntity `json:"users,omitempty"`
}

type StatusEntity struct {
	ID         uint             `json:"status_id,omitempty" form:"status_id"`
	StatusName string           `gorm:"status_name;not null"`
	CreatedAt  time.Time        `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at,omitempty" form:"updated_at"`
	DeletedAt  time.Time        `json:"deleted_at,omitempty" form:"deleted_at"`
	Mentees    []MenteeEntity   `json:"mentees,omitempty"`
	Feedbacks  []FeedbackEntity `json:"feedbacks,omitempty"`
}

type FeedbackEntity struct {
	ID        uint         `json:"feedback_id,omitempty" form:"feedback_id"`
	MenteeID  uint         `json:"mentee_id,omitempty" form:"mentee_id"`
	Mentee    MenteeEntity `json:"mentees,omitempty"`
	StatusID  uint         `json:"status_id,omitempty" form:"status_id"`
	Status    StatusEntity `json:"statuses,omitempty"`
	UserID    uint         `json:"user_id,omitempty" form:"user_id"`
	User      UserEntity   `json:"users,omitempty"`
	Notes     string       `json:"notes,omitempty" form:"notes"`
	Proof     string       `json:"proof,omitempty" form:"proof"`
	CreatedAt time.Time    `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt time.Time    `json:"updated_at,omitempty" form:"updated_at"`
	DeletedAt time.Time    `json:"deleted_at,omitempty" form:"deleted_at"`
}

type MenteeEntity struct {
	ID              uint             `json:"mentee_id,omitempty" form:"mentee_id"`
	StatusID        uint             `json:"status_id,omitempty" form:"status_id"`
	Status          StatusEntity     `json:"statuses,omitempty"`
	ClassID         uint             `json:"class_id,omitempty" form:"class_id"`
	Class           ClassEntity      `json:"classes,omitempty"`
	FullName        string           `json:"full_name,omitempty" form:"full_name"`
	NickName        string           `json:"nick_name,omitempty" form:"nick_name"`
	Email           string           `json:"email,omitempty" form:"email"`
	Phone           string           `json:"phone,omitempty" form:"phone"`
	CurrentAddress  string           `json:"current_address,omitempty" form:"current_address"`
	HomeAddress     string           `json:"home_address,omitempty" form:"home_address"`
	Telegram        string           `json:"telegram,omitempty" form:"telegram"`
	Gender          string           `json:"gender,omitempty" form:"gender"`
	EducationType   string           `json:"education_type,omitempty" form:"education_type"`
	Major           string           `json:"major,omitempty" form:"major"`
	Graduate        string           `json:"graduate,omitempty" form:"graduate"`
	Institution     string           `json:"institution,omitempty" form:"institution"`
	EmergencyName   string           `json:"emergency_name,omitempty" form:"emergency_name"`
	EmergencyPhone  string           `json:"emergency_phone,omitempty" form:"emergency_phone"`
	EmergencyStatus string           `json:"emergency_status,omitempty" form:"emergency_status"`
	CreatedAt       time.Time        `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at,omitempty" form:"updated_at"`
	DeletedAt       time.Time        `json:"deleted_at,omitempty" form:"deleted_at"`
	Feedbacks       []FeedbackEntity `json:"feedbacks,omitempty"`
}

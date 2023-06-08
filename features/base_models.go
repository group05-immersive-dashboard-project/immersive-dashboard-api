package features

import "gorm.io/gorm"

type User struct {
	gorm.Model
	TeamID    uint       `gorm:"column:team_id;not null"`
	Team      Team       `gorm:"foreignKey:TeamID"`
	FullName  string     `gorm:"column:full_name;not null"`
	Email     string     `gorm:"column:email;unique;not null"`
	Password  string     `gorm:"column:password;not null"`
	Status    string     `gorm:"type:enum('active','inactive', 'deleted');default:'active';column:status;not null"`
	Role      string     `gorm:"type:enum('admin','user');default:'user';column:role;not null"`
	Classes   []Class    `gorm:"foreignKey:UserID"`
	Feedbacks []Feedback `gorm:"foreignKey:UserID"`
}

// class belongs to user
type Class struct {
	gorm.Model
	UserID       uint     `gorm:"column:user_id;not null"`
	User         User     `gorm:"foreignKey:UserID"`
	Name         string   `gorm:"column:class_name;unique;not null"`
	StartDate    string   `gorm:"column:start_date;not null"`
	GraduateDate string   `gorm:"column:graduate_date;not null"`
	Mentees      []Mentee `gorm:"foreignKey:ClassID"`
}

type Team struct {
	gorm.Model
	TeamName string `gorm:"team_name;unique;not null" json:"team_name"`
	Users    []User `gorm:"foreignKey:TeamID"`
}

type Status struct {
	gorm.Model
	StatusName string     `gorm:"status_name;unique;not null"`
	Mentees    []Mentee   `gorm:"foreignKey:StatusID"`
	Feedbacks  []Feedback `gorm:"foreignKey:StatusID"`
}

type Feedback struct {
	gorm.Model
	MenteeID uint   `gorm:"column:mentee_id;not null"`
	Mentee   Mentee `gorm:"foreignKey:MenteeID"`
	StatusID uint   `gorm:"column:status_id;not null"`
	Status   Status `gorm:"foreignKey:StatusID"`
	UserID   uint   `gorm:"column:user_id;not null"`
	User     User   `gorm:"foreignKey:UserID"`
	Notes    string `gorm:"column:notes;not null"`
	Proof    string `gorm:"column:proofs;not null"`
}

type Mentee struct {
	gorm.Model
	StatusID        uint       `gorm:"column:status_id;not null"`
	Status          Status     `gorm:"foreignKey:StatusID"`
	ClassID         uint       `gorm:"column:class_id;not null"`
	Class           Class      `gorm:"foreignKey:ClassID"`
	FullName        string     `gorm:"column:full_name;not null"`
	NickName        string     `gorm:"column:nick_name"`
	Email           string     `gorm:"column:email;unique"`
	Phone           string     `gorm:"column:phone;unique"`
	CurrentAddress  string     `gorm:"column:current_address"`
	HomeAddress     string     `gorm:"column:home_address"`
	Telegram        string     `gorm:"column:telegram;unique"`
	Gender          string     `gorm:"type:enum('male','female');default:'male';column:gender;not null"`
	EducationType   string     `gorm:"type:enum('informatics','non-informatics');default:'informatics';column:education_type;not null"`
	Major           string     `gorm:"column:major"`
	Graduate        string     `gorm:"column:graduate"`
	Institution     string     `gorm:"column:institution"`
	EmergencyName   string     `gorm:"column:emergency_name"`
	EmergencyPhone  string     `gorm:"column:emergency_phone;unique"`
	EmergencyStatus string     `gorm:"type:enum('parent','grandparents', 'siblings');default:'parent';column:emergency_status;not null"`
	Feedbacks       []Feedback `gorm:"foreignKey:MenteeID"`
}

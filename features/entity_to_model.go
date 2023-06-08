package features

func UserEntityToModel(user UserEntity) User {
	return User{
		TeamID:   user.TeamID,
		Team:     TeamEntityToModel(user.Team),
		FullName: user.FullName,
		Email:    user.Email,
		Password: user.Password,
		Status:   user.Status,
		Role:     user.Role,
	}
}

func ClassEntityToModel(class ClassEntity) Class {
	return Class{
		UserID:       class.UserID,
		User:         UserEntityToModel(class.User),
		Name:         class.Name,
		StartDate:    class.StartDate,
		GraduateDate: class.GraduateDate,
	}
}

func TeamEntityToModel(team TeamEntity) Team {
	return Team{
		TeamName: team.TeamName,
	}
}

func StatusEntityToModel(status StatusEntity) Status {
	return Status{
		StatusName: status.StatusName,
	}
}

func FeedbackEntityToModel(feedback FeedbackEntity) Feedback {
	return Feedback{
		MenteeID: feedback.MenteeID,
		Mentee:   MenteeEntityToModel(feedback.Mentee),
		StatusID: feedback.StatusID,
		Status:   StatusEntityToModel(feedback.Status),
		UserID:   feedback.UserID,
		User:     UserEntityToModel(feedback.User),
		Notes:    feedback.Notes,
		Proof:    feedback.Proof,
	}
}

func MenteeEntityToModel(mentee MenteeEntity) Mentee {
	return Mentee{
		StatusID:        mentee.StatusID,
		Status:          StatusEntityToModel(mentee.Status),
		ClassID:         mentee.ClassID,
		Class:           ClassEntityToModel(mentee.Class),
		FullName:        mentee.FullName,
		NickName:        mentee.NickName,
		Email:           mentee.Email,
		Phone:           mentee.Phone,
		CurrentAddress:  mentee.CurrentAddress,
		HomeAddress:     mentee.HomeAddress,
		Telegram:        mentee.Telegram,
		Gender:          mentee.Gender,
		EducationType:   mentee.EducationType,
		Major:           mentee.Major,
		Graduate:        mentee.Graduate,
		Institution:     mentee.Institution,
		EmergencyName:   mentee.EmergencyName,
		EmergencyPhone:  mentee.EmergencyPhone,
		EmergencyStatus: mentee.EmergencyStatus,
	}
}

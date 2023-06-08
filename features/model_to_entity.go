package features

func UserModelToEntity(user User) UserEntity {
	var classEntities []ClassEntity
	for _, class := range user.Classes {
		classEntities = append(classEntities, ClassModelToEntity(class))
	}

	// Convert feedback models to Feedback entities
	var feedbackEntities []FeedbackEntity
	for _, feedback := range user.Feedbacks {
		feedbackEntities = append(feedbackEntities, FeedbackModelToEntity(feedback))
	}

	return UserEntity{
		ID:        user.ID,
		TeamID:    user.TeamID,
		Team:      TeamModelToEntity(user.Team),
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  user.Password,
		Status:    user.Status,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt.Time,
		Classes:   classEntities,
		Feedbacks: feedbackEntities,
	}
}

func ClassModelToEntity(class Class) ClassEntity {
	// Convert mentee models to mentee entities
	var menteeEntities []MenteeEntity
	for _, mentee := range class.Mentees {
		menteeEntities = append(menteeEntities, MenteeModelToEntity(mentee))
	}

	return ClassEntity{
		ID:           class.ID,
		UserID:       class.UserID,
		User:         UserModelToEntity(class.User),
		Name:         class.Name,
		StartDate:    class.StartDate,
		GraduateDate: class.GraduateDate,
		CreatedAt:    class.CreatedAt,
		UpdatedAt:    class.UpdatedAt,
		DeletedAt:    class.DeletedAt.Time,
		Mentees:      menteeEntities,
	}
}

func TeamModelToEntity(team Team) TeamEntity {
	// Convert user models to user entities
	var userEntities []UserEntity
	for _, user := range team.Users {
		userEntities = append(userEntities, UserModelToEntity(user))
	}

	return TeamEntity{
		ID:        team.ID,
		TeamName:  team.TeamName,
		CreatedAt: team.CreatedAt,
		UpdatedAt: team.UpdatedAt,
		DeletedAt: team.DeletedAt.Time,
		Users:     userEntities,
	}
}

func StatusModelToEntity(status Status) StatusEntity {
	// Convert mentee models to mentee entities
	var menteeEntities []MenteeEntity
	for _, mentee := range status.Mentees {
		menteeEntities = append(menteeEntities, MenteeModelToEntity(mentee))
	}

	// Convert feedback models to Feedback entities
	var feedbackEntities []FeedbackEntity
	for _, feedback := range status.Feedbacks {
		feedbackEntities = append(feedbackEntities, FeedbackModelToEntity(feedback))
	}

	return StatusEntity{
		ID:         status.ID,
		StatusName: status.StatusName,
		CreatedAt:  status.CreatedAt,
		UpdatedAt:  status.UpdatedAt,
		DeletedAt:  status.DeletedAt.Time,
		Mentees:    menteeEntities,
		Feedbacks:  feedbackEntities,
	}
}

func FeedbackModelToEntity(feedback Feedback) FeedbackEntity {
	return FeedbackEntity{
		ID:        feedback.ID,
		MenteeID:  feedback.MenteeID,
		Mentee:    MenteeModelToEntity(feedback.Mentee),
		StatusID:  feedback.StatusID,
		Status:    StatusModelToEntity(feedback.Status),
		UserID:    feedback.UserID,
		User:      UserModelToEntity(feedback.User),
		Notes:     feedback.Notes,
		Proof:     feedback.Proof,
		CreatedAt: feedback.CreatedAt,
		UpdatedAt: feedback.UpdatedAt,
		DeletedAt: feedback.DeletedAt.Time,
	}
}

func MenteeModelToEntity(mentee Mentee) MenteeEntity {
	// Convert feedback models to Feedback entities
	var feedbackEntities []FeedbackEntity
	for _, feedback := range mentee.Feedbacks {
		feedbackEntities = append(feedbackEntities, FeedbackModelToEntity(feedback))
	}

	return MenteeEntity{
		ID:              mentee.ID,
		StatusID:        mentee.StatusID,
		Status:          StatusModelToEntity(mentee.Status),
		ClassID:         mentee.ClassID,
		Class:           ClassModelToEntity(mentee.Class),
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
		CreatedAt:       mentee.CreatedAt,
		UpdatedAt:       mentee.UpdatedAt,
		DeletedAt:       mentee.DeletedAt.Time,
		Feedbacks:       feedbackEntities,
	}
}

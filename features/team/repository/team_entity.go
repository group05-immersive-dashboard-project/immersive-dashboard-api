package repository

// import (
// 	userEntity "alta-immersive-dashboard/features/user/repository"
// 	"time"
// )

// type TeamEntity struct {
// 	ID        uint                    `json:"team_id,omitempty" form:"team_id"`
// 	TeamName  string                  `json:"team_name,omitempty" form:"team_name"`
// 	CreatedAt time.Time               `json:"created_at,omitempty" form:"created_at"`
// 	UpdatedAt time.Time               `json:"updated_at,omitempty" form:"updated_at"`
// 	DeletedAt time.Time               `json:"deleted_at,omitempty" form:"deleted_at"`
// 	Users     []userEntity.UserEntity `json:"users,omitempty"`
// }

// func ModelToEntity(team Team) TeamEntity {
// 	// Convert user models to user entities
// 	var userEntities []userEntity.UserEntity
// 	for _, user := range team.Users {
// 		userEntities = append(userEntities, userEntity.ModelToEntity(user))
// 	}

// 	return TeamEntity{
// 		ID:        team.ID,
// 		TeamName:  team.TeamName,
// 		CreatedAt: team.CreatedAt,
// 		UpdatedAt: team.UpdatedAt,
// 		DeletedAt: team.DeletedAt.Time,
// 		Users:     userEntities,
// 	}
// }

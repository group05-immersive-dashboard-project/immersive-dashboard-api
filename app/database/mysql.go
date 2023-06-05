package database

import (
	"alta-immersive-dashboard/app/config"
	classModel "alta-immersive-dashboard/features/class/repository"
	feedbackModel "alta-immersive-dashboard/features/feedback/repository"
	menteeModel "alta-immersive-dashboard/features/mentee/repository"
	statusModel "alta-immersive-dashboard/features/status/repository"
	teamModel "alta-immersive-dashboard/features/team/repository"
	userModel "alta-immersive-dashboard/features/user/repository"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(cfg *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

func InitialMigration(DB *gorm.DB) {
	DB.AutoMigrate(&feedbackModel.Feedback{})
	DB.AutoMigrate(&menteeModel.Mentee{})
	DB.AutoMigrate(&statusModel.Status{})
	DB.AutoMigrate(&classModel.Class{})
	DB.AutoMigrate(&userModel.User{})
	DB.AutoMigrate(&teamModel.Team{})
}

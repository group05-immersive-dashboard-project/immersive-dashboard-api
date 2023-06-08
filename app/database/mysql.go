package database

import (
	"alta-immersive-dashboard/app/config"
	"alta-immersive-dashboard/utils"
	"fmt"

	models "alta-immersive-dashboard/features"

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
	DB.AutoMigrate(&models.Team{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Class{})
	DB.AutoMigrate(&models.Status{})
	DB.AutoMigrate(&models.Mentee{})
	DB.AutoMigrate(&models.Feedback{})
}

func InitialTeamData(db *gorm.DB) {
	teams := []models.Team{
		{TeamName: "Manager"},
		{TeamName: "Mentor"},
		{TeamName: "Placement Team"},
		{TeamName: "People Skill Team"},
	}

	for _, team := range teams {
		db.Create(&team)
	}
}

func InitialUserData(db *gorm.DB) {
	user := models.User{
		TeamID:   1,
		FullName: "John Doe",
		Email:    "johndoe@mail.com",
		Password: "qwerty",
		Status:   "active",
		Role:     "admin",
	}

	user.Password = utils.HashPass(user.Password)

	db.Create(&user)
}

func InitialStatusData(db *gorm.DB) {
	statuses := []models.Status{
		{StatusName: "Interview"},
		{StatusName: "Join Class"},
		{StatusName: "Unit 1"},
		{StatusName: "Unit 2"},
		{StatusName: "Unit 3"},
		{StatusName: "Repeat Unit 1"},
		{StatusName: "Repeat Unit 2"},
		{StatusName: "Repeat Unit 3"},
		{StatusName: "Placement"},
		{StatusName: "Eliminated"},
		{StatusName: "Graduated"},
	}

	for _, status := range statuses {
		db.Create(&status)
	}
}

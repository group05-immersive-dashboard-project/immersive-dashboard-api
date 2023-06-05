package router

import (
	usrCtrl "alta-immersive-dashboard/features/user/controllers"
	usrRepo "alta-immersive-dashboard/features/user/repository"
	usrSrv "alta-immersive-dashboard/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userRepo := usrRepo.New(db)
	userService := usrSrv.New(userRepo)
	userHandlerAPI := usrCtrl.New(userService)

	usersGroup := e.Group("/users")
	{
		usersGroup.POST("/login", userHandlerAPI.LoginUser)
	}
}

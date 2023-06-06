package router

import (
	"alta-immersive-dashboard/app/middlewares"
	clsCtrl "alta-immersive-dashboard/features/class/controllers"
	clsRepo "alta-immersive-dashboard/features/class/repository"
	clsSrv "alta-immersive-dashboard/features/class/service"
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
		usersGroup.POST("/admin", userHandlerAPI.CreateUser)
		usersGroup.POST("", userHandlerAPI.CreateUser, middlewares.JWTMiddlewareFunc(), middlewares.AdminAuth)
		usersGroup.GET("", userHandlerAPI.ReadAllUser, middlewares.JWTMiddlewareFunc())
		usersGroup.GET("/:user_id", userHandlerAPI.ReadUser, middlewares.JWTMiddlewareFunc())
		usersGroup.PUT("", userHandlerAPI.UpdateUser, middlewares.JWTMiddlewareFunc())
		usersGroup.PUT("/:user_id", userHandlerAPI.UpdateUser, middlewares.JWTMiddlewareFunc())
		usersGroup.DELETE("", userHandlerAPI.DeleteUser, middlewares.JWTMiddlewareFunc())
		usersGroup.DELETE("/:user_id", userHandlerAPI.DeleteUser, middlewares.JWTMiddlewareFunc())
	}

	classRepo := clsRepo.New(db)
	classService := clsSrv.New(classRepo)
	classHandlerAPI := clsCtrl.New(classService)

	classesGroup := e.Group("/classes")
	{
		classesGroup.POST("", classHandlerAPI.CreateClass, middlewares.JWTMiddlewareFunc())
	}
}

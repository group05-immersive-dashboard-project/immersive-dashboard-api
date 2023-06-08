package router

import (
	"alta-immersive-dashboard/app/middlewares"
	clsCtrl "alta-immersive-dashboard/features/class/controllers"
	clsRepo "alta-immersive-dashboard/features/class/repository"
	clsSrv "alta-immersive-dashboard/features/class/service"
	fbCtrl "alta-immersive-dashboard/features/feedback/controllers"
	fbRepo "alta-immersive-dashboard/features/feedback/repository"
	fbSrv "alta-immersive-dashboard/features/feedback/service"
	mntCtrl "alta-immersive-dashboard/features/mentee/controllers"
	mntRepo "alta-immersive-dashboard/features/mentee/repository"
	mntSrv "alta-immersive-dashboard/features/mentee/service"
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

	// e.POST("/login", userHandlerAPI.LoginUser)

	usersGroup := e.Group("/users")
	{
		usersGroup.POST("", userHandlerAPI.CreateUser)
		// usersGroup.POST("", userHandlerAPI.CreateUser, middlewares.JWTMiddlewareFunc(), middlewares.AdminAuth)
		// usersGroup.GET("", userHandlerAPI.ReadAllUser, middlewares.JWTMiddlewareFunc())
		// usersGroup.GET("/:user_id", userHandlerAPI.ReadUser, middlewares.JWTMiddlewareFunc())
		// usersGroup.PUT("", userHandlerAPI.UpdateUser, middlewares.JWTMiddlewareFunc())
		// usersGroup.PUT("/:user_id", userHandlerAPI.UpdateUser, middlewares.JWTMiddlewareFunc())
		// usersGroup.DELETE("", userHandlerAPI.DeleteUser, middlewares.JWTMiddlewareFunc())
		// usersGroup.DELETE("/:user_id", userHandlerAPI.DeleteUser, middlewares.JWTMiddlewareFunc())
	}

	classRepo := clsRepo.New(db)
	classService := clsSrv.New(classRepo)
	classHandlerAPI := clsCtrl.New(classService)

	classesGroup := e.Group("/classes")
	{
		classesGroup.POST("", classHandlerAPI.CreateClass, middlewares.JWTMiddlewareFunc())
		classesGroup.GET("", classHandlerAPI.ReadAllClass, middlewares.JWTMiddlewareFunc())
		classesGroup.GET("/:class_id", classHandlerAPI.ReadClass, middlewares.JWTMiddlewareFunc())
		classesGroup.PUT("/:class_id", classHandlerAPI.UpdateClass, middlewares.JWTMiddlewareFunc())
		classesGroup.DELETE("/:class_id", classHandlerAPI.DeleteClass, middlewares.JWTMiddlewareFunc())
	}

	menteeRepo := mntRepo.New(db)
	menteeService := mntSrv.New(menteeRepo)
	menteeHandlerAPI := mntCtrl.New(menteeService)

	menteesGroup := e.Group("/mentees")
	{
		menteesGroup.POST("", menteeHandlerAPI.CreateMentee, middlewares.JWTMiddlewareFunc())
		menteesGroup.GET("", menteeHandlerAPI.ReadAllMentee, middlewares.JWTMiddlewareFunc())
		menteesGroup.GET("/:mentee_id", menteeHandlerAPI.ReadMentee, middlewares.JWTMiddlewareFunc())
		menteesGroup.GET("/:mentee_id/feedbacks", menteeHandlerAPI.ReadMenteeFeedbacks, middlewares.JWTMiddlewareFunc())
		menteesGroup.PUT("/:mentee_id", menteeHandlerAPI.UpdateMentee, middlewares.JWTMiddlewareFunc())
		menteesGroup.DELETE("/:mentee_id", menteeHandlerAPI.DeleteMentee, middlewares.JWTMiddlewareFunc())
	}

	feedbackRepo := fbRepo.New(db)
	feedbackService := fbSrv.New(feedbackRepo)
	feedbackHandlerAPI := fbCtrl.New(feedbackService)

	feedbacksGroup := e.Group("/feedbacks")
	{
		feedbacksGroup.POST("", feedbackHandlerAPI.CreateFeedback, middlewares.JWTMiddlewareFunc())
		feedbacksGroup.PUT("/:feedbacks_id", feedbackHandlerAPI.UpdateFeedback, middlewares.JWTMiddlewareFunc())
		feedbacksGroup.DELETE("/:feedbacks_id", feedbackHandlerAPI.DeleteFeedback, middlewares.JWTMiddlewareFunc())
	}
}

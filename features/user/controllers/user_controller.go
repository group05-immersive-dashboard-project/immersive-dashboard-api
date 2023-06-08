package controllers

import (
	models "alta-immersive-dashboard/features"
	userSrv "alta-immersive-dashboard/features/user/service"
	"alta-immersive-dashboard/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type userController struct {
	userService userSrv.UserService
}

func New(service userSrv.UserService) *userController {
	return &userController{
		userService: service,
	}
}

func (uc *userController) CreateUser(c echo.Context) error {
	var user models.UserEntity
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("failed to bind user data", nil))
	}

	// Insert user data to database
	userID, err := uc.userService.CreateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "insert failed") {
			return c.JSON(http.StatusBadRequest, utils.FailResponse(err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, utils.FailResponse("failed to insert data. "+err.Error(), nil))
		}
	}

	userOutput, err := uc.userService.GetUser(uint(userID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("user not found", nil))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("user created successfully", userOutput))
}

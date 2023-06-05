package controllers

import (
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

func (uc *userController) LoginUser(c echo.Context) error {
	req := AuthRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("failed to bind data", nil))
	}

	user, token, err := uc.userService.Login(req.Email, req.Password)
	if err != nil {
		if strings.Contains(err.Error(), "login failed") {
			return c.JSON(http.StatusBadRequest, utils.FailResponse(err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, utils.FailResponse("internal server error", nil))
		}
	}

	response := map[string]interface{}{
		"email":   user.Email,
		"user_id": user.ID,
		"token":   token,
		"role":    user.Role,
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("login success", response))
}

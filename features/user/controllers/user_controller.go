package controllers

import (
	"alta-immersive-dashboard/app/middlewares"
	userRepo "alta-immersive-dashboard/features/user/repository"
	userSrv "alta-immersive-dashboard/features/user/service"
	"alta-immersive-dashboard/utils"
	"net/http"
	"strconv"
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

func (uc *userController) UpdateUser(c echo.Context) error {
	var updatedUser userRepo.UserEntity
	err := c.Bind(&updatedUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("failed to bind user data", nil))
	}

	role := middlewares.ExtractRoleFromToken(c)
	idParam := c.Param("user_id")

	var userID int

	if (role == "admin" || role == "user") && idParam == "" {
		userID = middlewares.ExtractUserIDFromToken(c)

	} else if role == "admin" && idParam != "" {
		ID, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid user ID", nil))
		}
		userID = ID
	} else {
		return c.JSON(http.StatusUnauthorized, utils.FailResponse("unauthorized access", nil))
	}

	err = uc.userService.UpdateUser(uint(userID), updatedUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.FailResponse("status internal error", nil))
	}

	// Get user data for response
	user, err := uc.userService.GetUser(uint(userID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("user not found", nil))
	}

	userResponse := EntityToUserResponse(user)

	return c.JSON(http.StatusOK, utils.SuccessResponse("user updated successfully", userResponse))
}

func (uc *userController) DeleteUser(c echo.Context) error {
	role := middlewares.ExtractRoleFromToken(c)
	idParam := c.Param("user_id")

	var userID int

	// Get user data for response
	if (role == "admin" || role == "user") && idParam == "" {
		userID = middlewares.ExtractUserIDFromToken(c)

	} else if role == "admin" && idParam != "" {
		ID, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid user ID", nil))
		}
		userID = ID
	} else {
		return c.JSON(http.StatusUnauthorized, utils.FailResponse("unauthorized access", nil))
	}

	user, err := uc.userService.GetUser(uint(userID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("user not found", nil))
	}

	// Update user status in database
	updatedUser := userRepo.UserEntity{
		Status: "deleted",
	}
	err = uc.userService.UpdateUser(uint(userID), updatedUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.FailResponse("status internal error", nil))
	}

	// Delete user data from database
	err = uc.userService.DeleteUser(uint(userID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse(err.Error(), nil))
	}

	// Response user
	userResponse := EntityToUserResponse(user)
	userResponse.Status = "deleted"

	return c.JSON(http.StatusOK, utils.SuccessResponse("user deleted successfully", userResponse))
}

func (uc *userController) ReadAllUser(c echo.Context) error {
	users, err := uc.userService.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.FailResponse("users not found", nil))
	}

	var userResponses []UserResponse
	for _, userEntity := range users {
		userResponses = append(userResponses, EntityToUserResponse(userEntity))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("users retrieved successfully", userResponses))
}

func (uc *userController) ReadUser(c echo.Context) error {
	idParam := c.Param("user_id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid user ID", nil))
	}

	user, err := uc.userService.GetUser(uint(userID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("user not found", nil))
	}

	userResponse := EntityToUserResponse(user)

	return c.JSON(http.StatusOK, utils.SuccessResponse("user retrieved successfully", userResponse))
}

func (uc *userController) CreateUser(c echo.Context) error {
	var user userRepo.UserEntity
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
	user.ID = userID
	userResponse := EntityToUserResponse(user)

	return c.JSON(http.StatusOK, utils.SuccessResponse("user created successfully", userResponse))
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

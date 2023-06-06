package controllers

import (
	classRepo "alta-immersive-dashboard/features/class/repository"
	classSrv "alta-immersive-dashboard/features/class/service"
	"alta-immersive-dashboard/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type classController struct {
	classService classSrv.ClassService
}

func New(service classSrv.ClassService) *classController {
	return &classController{
		classService: service,
	}
}

func (cc *classController) CreateClass(c echo.Context) error {
	var class classRepo.ClassEntity
	err := c.Bind(&class)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("failed to bind class data", nil))
	}

	// Insert class data to database
	classID, err := cc.classService.CreateClass(class)
	if err != nil {
		if strings.Contains(err.Error(), "insert failed") {
			return c.JSON(http.StatusBadRequest, utils.FailResponse(err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, utils.FailResponse("failed to insert data. "+err.Error(), nil))
		}
	}
	class.ID = classID
	classResponse := EntityToClassResponse(class)

	return c.JSON(http.StatusOK, utils.SuccessResponse("class created successfully", classResponse))
}

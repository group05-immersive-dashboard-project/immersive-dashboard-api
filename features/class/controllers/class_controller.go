package controllers

import (
	classRepo "alta-immersive-dashboard/features/class/repository"
	classSrv "alta-immersive-dashboard/features/class/service"
	"alta-immersive-dashboard/utils"
	"net/http"
	"strconv"
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
	classResponse := EntityToCreateDeleteClassResponse(class)

	return c.JSON(http.StatusOK, utils.SuccessResponse("class created successfully", classResponse))
}

func (cc *classController) ReadClass(c echo.Context) error {
	idParam := c.Param("class_id")
	classID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid class ID", nil))
	}

	class, err := cc.classService.GetClass(uint(classID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("class not found", nil))
	}

	classResponse := EntityToGetUpdateClassResponse(class)

	return c.JSON(http.StatusOK, utils.SuccessResponse("class retrieved successfully", classResponse))
}

func (cc *classController) ReadAllClass(c echo.Context) error {
	classes, err := cc.classService.GetAllClass()
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.FailResponse("classes not found", nil))
	}

	var classResponses []GetUpdateClassResponse
	for _, classEntity := range classes {
		classResponses = append(classResponses, EntityToGetUpdateClassResponse(classEntity))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("classes retrieved successfully", classResponses))
}

func (cc *classController) UpdateClass(c echo.Context) error {
	var updatedClass classRepo.ClassEntity
	err := c.Bind(&updatedClass)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("failed to bind class data", nil))
	}

	idParam := c.Param("class_id")
	classID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid class ID", nil))
	}

	err = cc.classService.UpdateClass(uint(classID), updatedClass)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.FailResponse("status internal error", nil))
	}

	// Get class data for response
	class, err := cc.classService.GetClass(uint(classID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("class not found", nil))
	}

	classResponse := EntityToGetUpdateClassResponse(class)

	return c.JSON(http.StatusOK, utils.SuccessResponse("class updated successfully", classResponse))
}

func (cc *classController) DeleteClass(c echo.Context) error {
	idParam := c.Param("class_id")
	classID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid class ID", nil))
	}

	class, err := cc.classService.GetClass(uint(classID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("class not found", nil))
	}

	// Delete class data from database
	err = cc.classService.DeleteClass(uint(classID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse(err.Error(), nil))
	}

	// Response class
	classResponse := EntityToCreateDeleteClassResponse(class)

	return c.JSON(http.StatusOK, utils.SuccessResponse("class deleted successfully", classResponse))
}

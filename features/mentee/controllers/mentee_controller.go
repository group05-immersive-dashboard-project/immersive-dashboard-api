package controllers

import (
	"alta-immersive-dashboard/features/mentee/repository"
	menteeRepo "alta-immersive-dashboard/features/mentee/repository"
	menteeSrv "alta-immersive-dashboard/features/mentee/service"
	"alta-immersive-dashboard/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type menteeController struct {
	menteeService menteeSrv.MenteeService
}

func New(service menteeSrv.MenteeService) *menteeController {
	return &menteeController{
		menteeService: service,
	}
}

func (mc *menteeController) CreateMentee(c echo.Context) error {
	var mentee menteeRepo.MenteeEntity
	err := c.Bind(&mentee)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("failed to bind mentee data", nil))
	}

	// Insert mentee data to database
	menteeID, err := mc.menteeService.CreateMentee(mentee)
	if err != nil {
		if strings.Contains(err.Error(), "insert failed") {
			return c.JSON(http.StatusBadRequest, utils.FailResponse(err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, utils.FailResponse("failed to insert data. "+err.Error(), nil))
		}
	}
	mentee.ID = menteeID
	menteeResponse := EntityToCreateDeleteMenteeResponse(mentee)

	return c.JSON(http.StatusOK, utils.SuccessResponse("mentee created successfully", menteeResponse))
}

func (mc *menteeController) ReadMentee(c echo.Context) error {
	idParam := c.Param("mentee_id")
	menteeID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid mentee ID", nil))
	}

	mentee, err := mc.menteeService.GetMentee(uint(menteeID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("mentee not found", nil))
	}

	menteeResponse := EntityToReadUpdateMenteeResponse(mentee)

	return c.JSON(http.StatusOK, utils.SuccessResponse("mentee retrieved successfully", menteeResponse))
}

func (mc *menteeController) ReadMenteeFeedbacks(c echo.Context) error {
	idParam := c.Param("mentee_id")
	menteeID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid mentee ID", nil))
	}

	mentee, err := mc.menteeService.GetMentee(uint(menteeID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("mentee not found", nil))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("mentee feedbacks retrieved successfully", mentee.Feedbacks))
}

func (mc *menteeController) ReadAllMentee(c echo.Context) error {
	filters := repository.MenteeFilter{
		ClassID:  c.QueryParam("class_id"),
		StatusID: c.QueryParam("status_id"),
		Category: c.QueryParam("category"),
	}

	var mentees []menteeRepo.MenteeEntity
	var err error

	if filters.IsEmpty() {
		mentees, err = mc.menteeService.GetAllMentee()
	} else {
		mentees, err = mc.menteeService.GetAllMenteeByFilters(filters)
	}

	if err != nil {
		return c.JSON(http.StatusNotFound, utils.FailResponse("mentees not found", nil))
	}

	var menteeResponses []CreateDeleteMenteeResponse
	for _, menteeEntity := range mentees {
		menteeResponses = append(menteeResponses, EntityToCreateDeleteMenteeResponse(menteeEntity))
	}

	if len(menteeResponses) == 0 {
		return c.JSON(http.StatusOK, utils.SuccessResponse("mentees not found", nil))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("mentees retrieved successfully", menteeResponses))
}

func (mc *menteeController) DeleteMentee(c echo.Context) error {
	idParam := c.Param("mentee_id")
	menteeID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid mentee ID", nil))
	}

	mentee, err := mc.menteeService.GetMentee(uint(menteeID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("mentee not found", nil))
	}

	// Delete mentee data from database
	err = mc.menteeService.DeleteMentee(uint(menteeID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse(err.Error(), nil))
	}

	menteeResponse := EntityToCreateDeleteMenteeResponse(mentee)

	return c.JSON(http.StatusOK, utils.SuccessResponse("mentee deleted successfully", menteeResponse))
}

func (mc *menteeController) UpdateMentee(c echo.Context) error {
	var updatedMentee menteeRepo.MenteeEntity
	err := c.Bind(&updatedMentee)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("failed to bind mentee data", nil))
	}

	idParam := c.Param("mentee_id")
	menteeID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid mentee ID", nil))
	}

	err = mc.menteeService.UpdateMentee(uint(menteeID), updatedMentee)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.FailResponse("status internal error", nil))
	}

	// Get mentee data for response
	mentee, err := mc.menteeService.GetMentee(uint(menteeID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("mentee not found", nil))
	}

	menteeResponse := EntityToReadUpdateMenteeResponse(mentee)

	return c.JSON(http.StatusOK, utils.SuccessResponse("mentee updated successfully", menteeResponse))
}

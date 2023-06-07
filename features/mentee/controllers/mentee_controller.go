package controllers

import (
	menteeRepo "alta-immersive-dashboard/features/mentee/repository"
	menteeSrv "alta-immersive-dashboard/features/mentee/service"
	"alta-immersive-dashboard/utils"
	"net/http"
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

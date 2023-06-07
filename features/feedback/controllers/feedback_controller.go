package controllers

import (
	feedbackRepo "alta-immersive-dashboard/features/feedback/repository"
	feedbackSrv "alta-immersive-dashboard/features/feedback/service"
	"alta-immersive-dashboard/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type feedbackController struct {
	feedbackService feedbackSrv.FeedbackService
}

func New(service feedbackSrv.FeedbackService) *feedbackController {
	return &feedbackController{
		feedbackService: service,
	}
}

func (fc *feedbackController) CreateFeedback(c echo.Context) error {
	var feedback feedbackRepo.FeedbackEntity
	err := c.Bind(&feedback)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("failed to bind feedback data", nil))
	}

	// Insert feedback data to database
	feedbackID, err := fc.feedbackService.CreateFeedback(feedback)
	if err != nil {
		if strings.Contains(err.Error(), "insert failed") {
			return c.JSON(http.StatusBadRequest, utils.FailResponse(err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, utils.FailResponse("failed to insert data. "+err.Error(), nil))
		}
	}
	feedback.ID = feedbackID

	return c.JSON(http.StatusOK, utils.SuccessResponse("mentee created successfully", feedback))
}

// func (fc *feedbackController) UpdateFeedback(c echo.Context) error {
// 	var updatedFeedback feedbackRepo.FeedbackEntity
// 	err := c.Bind(&updatedFeedback)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, utils.FailResponse("failed to bind feedback data", nil))
// 	}

// 	idParam := c.Param("feedback_id")
// 	feedbackID, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid feedback ID", nil))
// 	}

// 	err = fc.feedbackService.UpdateFeedback(uint(feedbackID), updatedFeedback)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, utils.FailResponse("status internal error", nil))
// 	}

// 	// Get feedback data for response
// 	feedback, err := fc.feedbackService.GetFeedback(uint(feedbackID))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, utils.FailResponse("feedback not found", nil))
// 	}

// 	menteeResponse := EntityToReadUpdateMenteeResponse(mentee)

// 	return c.JSON(http.StatusOK, utils.SuccessResponse("feedback updated successfully", menteeResponse))
// }

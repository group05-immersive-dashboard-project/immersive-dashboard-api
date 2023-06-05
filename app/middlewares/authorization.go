package middlewares

import (
	"alta-immersive-dashboard/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := ExtractRoleFromToken(c)
		if role != "admin" {
			return c.JSON(http.StatusForbidden, utils.FailResponse("Unauthorized access", nil))
		}

		return next(c)
	}
}

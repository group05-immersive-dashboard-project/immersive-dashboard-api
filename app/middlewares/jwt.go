package middlewares

import (
	"alta-immersive-dashboard/app/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddlewareFunc() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(config.SECRET_JWT),
		SigningMethod: "HS256",
	})
}

func GenerateToken(userId int, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 24 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SECRET_JWT))
}

func ExtractUserIDFromToken(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userID := claims["userID"].(float64)
		return int(userID)
	}
	return 0
}

func ExtractRoleFromToken(e echo.Context) string {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		role := claims["role"].(string)
		return role
	}
	return ""
}

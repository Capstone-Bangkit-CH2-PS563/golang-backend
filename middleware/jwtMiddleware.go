package middleware

import (
	"capstone/constants"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var IsLoggedIn = echojwt.WithConfig(echojwt.Config{
	SigningMethod: "HS256",
	SigningKey:    []byte(constants.SECRET_JWT),
})

func CreateToken(userId int, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	byteSecret := []byte(constants.SECRET_JWT)
	return token.SignedString(byteSecret)
}

// check role user
func IsUser(c echo.Context) (int, error) {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}
	claims := user.Claims.(jwt.MapClaims)
	if claims["role"] != constants.User {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}
	userId := int(claims["userId"].(float64))

	return userId, nil
}

// check role admin
func IsAdmin(c echo.Context) (int, error) {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}
	claims := user.Claims.(jwt.MapClaims)
	if claims["role"] != constants.Admin {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}
	userId := int(claims["userId"].(float64))

	return userId, nil
}

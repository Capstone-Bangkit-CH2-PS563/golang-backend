package controller

import (
	"capstone/middleware"
	"capstone/model/payload"
	"capstone/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegisterUserController(c echo.Context) error {
	payloadUser := payload.CreateUserRequest{}
	c.Bind(&payloadUser)

	if err := c.Validate(&payloadUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload create user",
			"error":   err.Error(),
		})
	}

	response, err := service.CreateUser(&payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error create user",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success register user",
		Status:  http.StatusText(200),
		Data:    response,
	})
}

func LoginUserController(c echo.Context) error {
	payloadUser := payload.LoginUserRequest{}

	c.Bind(&payloadUser)

	if err := c.Validate(&payloadUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload login user",
			"error":   err.Error(),
		})
	}

	response, err := service.LoginUser(&payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Login",
		Status:  http.StatusText(200),
		Data:    response,
	})
}

func GetUsersController(c echo.Context) error {
	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "this route only for admin",
		})
	}

	users, err := service.GetUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "Succes get all users",
		Data:    users,
	})
}

func UpdateUserController(c echo.Context) error {
	payloadUser := payload.UpdateUserRequest{}
	c.Bind(&payloadUser)
	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "this route only for admin",
		})
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	user, err := service.GetUserById(int(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}


	if err := c.Validate(payloadUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload update user",
			"error":   err.Error(),
		})
	}

	response, err := service.UpdateUser(user, &payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success update User",
		Data:    response,
	})
}

package controller

import (
	"capstone/middleware"
	"capstone/model/payload"
	"capstone/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddBookmarkPlantController(c echo.Context) error {
	payloadBookmark := payload.CreateBookmarkRequest{}

	userId, err := middleware.IsUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "this route only for user",
		})
	}

	c.Bind(&payloadBookmark)

	if err := c.Validate(&payloadBookmark); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload favorite warehouse",
			"error":   err.Error(),
		})
	}

	response, err := service.CreateBookmarkPlant(userId, &payloadBookmark)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": response,
	})
}

func GetBookmarkPlantListByUserIdController(c echo.Context) error {
	userId, err := middleware.IsUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "this route only for user",
		})
	}

	favorite, err := service.GetBookmarkListByUserId(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "Success get bookmark list by user id",
		Data:    favorite,
	})

}

func DeleteBookmarkController(c echo.Context) error {
	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "this route only for admin",
		})
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	bookmark, err := service.GetBookmarkById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = service.DeleteBookmark(bookmark)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "delete bookmark complete")
}

func GetAllBookmarkPlantController(c echo.Context) error {
	response, err := service.GetAllBookmarks()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "Succes get all bookmarks plant",
		Data:    response,
	})
}
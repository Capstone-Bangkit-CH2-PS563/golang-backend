package controller

import (
	"capstone/middleware"
	"capstone/model/payload"
	"capstone/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreatePlantController(c echo.Context) error {
	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "this route only for admin",
		})
	}

	payloadPlant := payload.CreatePlantRequest{}
	c.Bind(&payloadPlant)

	if err := c.Validate(payloadPlant); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload create Plant",
			"error":   err.Error(),
		})
	}

	response, err := service.CreatePlant(&payloadPlant)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error create Plant",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success create plant",
		Status:  http.StatusText(200),
		Data:    response,
	})
}

func GetAllPlantController(c echo.Context) error {
	response, err := service.GetAllPlants()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "Succes get all plant",
		Data:    response,
	})
}

func GetPlantByIdController(c echo.Context) error {
	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "this route only for admin",
		})
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	response, err := service.GetPlantById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "Succes get plant by id",
		Data:    response,
	})
}

func UpdatePlantController(c echo.Context) error {
	payloadPlant := payload.UpdatePlantRequest{}
	c.Bind(&payloadPlant)
	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "this route only for admin",
		})
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	plant, err := service.GetPlantById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(payloadPlant); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload update plant",
			"error":   err.Error(),
		})
	}

	response, err := service.UpdatePlant(plant, &payloadPlant)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success update plant",
		Data:    response,
	})
}

func DeletePlantController(c echo.Context) error {
	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "this route only for admin",
		})
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	plant, err := service.GetPlantById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = service.DeletePlant(plant)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "delete plant complete")
}

func SearchPlantController(c echo.Context) error {

	keyword := c.QueryParam("keyword")
	response, err := service.SearchPlant(keyword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "Succes get search data",
		Data:    response,
	})
}
package controller

import (
	"capstone/model/payload"
	"capstone/util"
	"net/http"

	"github.com/labstack/echo/v4"
)


func UploadFileController(c echo.Context) error {

	file, err := c.FormFile("image_url")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error To upload Image")
	}
	src, err := file.Open()
	if err != nil {
		return  err
	}
	
	defer src.Close()

	resp, err := util.UploadFile(src, file.Filename)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error To upload Image")
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "image upload success",
		Data:    resp,
	})
}

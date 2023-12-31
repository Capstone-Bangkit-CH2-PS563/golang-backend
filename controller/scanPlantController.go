package controller

import (
	"bytes"
	"capstone/middleware"
	"capstone/model/payload"
	"capstone/service"
	"capstone/util"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
)

func ScanPlantController(c echo.Context) error {
	payloadScanPlant := payload.CreateScanPlantRequest{}
	userId, err := middleware.IsUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "this route only for user",
		})
	}
	payloadScanPlant.UserID = uint(userId)

	file, err := c.FormFile("image_for_scan")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error To upload Image")
	}
	src, err := file.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	image_url, err := util.UploadFile(src, file.Filename)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error To upload Image")
	}

	payloadScanPlant.ScanPlantImage = image_url

	payloadML := map[string]interface{}{
		"image_url": image_url,
	}

	jsonPayload, err := json.Marshal(payloadML)
	if err != nil {
		return err
	}

	modelResp, err := http.Post("url_model", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	defer modelResp.Body.Close()

	body, err := io.ReadAll(modelResp.Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error To read response body from ML")
	}

	var respML payload.ResponseBodyFromML

	err = json.Unmarshal(body, &respML)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error To unmarshaling JSON response")
	}

	plant, err := service.GetPlantBySlug(slug.Make(respML.PlantName))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	payloadScanPlant.PlantID = plant.ID

	resp, err := service.CreateScanPlant(&payloadScanPlant)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": resp,
	})

}

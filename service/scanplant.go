package service

import (
	"capstone/model"
	"capstone/model/payload"
	"capstone/repository/database"
)

func CreateScanPlant(req *payload.CreateScanPlantRequest) (resp payload.ManageScanPlantResponse, err error) {

	newScanPlant := &model.ScanPlant{
		ScanPlantImage: req.ScanPlantImage,
		PlantID:        req.PlantID,
		UserID:         req.UserID,
	}

	err = database.CreateScanPlant(newScanPlant)
	if err != nil {
		return
	}

	plant, err := database.GetPlantById(uint64(newScanPlant.ID))
	if err != nil {
		return
	}

	resp = payload.ManageScanPlantResponse{
		ScanPlantImage: newScanPlant.ScanPlantImage,
		PlantID:        newScanPlant.PlantID,
		UserID:         newScanPlant.UserID,
		Plant:          plant,
	}

	return
}

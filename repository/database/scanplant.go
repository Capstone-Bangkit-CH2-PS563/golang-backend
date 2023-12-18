package database

import (
	"capstone/config"
	"capstone/model"
)

func CreateScanPlant(ScanPlant *model.ScanPlant) error {
	if err := config.DB.Preload("Plant").Preload("User").Create(&ScanPlant).Error; err != nil {
		return err
	}

	return nil
}

func GetAllScanPlants() (ScanPlants []model.ScanPlant, err error) {
	if err = config.DB.Preload("Plant").Preload("User").Find(&ScanPlants).Error; err != nil {
		return
	}
	return
}

func DeleteScanPlants(ScanPlant *model.ScanPlant) error {
	if err := config.DB.Delete(&ScanPlant).Error; err != nil {
		return err
	}

	return nil
}

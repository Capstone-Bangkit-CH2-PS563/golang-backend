package model

import "gorm.io/gorm"

type ScanPlant struct {
	gorm.Model
	ScanPlantImage string `json:"scan_plant_img" form:"scan_plant_img"`
	PlantID        uint   `json:"plant_id" form:"plant_id"`
	Plant          Plant
	UserID         uint `json:"user_id" form:"user_id"`
	User           User
}

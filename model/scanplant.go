package model

import "gorm.io/gorm"

type ScanPlant struct {
	gorm.Model
	ScanPlantImage string `json:"scan_plant_img" form:"scan_plant_img"`
	UserID         uint   `json:"user_id" form:"user_id"`
	User           User
}

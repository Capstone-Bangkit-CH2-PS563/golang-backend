package model

import "gorm.io/gorm"

type Plant struct {
	gorm.Model
	PlantName string `json:"plant_name" form:"plant_name"`
	PlantSlug string `json:"plant_slug" form:"plant_slug"`
	PlantImage string `json:"plant_image" form:"plant_image"`
	PlantStatusId string `json:"plant_status_id" form:"plant_status_id"`
	PlantNutritionId string `json:"plant_nutrition_id" form:"plant_nutrition_id"`
	PlantStatusEn string `json:"plant_status_en" form:"plant_status_en"`
	PlantNutritionEn string `json:"plant_nutrition_en" form:"plant_nutrition_en"`
}

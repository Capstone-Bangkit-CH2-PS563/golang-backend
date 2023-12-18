package payload

type CreatePlantRequest struct {
	PlantName        string `json:"plant_name" form:"plant_name" validate:"required"`
	PlantImage       string `json:"plant_image" form:"plant_image" validate:"required"`
	PlantStatusId    string `json:"plant_status_id" form:"plant_status_id" validate:"required"`
	PlantNutritionId string `json:"plant_nutrition_id" form:"plant_nutrition_id" validate:"required"`
	PlantStatusEn    string `json:"plant_status_en" form:"plant_status_en" validate:"required"`
	PlantNutritionEn string `json:"plant_nutrition_en" form:"plant_nutrition_en" validate:"required"`
}
type UpdatePlantRequest struct {
	PlantName        string `json:"plant_name" form:"plant_name"`
	PlantImage       string `json:"plant_image" form:"plant_image"`
	PlantStatusId    string `json:"plant_status_id" form:"plant_status_id"`
	PlantNutritionId string `json:"plant_nutrition_id" form:"plant_nutrition_id"`
	PlantStatusEn    string `json:"plant_status_en" form:"plant_status_en"`
	PlantNutritionEn string `json:"plant_nutrition_en" form:"plant_nutrition_en"`
}

type ManagePlantResponse struct {
	PlantName        string `json:"plant_name"`
	PlantImage       string `json:"plant_image"`
	PlantStatusId    string `json:"plant_status_id"`
	PlantNutritionId string `json:"plant_nutrition_id"`
	PlantStatusEn    string `json:"plant_status_en"`
	PlantNutritionEn string `json:"plant_nutrition_en"`
}

type GetAllPlantResponse struct {
	PlantID          uint `json:"plant_id"`
	PlantName        string `json:"plant_name"`
	PlantImage       string `json:"plant_image"`
	PlantStatusId    string `json:"plant_status_id"`
	PlantNutritionId string `json:"plant_nutrition_id"`
	PlantStatusEn    string `json:"plant_status_en"`
	PlantNutritionEn string `json:"plant_nutrition_en"`
}

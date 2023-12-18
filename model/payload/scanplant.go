package payload

import "capstone/model"

type CreateScanPlantRequest struct {
	ScanPlantImage string `json:"scan_plant_img" form:"scan_plant_img"`
	PlantID        uint   `json:"plant_id" form:"plant_id"`
	UserID         uint   `json:"user_id" form:"user_id"`
}

type ManageScanPlantResponse struct {
	UserID         uint `json:"user_id"`
	ScanPlantImage string `json:"scan_plant_img"`
	PlantID        uint   `json:"plant_id"`
	Plant          *model.Plant
	
}

type ResponseBodyFromML struct {
	PlantID   int    `json:"plant_id"`
	PlantName string `json:"plant_name"`
}

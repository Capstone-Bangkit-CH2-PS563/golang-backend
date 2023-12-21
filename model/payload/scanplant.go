package payload

import "capstone/model"

type CreateScanPlantRequest struct {
	ScanPlantImage string `json:"scan_plant_img" form:"scan_plant_img"`
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


type PredictionResponse struct {
	Data struct {
		Benefit             []string  `json:"benefit"`
		Calories            string    `json:"calories"`
		Confidence          float64   `json:"confidence"`
		VegetablePrediction string    `json:"vegetable_prediction"`
	} `json:"data"`
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
}

package payload

type CreateBookmarkRequest struct {
	UserID  uint `json:"user_id" form:"user_id"`
	PlantID uint `json:"plant_id" form:"plant_id" validate:"required"`
}
type CreateBookmarkResponse struct {
	PlantID uint`json:"plant_id" form:"plant_id"`
	Plant  GetAllPlantResponse 
}
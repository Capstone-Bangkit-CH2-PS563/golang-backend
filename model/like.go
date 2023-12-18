package model

type Like struct {
	ID      uint `json:"id" form:"id" gorm:"primarykey"`
	UserID  uint `json:"user_id" form:"user_id"`
	PlantID uint `json:"plant_id" form:"plant_id"`
	Plant   Plant
	User    User
}
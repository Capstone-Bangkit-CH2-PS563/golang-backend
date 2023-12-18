package model

type Bookmark struct {
	ID      uint64 `json:"id" form:"id" gorm:"primarykey"`
	UserID  uint `json:"user_id" form:"user_id"`
	PlantID uint `json:"plant_id" form:"plant_id"`
	Plant   Plant
	User    User `json:"-"`
}

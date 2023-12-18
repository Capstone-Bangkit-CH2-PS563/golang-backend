package database

import (
	"capstone/config"
	"capstone/model"
)

func CreatePlant(plant *model.Plant) error {
	if err := config.DB.Create(plant).Error; err != nil {
		return err
	}
	return nil
}

func GetAllPlants() (plants []model.Plant, err error) {
	if err = config.DB.Find(&plants).Error; err != nil {
		return
	}
	return
}
func GetPlantById(id uint64) (plant *model.Plant, err error) {
	if err = config.DB.Where("id = ?", id).First(&plant).Error; err != nil {
		return
	}
	return
}

func UpdatePlant(plant *model.Plant) error {
	if err := config.DB.Updates(&plant).Error; err != nil {
		return err
	}
	return nil
}

func DeletePlant(plant *model.Plant) error {
	if err := config.DB.Delete(plant).Error; err != nil {
		return err
	}

	return nil
}

func SearchPlant(keyword string) (plants []model.Plant, err error) {
	if err = config.DB.Where("plant_name LIKE ?", "%"+keyword+"%").Find(&plants).Error; err != nil {
		return
	}

	return
}

func GetPlantBySlug(slug string) (plant *model.Plant, err error) {
	if err = config.DB.Where("plant_slug = ?", slug).First(&plant).Error; err != nil {
		return
	}
	return
}

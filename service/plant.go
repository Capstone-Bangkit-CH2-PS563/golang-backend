package service

import (
	"capstone/model"
	"capstone/model/payload"
	"capstone/repository/database"
	"errors"

	"github.com/gosimple/slug"
)

func CreatePlant(req *payload.CreatePlantRequest) (resp payload.ManagePlantResponse, err error) {

	newPlant := &model.Plant{
		PlantName:        req.PlantName,
		PlantSlug:        slug.Make(req.PlantName),
		PlantImage:       req.PlantImage,
		PlantStatusId:    req.PlantStatusId,
		PlantNutritionId: req.PlantNutritionId,
		PlantStatusEn:    req.PlantStatusEn,
		PlantNutritionEn: req.PlantNutritionEn,
	}

	err = database.CreatePlant(newPlant)
	if err != nil {
		return
	}

	resp = payload.ManagePlantResponse{
		PlantName:        newPlant.PlantName,
		PlantImage:       newPlant.PlantImage,
		PlantStatusId:    newPlant.PlantStatusId,
		PlantNutritionId: newPlant.PlantNutritionId,
		PlantStatusEn:    newPlant.PlantStatusEn,
		PlantNutritionEn: newPlant.PlantNutritionEn,
	}

	return
}

func GetAllPlants() ([]model.Plant, error) {
	plants, err := database.GetAllPlants()
	if err != nil {
		return nil, err
	}

	return plants, nil
}

func GetPlantById(id uint64) (*model.Plant, error) {
	plant, err := database.GetPlantById(id)
	if err != nil {
		return nil, err
	}

	return plant, nil
}

func UpdatePlant(plant *model.Plant, req *payload.UpdatePlantRequest) (resp payload.ManagePlantResponse, err error) {

	plant.PlantName = req.PlantName
	plant.PlantSlug = slug.Make(req.PlantName)
	plant.PlantImage = req.PlantImage
	plant.PlantStatusId = req.PlantStatusId
	plant.PlantNutritionId = req.PlantNutritionId
	plant.PlantStatusEn = req.PlantStatusEn
	plant.PlantNutritionEn = req.PlantNutritionEn

	err = database.UpdatePlant(plant)
	if err != nil {
		return resp, errors.New("Can't update plant")
	}

	updatedPlant, _ := database.GetPlantById(uint64(plant.ID))

	resp = payload.ManagePlantResponse{
		PlantName:        updatedPlant.PlantName,
		PlantImage:       updatedPlant.PlantImage,
		PlantStatusId:    updatedPlant.PlantStatusId,
		PlantNutritionId: updatedPlant.PlantNutritionId,
		PlantStatusEn:    updatedPlant.PlantStatusEn,
		PlantNutritionEn: updatedPlant.PlantNutritionEn,
	}

	return resp, nil
}

func DeletePlant(plant *model.Plant) error {
	err := database.DeletePlant(plant)
	if err != nil {
		return err
	}
	return nil
}

func SearchPlant(keyword string)([]model.Plant, error){
	plants, err:= database.SearchPlant(keyword)
	if err != nil {
		return nil, err
	}
	return plants, nil
}

func GetPlantBySlug(slug string) (*model.Plant, error) {
	plant, err := database.GetPlantBySlug(slug)
	if err != nil {
		return nil, err
	}

	return plant, nil
}
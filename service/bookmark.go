package service

import (
	"capstone/model"
	"capstone/model/payload"
	"capstone/repository/database"
	"errors"
)

func CreateBookmarkPlant(id int, req *payload.CreateBookmarkRequest) (resp any, err error) {
	user, err := database.GetUserByID(id)
	if err != nil {
		return resp, errors.New("User not found")
	}

	plant, err := database.GetPlantById(uint64(req.PlantID))
	if err != nil {
		return resp, errors.New("Plant data not found")
	}

	newBookmark := &model.Bookmark{
		UserID:  user.ID,
		PlantID: req.PlantID,
	}

	bookmark, err := database.CheckBookmark(newBookmark)
	if err != nil {
		err = database.CreateBookmark(newBookmark)
		if err != nil {
			return resp, errors.New("Can't Create Bookmark")
		}

		resp = payload.CreateBookmarkResponse{
			PlantID: newBookmark.PlantID,
			Plant: payload.GetAllPlantResponse{
				PlantID:          plant.ID,
				PlantName:        plant.PlantName,
				PlantImage:       plant.PlantImage,
				PlantStatusId:    plant.PlantStatusId,
				PlantNutritionId: plant.PlantNutritionId,
				PlantStatusEn:    plant.PlantStatusEn,
				PlantNutritionEn: plant.PlantNutritionEn,
			},
		}
	} else {
		err = database.DeleteBookmark(bookmark)
		if err != nil {
			return resp, errors.New("Can't Delete bookmark")
		}
		resp = "Success Delete Favorite"

		return
	}

	return
}


func GetBookmarkListByUserId(userId int) (bookmark []model.Bookmark, err error) {

	bookmark, err = database.GetBookmarkUserByID(userId)
	if err != nil {
		return nil, errors.New("Failed to get bookmark")
	}

	return bookmark, nil
}


func GetBookmarkById(id uint64) (*model.Bookmark, error) {
	bookmark, err := database.GetBookmarkById(id)
	if err != nil {
		return nil, err
	}

	return bookmark, nil
}

func DeleteBookmark(bookmark *model.Bookmark) error {
	err := database.DeleteBookmark(bookmark)
	if err != nil {
		return err
	}
	return nil
}

func GetAllBookmarks() ([]model.Bookmark, error) {
	bookmarks, err := database.GetAllBookmarks()
	if err != nil {
		return nil, err
	}

	return bookmarks, nil
}

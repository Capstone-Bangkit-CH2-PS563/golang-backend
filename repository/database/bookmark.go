package database

import (
	"capstone/config"
	"capstone/model"
)

func CreateBookmark(bookmark *model.Bookmark) error {
	if err := config.DB.Preload("Plant").Create(&bookmark).Error; err != nil {
		return err
	}

	return nil
}

func CheckBookmark(bookmark *model.Bookmark) (bookmarks *model.Bookmark, err error) {
	if err := config.DB.Where("user_id = ? AND plant_id = ?", bookmark.UserID, bookmark.PlantID).First(&bookmark).Error; err != nil {
		return nil, err
	}

	bookmarks = bookmark

	return bookmarks, nil
}

func GetAllBookmarks() (bookmarks []model.Bookmark, err error) {
	if err = config.DB.Preload("Plant").Find(&bookmarks).Error; err != nil {
		return
	}
	return
}

func DeleteBookmark(bookmark *model.Bookmark) error {
	if err := config.DB.Delete(&bookmark).Error; err != nil {
		return err
	}

	return nil
}

func GetBookmarkUserByID(userId int) (bookmark []model.Bookmark, err error) {
	if err = config.DB.Preload("Plant").Where("user_id = ?", userId).Find(&bookmark).Error; err != nil {
		return
	}

	return bookmark, nil
}

func GetBookmarkById(id uint64) (bookmark *model.Bookmark, err error) {
	if err = config.DB.Where("id = ?", id).First(&bookmark).Error; err != nil {
		return
	}
	return
}


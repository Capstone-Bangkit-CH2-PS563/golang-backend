package database

import (
	"capstone/config"
	"capstone/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(user *model.User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func IsEmailAvailable(email string) bool {
	var count int64
	user := model.User{}
	if err := config.DB.Model(&user).Where("email = ?", email).Count(&count).Error; err != nil {
		echo.NewHTTPError(http.StatusNotFound, err.Error())
		return false
	}

	return count == 0
}

func GetUserByEmail(email string) (user model.User, err error) {
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GetUsers() (users []model.User, err error) {
	if err = config.DB.Model(&model.User{}).Find(&users).Error; err != nil {
		return
	}

	return
}

func CountUsers() (count int64, err error) {
	if err := config.DB.Model(&model.User{}).Where("role = ?", "USER").Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func UpdateUser(user *model.User) error {
	if err := config.DB.Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(user *model.User) error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByID(id int) (user *model.User, err error) {
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}


package service

import (
	"capstone/middleware"
	"capstone/model"
	"capstone/model/payload"
	"capstone/repository/database"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(req *payload.CreateUserRequest) (resp payload.CreateUserResponse, err error) {
	if req.ConfirmPassword != req.Password {
		return resp, errors.New("Password not match")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	if !database.IsEmailAvailable(req.Email) {
		return resp, errors.New("email is already registered")
	}

	newUser := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(passwordHash),
	}

	err = database.CreateUser(newUser)
	if err != nil {
		return
	}

	userSuccess, err := database.GetUserByEmail(newUser.Email)
	if err != nil {
		return
	}

	resp = payload.CreateUserResponse{
		UserID: userSuccess.ID,
		Email:  userSuccess.Email,
	}

	return
}

func LoginUser(req *payload.LoginUserRequest) (res payload.LoginUserResponse, err error) {

	user, err := database.GetUserByEmail(req.Email)
	if err != nil {
		return res, errors.New("Email Not Registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return res, errors.New("Wrong Password")
	}

	token, err := middleware.CreateToken(int(user.ID), user.Role)
	if err != nil {
		return res, errors.New("Failed To Create Token")
	}

	user.Token = token

	res = payload.LoginUserResponse{
		UserID: user.ID,
		Email:  user.Email,
		Token:  user.Token,
	}

	return
}

func GetUsers() (resp []payload.GetAllUserResponse, err error) {
	users, err := database.GetUsers()
	if err != nil {
		return nil, errors.New("Error getting users")
	}

	resp = []payload.GetAllUserResponse{}
	for _, user := range users {
		resp = append(resp, payload.GetAllUserResponse{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
			Role:  user.Role,
		})
	}

	return
}

func GetUserById(id int) (*model.User, error) {
	plant, err := database.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return plant, nil
}

func UpdateUser(User *model.User, req *payload.UpdateUserRequest) (resp payload.UpdateUserResponse, err error) {

	User.Name = req.Name
	User.Role = req.Role
	

	err = database.UpdateUser(User)
	if err != nil {
		return resp, errors.New("Can't update User")
	}

	updatedUser, _ := database.GetUserByID(int(User.ID))

	resp = payload.UpdateUserResponse{
		UserID: updatedUser.ID,
		Email:   updatedUser.Email,
		Role: updatedUser.Role,
	}

	return resp, nil
}
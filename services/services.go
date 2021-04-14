package services

import (
	"test2/database"
	"test2/domain"
)

var (
	UserService userServiceInterface = &userServiceStruct{}
)

type userServiceStruct struct {
}

type userServiceInterface interface {
	CreateUser(*domain.User) (int, error)
	Login(*domain.UserLogin) error
}

func (service *userServiceStruct) CreateUser(details *domain.User) (int, error) {
	if err := details.CheckDetails(); err != nil {
		return -1, err
	}
	userID := database.UserDB.SaveUserDetails(*details)
	return userID, nil
}

func (service *userServiceStruct) Login(details *domain.UserLogin) error {
	if err := details.CheckCredentials(); err != nil {
		return err
	}
	err := database.UserDB.VerifyCredentials(*details)
	if err != nil {
		return err
	}
	return nil
}

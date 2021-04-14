package domain

import (
	"errors"
)

type User struct {
	UserEmail string `json:"email"`
	Password  string `json:"password"`
}

type UserLogin struct {
	UserID       int    `json:"userID"`
	UserEmail    string `json:"email"`
	UserPassword string `json:"password"`
}

func (info *User) CheckDetails() error {
	if info.UserEmail == "" || len(info.UserEmail) < 5 {
		return errors.New("Please Enter The Valid Email Address")
	} else if len(info.Password) < 5 {
		return errors.New("Please enter the password again length of password should be atleast 5")
	}
	return nil
}

func (checkDetails *UserLogin) CheckCredentials() error {
	if checkDetails.UserID < 0 {
		errors.New("userID is not valid")
	}
	return nil
}

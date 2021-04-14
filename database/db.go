package database

import (
	"errors"
	"test2/domain"
)

var (
	UserDB userDBInterface = &userDBstruct{}
	db     map[int]domain.User
	userID int = 1
)

type userDBstruct struct {
}

type userDBInterface interface {
	SaveUserDetails(domain.User) int
	VerifyCredentials(domain.UserLogin) error
}

func init() {
	db = make(map[int]domain.User)
}

//SaveUserDetails : To  Save the user Details
func (user *userDBstruct) SaveUserDetails(details domain.User) int {
	db[userID] = details
	savedUserID := userID
	userID++ //to increase the primary key of the user ID by 1 after creating the first user
	return savedUserID
}

func (user *userDBstruct) VerifyCredentials(details domain.UserLogin) error {
	_, found := db[details.UserID]
	if !found {
		return errors.New("user not found")
	}
	savedDetails := db[details.UserID]
	if savedDetails.UserEmail != details.UserEmail {
		return errors.New("email address is not valid")
	} else if savedDetails.Password != details.UserPassword {
		return errors.New("Invalid credentails")
	}
	return nil
}

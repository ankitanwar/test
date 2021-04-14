package controllers

import (
	"net/http"
	"test2/domain"
	"test2/services"

	"github.com/gin-gonic/gin"
)

var (
	UserLogin userLoginInterface = &userLoginStruct{}
)

type userLoginStruct struct {
}

type userLoginInterface interface {
	CreateUser(c *gin.Context)
	Login(c *gin.Context)
}

func (contorller *userLoginStruct) CreateUser(c *gin.Context) {
	details := &domain.User{}
	if err := c.ShouldBindJSON(details); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to Decode the  emailAddress and Password"})
		return
	}
	userID, err := services.UserService.CreateUser(details)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "user has been created successfully", "userID": userID})

}

func (controller *userLoginStruct) Login(c *gin.Context) {
	loginDetails := &domain.UserLogin{}
	if err := c.ShouldBindJSON(loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to get the user details"})
		return
	}
	err := services.UserService.Login(loginDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "User has been logged in successfully"})

}

package application

import (
	"test2/controllers"
)

func urlMapping() {
	router.POST("/newUser", controllers.UserLogin.CreateUser)
	router.POST("/login", controllers.UserLogin.Login)
}

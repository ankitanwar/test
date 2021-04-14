package application

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {
	urlMapping()
	router.Run(":8080")
}

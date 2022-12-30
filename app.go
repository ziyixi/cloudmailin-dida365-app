package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/", HandleCMIPost)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}

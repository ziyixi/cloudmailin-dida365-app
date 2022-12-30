package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func testEnv() {
	_, err := LoginDidaClient()
	if err != nil {
		panic("Couldn't login in dida365. Check .env file")
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// basic auth
	err := godotenv.Load()
	if err != nil {
		panic("Couldn't load .env")
	}
	cmiUser := os.Getenv("cloudmailin_username")
	cmiPass := os.Getenv("cloudmailin_password")
	if len(cmiUser) == 0 || len(cmiPass) == 0 {
		panic("cloudmailin_username or cloudmailin_password is not in .env")
	}
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		cmiUser: cmiPass,
	}))

	authorized.POST("/", HandleCMIPost)

	return r
}

func main() {
	testEnv()
	gin.SetMode(gin.ReleaseMode)

	r := setupRouter()
	r.Run(":8080")
}

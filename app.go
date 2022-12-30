package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func testEnv() {
	_, err := LoginDidaClient()
	if err != nil {
		panic(err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// basic auth
	godotenv.Load() // it's OK if no .env, as we read from ENV variables instead
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
	r.Run(":" + os.Getenv("PORT"))
}

package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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

	authorized.POST("/api/CloudmailinDida365App", HandleCMIPost)

	return r
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	for _, pair := range os.Environ() {
		fmt.Println(pair)
	}

	listenAddr := ":" + os.Getenv("PORT")
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	r := setupRouter()
	r.Run(listenAddr)
}

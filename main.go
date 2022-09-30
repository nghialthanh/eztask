package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"eztask/routes"

	"eztask/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	str := time.Now().Local().Add(time.Hour * time.Duration(24))
	fmt.Println(time.Now().Local())
	fmt.Println(str)

	database.Init()

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + port)

	log.Println("API is running!")
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vishwaszadte/event-magnet-be-core/api"
	"github.com/vishwaszadte/event-magnet-be-core/configs"
	"github.com/vishwaszadte/event-magnet-be-core/utils"
)

func init() {
	var err error

	if os.Getenv("ENV") == "DEV" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Failed to load")
		}
	}

	configs.DB, err = utils.ConnectToDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello there",
		})
	})

	api.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Starting the server on port:", port)
	r.Run()
}

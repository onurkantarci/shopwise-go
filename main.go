package main

import (
	"time"

	"github.com/gin-contrib/cors" // Import the CORS package
	"github.com/gin-gonic/gin"
	"github.com/onurkantarci/shopwise/config"
	"github.com/onurkantarci/shopwise/routes"
)

func main() {
	router := gin.New()

	config.Connect()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.DealRoute(router)
	routes.StoreRoute(router)

	router.Run(":4545")
}

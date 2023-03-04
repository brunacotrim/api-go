package main

import (
	"log"

	"github.com/brunacotrim/api-go/src/configuration/logger"
	"github.com/brunacotrim/api-go/src/controller/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("About to start user application")

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"

	"github.com/Daniel-Fonseca-da-Silva/Golang-Secret-API/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	route := gin.Default()

	routes.InitRoutes(&route.RouterGroup)
	if err := route.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

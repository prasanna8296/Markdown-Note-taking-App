package main

import (
	"markdown-notes/config"
	"markdown-notes/models"
	"markdown-notes/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.ConnectDb()
	config.DB.AutoMigrate(&models.Note{})
	r := gin.Default()
	routes.SetUpRoutes(r)
	r.Run(":8080")

}

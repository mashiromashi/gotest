package main

import (
	"mashiromashi/gotest/controllers"
	"mashiromashi/gotest/models"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "goTest"
)

func setupRouter() *gin.Engine {

	db := models.SetupModels()
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/books", controllers.FindBooks)

	// Get user value

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

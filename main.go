package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {

	router := gin.Default()

	// Serve frontend static files
	router.Use(static.ServeRoot("/", "./public"))

	// Setup route group for the API
	api := router.Group("/")
	{
		api.GET("/d", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H {
			"message": "pong",
			})
		})
	}

	return router
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
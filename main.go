package main

import (
	"api_assignment/auth"
	"api_assignment/config"
	"api_assignment/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	config.Connect()

	// Create a new Gin router
	r := gin.Default()

	// Routes
	r.POST("/signup", handlers.SignUp)
	r.POST("/signin", handlers.SignIn)
	r.POST("/refresh", handlers.RefreshToken)

	r.GET("/protected", auth.Authenticate(), handlers.ProtectedRoute)

	// Start the server
	r.Run(":8080")
}

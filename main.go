package main

import (
	"github.com/AlmedinNasufi/go-jwt/controllers"
	"github.com/AlmedinNasufi/go-jwt/initializers"
	"github.com/AlmedinNasufi/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}

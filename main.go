package main

import (
	"webimg/config"
	"webimg/controllers"
	"webimg/initializers"
	"webimg/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadConfig()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
	initializers.ConnectToMinio()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.POST("/upload", middleware.RequireAuth, controllers.Upload)
	r.GET("/image/:filename", middleware.RequireAuth, controllers.GetImg)

	r.Run(":" + config.GetConfig().Port)
}

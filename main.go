package main

import (
	"exchange-api/controllers"
	"exchange-api/initializers"
	"exchange-api/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {	
	r := gin.Default()

	r.POST("/signup", controllers.UserCreate)
	r.POST("/login", controllers.UserLogin)

	r.GET("/user", middleware.RequireAuth, controllers.UserIndex)
	r.GET("/user/:id", middleware.RequireAuth, controllers.UserShow)
	r.PUT("/user/:id", middleware.RequireAuth, controllers.UserUpdate)
	r.DELETE("/user/:id", middleware.RequireAuth, controllers.UserDelete)

	r.GET("/log", middleware.RequireAuth, controllers.LogIndex)
	r.GET("/log/:id", middleware.RequireAuth, controllers.LogShow)

	r.GET("/exchange", middleware.RequireAuth, controllers.GetExchange)

	r.Run()
}
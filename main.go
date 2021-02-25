package main

import (
	"go-rest-api-starter/config"
	"go-rest-api-starter/controller"
	"go-rest-api-starter/handler"
	"go-rest-api-starter/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.DBInit()
	controller := controller.NewController(db)
	hdl := handler.NewHandler(*controller)

	routes := gin.Default()
	routes.Use(service.CORSMiddleware())
	v1 := routes.Group("api/v1/")
	v1.POST("create", hdl.CreateUser)
	v1.GET("user", hdl.GetUserById)
	v1.GET("users", hdl.GetAllUsers)
	v1.POST("update", hdl.UpdateUser)

	routes.Run(":8800")
}

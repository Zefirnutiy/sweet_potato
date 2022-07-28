package routes

import (
	"github.com/Zefirnutiy/sweet_potato.git/controllers"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)

func Routs(port string) {
	router := gin.Default()

	auth := router.Group("/api/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	test := router.Group("/api/test")
	{
		test.GET("/", utils.TokenCheckedFromHeader, controllers.Test)
	}
	client := router.Group("/api/client")
	{	
		client.GET("/getClients/:tokken", controllers.GetClients)
		client.GET("/getClientById/:tokken/:id", controllers.GetClientById)
		client.POST("/create", controllers.CreateClient)
		client.PATCH("/update", controllers.UpdateClient)
		client.DELETE("/delete/:tokken/:id", controllers.DeleteClient)
	}
	router.Run(port)

}

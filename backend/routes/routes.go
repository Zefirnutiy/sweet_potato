package routes

import (
	"github.com/Zefirnutiy/sweet_potato.git/controllers"
	// "github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)

func Routs(port string) {
	router := gin.Default()

	organization := router.Group("/api/auth")
	{
		organization.POST("/register", controllers.Register)
		organization.POST("/login", controllers.Login)
	}

	// client := router.Group("/api/client")
	// {
	// 	client.GET("/getClients/:token", controllers.GetClients)
	// 	client.GET("/getClientById/:tokken/:id", controllers.GetClientById)
	// 	client.POST("/create", controllers.CreateClient)
	// 	client.PATCH("/update", controllers.UpdateClient)
	// 	client.DELETE("/delete/:tokken/:id", controllers.DeleteClient)
	// }
	router.Run(port)

}

package routes

import (
	"github.com/Zefirnutiy/sweet_potato.git/controllers"
	"github.com/Zefirnutiy/sweet_potato.git/utils"

	// "github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)

func Routs(port string) {
	router := gin.Default()

	organization := router.Group("/api/organization")
	{
		organization.POST("/register", controllers.Register)
		organization.POST("/login", controllers.Login)
	}

	admin := router.Group("/api/admin")
	{
		admin.GET("/", utils.TokenCheckedFromHeader, controllers.GetAdmins)
		admin.GET("/:id", utils.TokenCheckedFromHeader, controllers.GetAdminById)
		admin.POST("/create", utils.TokenCheckedFromHeader, controllers.CreateAdmin)
		admin.PATCH("/update/:id", utils.TokenCheckedFromHeader, controllers.UpdateAdmin)
		admin.DELETE("/delete/:id", utils.TokenCheckedFromHeader, controllers.DeleteAdmin)
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

package routes

import (
	"github.com/Zefirnutiy/sweet_potato.git/controllers"
	"github.com/gin-gonic/gin"
)

func Routs(port string) {
	router := gin.Default()

	auth := router.Group("/api/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	router.Run(port)

}

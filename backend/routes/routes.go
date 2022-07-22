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

	client := router.Group("/api/client")
	{
		client.POST("/register", controllers.Register)
		client.POST("/login", controllers.Login)
	}

	deadLine := router.Group("/api/deadline")
	{
		deadLine.GET("/", controllers.GetDeadLine)
		deadLine.GET("/:id", controllers.GetDeadLineById)
		deadLine.POST("/create", controllers.CreateDeadLine)
		deadLine.PATCH("/update", controllers.UpdateDeadLine)
		deadLine.DELETE("/delete", controllers.DeleteDeadLine)
	}
	router.Run(port)

}

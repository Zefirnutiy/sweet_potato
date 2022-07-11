package routes

import (
	"github.com/gin-gonic/gin"
)

func Routs(port string) {
	router := gin.Default()

	auth := router.Group("/api/auth")
	{
		auth.POST("/register")
		auth.POST("/login")
	}

	router.Run(port)

}

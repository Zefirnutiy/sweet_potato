package routes

import (
	"github.com/gin-gonic/gin"
)

func Routs(port string) {
	router := gin.Default()
	router.LoadHTMLGlob("../frontend/templates/*.html")
	router.Static("/js/", "../frontend/js")
	router.Static("/css/", "../frontend/css")

	auth := router.Group("/api/auth")
	{
		auth.POST("/register")
		auth.POST("/login")
	}

	router.Run(port)

}

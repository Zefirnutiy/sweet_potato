package main

import (
	"github.com/Zefirnutiy/sweet_potato.git/controllers"
	"github.com/gin-gonic/gin"
)

func app(port string) {
	router := gin.Default()

	router.GET("/test", controllers.TestIndex)

	router.Run(port)

}

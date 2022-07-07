package routs

import (
	"github.com/gin-gonic/gin"
)

func Routs(port string) {
	router := gin.Default()

	router.Run(port)

}

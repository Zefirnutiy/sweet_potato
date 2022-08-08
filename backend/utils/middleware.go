package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//Нужна для передачи данных и проверке, авторизован ли человек или нет
func TokenCheckedFromHeader(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Авторизуйтесь",
		})
		return
	}

	header := strings.Split(authHeader, " ")

	if len(header) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Авторизуйтесь",
		})
		return
	}

	if header[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Авторизуйтесь",
		})
		return
	}

	Model, err := ParseToken(header[1], []byte("woe_ifhb_o$wei_$rhb#fqol#eri@bf"))

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Авторизуйтесь",
		})
		return
	}

	c.Set("Model", Model)
}

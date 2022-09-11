package utils

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

//Нужна для передачи данных и проверке, авторизован ли человек или нет
func TokenCheckedFromHeader(c *gin.Context) {
	authHeader := c.Request.Header[`Authorization`]
	
	if authHeader[0] == ""{
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Авторизуйтесь",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	authHeaderMas := strings.Split(authHeader[0], " ") 

	if len(authHeaderMas) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Авторизуйтесь",
		})
		return
	}

	if authHeaderMas[0] != `Bearer` {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Авторизуйтесь",
		})
		return
	}

	Model, err := ParseToken(authHeaderMas[1], []byte(os.Getenv("SECRET_WORD_FOR_ORG")))

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Авторизуйтесь",
		})
		return
	}

	c.Set("Model", Model)
}

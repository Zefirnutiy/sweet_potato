package controllers

import "github.com/gin-gonic/gin"

func Register(c *gin.Context) {
	c.HTML(200, "registration.html", nil)
}

func Login(c *gin.Context) {
	//
}

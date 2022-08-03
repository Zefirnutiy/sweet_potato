package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context) {
	Organization := c.Keys["Organization"].(structs.Organization)
	c.JSON(http.StatusOK, gin.H{
		"title": Organization.Title,
		"email": Organization.Email,
	})
}

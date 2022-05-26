package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/gin-gonic/gin"
)

func TestIndex(c *gin.Context) {
	db.InsertQuest()
	c.JSON(200, gin.H{
		"lol": "kik",
	})
}

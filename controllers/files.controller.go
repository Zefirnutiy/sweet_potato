package controllers

import (
	// "github.com/Zefirnutiy/sweet_potato.git/db"
	"fmt"

	"github.com/gin-gonic/gin"
)


func GetFiles(c *gin.Context){

}

func UploadFile(c *gin.Context){
	file, _ := c.FormFile("file")

	fmt.Println(file.Filename)
}

func DeleteFile(c *gin.Context){
	// id := c.Param("id")
}
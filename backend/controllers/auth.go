package controllers

import (
	"fmt"
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context) {
	var organization structs.Organization

	if err := c.ShouldBindJSON(&organization); err != nil {
		c.String(http.StatusBadGateway, err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(organization.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	candidate := db.SelectOrganization(organization.Email)

	if candidate == true {
		fmt.Println(candidate)
		c.JSON(http.StatusConflict, gin.H{
			"message": "Такой пользователь уже существует",
		})
		return
	}

	_, err = db.QuestionRegister(organization.Title, string(hashedPassword), organization.Email, organization.EmailNotifications)

	if err != nil {
		c.String(http.StatusNotImplemented, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"organization": organization,
		"message":      "Организация успешно создана",
	})
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

package controllers

import (
	"database/sql"
	"fmt"
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func QuestionRegister(title, password, email string, emailnotifications bool) (sql.Result, error) {

	result, err := db.Dbpool.Exec(`
		INSERT INTO main."Organization" (title, password, email, emailnotifications) 
		VALUES ($1, $2, $3, $4);`,
		title, password, email, emailnotifications)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func SelectOrganization(email string) (structs.Organization, bool) {
	var organization structs.Organization

	err := db.Dbpool.QueryRow(`SELECT * FROM main."Organization" WHERE email IN ($1)`, email).Scan(
		&organization.Id,
		&organization.Title,
		&organization.Password,
		&organization.Email,
		&organization.EmailNotifications,
		&organization.LevelId,
	)

	if err != nil && err == sql.ErrNoRows {
		return structs.Organization{}, false
	}

	return organization, true
}

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

	_, candidate := SelectOrganization(organization.Email)

	if candidate == true {
		fmt.Println(candidate)
		c.JSON(http.StatusConflict, gin.H{
			"message": "Такой пользователь уже существует",
		})
		return
	}

	_, err = QuestionRegister(organization.Title, string(hashedPassword), organization.Email, organization.EmailNotifications)

	if err != nil {
		c.String(http.StatusNotImplemented, err.Error())
		return
	}

	token, err := utils.CreateToken(organization.Title)

	c.JSON(http.StatusCreated, gin.H{
		"organization": organization,
		"token":        token,
		"message":      "Организация успешно создана",
	})
}

func Login(c *gin.Context) {
	var organization structs.Organization

	if err := c.ShouldBindJSON(&organization); err != nil {
		c.String(http.StatusBadGateway, err.Error())
		return
	}

	result, yes := SelectOrganization(organization.Email)

	if !yes {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Вы ввели неправильные данные 1",
		})
		return
	}

	resultPasswordByte := []byte(result.Password)
	getPasswordByte := []byte(organization.Password)

	err := bcrypt.CompareHashAndPassword(resultPasswordByte, getPasswordByte)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     err,
			"message": "Вы ввели неправильные данные",
		})
		return
	}

	token, err := utils.CreateToken(result.Title)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Что-то пошло не так",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Вы успешно вошли",
		"token":   token,
	})
}

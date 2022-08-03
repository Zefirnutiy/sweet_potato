package controllers

import (
	"database/sql"
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func CreateOrganization(title, password, email string, emailnotifications bool) (sql.Result, error) {
	result, err := db.Dbpool.Exec(`
		INSERT INTO "Main"."Organization" ("Title", "Password", "Email", "EmailNotifications") 
		VALUES ($1, $2, $3, $4);`,
		title, password, email, emailnotifications)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetOrganization(email string) (structs.Organization, bool) {
	var organization structs.Organization

	err := db.Dbpool.QueryRow(`SELECT * FROM "Main"."Organization" WHERE "Email" IN ($1)`, email).Scan(
		&organization.Id,
		&organization.Title,
		&organization.Password,
		&organization.Email,
		&organization.EmailNotifications,
		&organization.LevelId,
	)

	if err != nil {
		return structs.Organization{}, false
	}

	return organization, true
}

func RegisterOrganization(c *gin.Context) {
	var organization structs.Organization

	if err := c.ShouldBindJSON(&organization); err != nil {
		c.String(http.StatusBadGateway, err.Error())
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(organization.Password), bcrypt.DefaultCost)

	_, candidate := GetOrganization(organization.Email)

	if candidate {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Такой пользователь уже существует",
		})
		return
	}

	_, err := CreateOrganization(
		organization.Title,
		string(hashedPassword),
		organization.Email,
		organization.EmailNotifications)

	if err != nil {
		c.String(http.StatusNotImplemented, err.Error(), gin.H{
			"message": "Что-то пошло не так",
		})
		return
	}

	token, err := utils.CreateToken(organization, structs.Client{})

	if err != nil {
		c.String(500, err.Error())
		return
	}

	organizationFromDb, _ := GetOrganization(organization.Email)
	organizationFromDb.Password = ""

	c.JSON(http.StatusCreated, gin.H{
		"organization": organizationFromDb,
		"token":        token,
		"message":      "Организация успешно создана",
	})
}

func LoginOrganization(c *gin.Context) {
	var organization structs.Organization

	if err := c.ShouldBindJSON(&organization); err != nil {
		c.String(http.StatusBadGateway, err.Error())
		return
	}

	result, candidate := GetOrganization(organization.Email)

	if !candidate {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Вы ввели неправильные данные",
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

	result.Password = ""
	token, err := utils.CreateToken(result, structs.Client{})

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Что-то пошло не так",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"organization": result,
		"message":      "Вы успешно вошли",
		"token":        token,
	})
}

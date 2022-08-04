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

func CreateOrganization(title, password, email, key string, emailnotifications bool) (sql.Result, error) {
	result, err := db.Dbpool.Exec(`
		INSERT INTO "Main"."Organization" ("Title", "Password", "Email", "EmailNotifications", "Key") 
		VALUES ($1, $2, $3, $4, $5);`,
		title, password, email, emailnotifications, key)

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
		&organization.Key,
	)

	if err != nil {
		fmt.Println(err)
		return structs.Organization{}, false
	}

	return organization, true
}

func RegisterOrganization(c *gin.Context) {
	var organization structs.Organization
	var claimOrganization structs.Claims

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
		string(hashedPassword[len(hashedPassword)-6:]),
		organization.EmailNotifications,
	)
	db.CreateTable("./db/organization.sql", string(hashedPassword[len(hashedPassword)-6:]))

	if err != nil {
		c.String(http.StatusNotImplemented, err.Error(), gin.H{
			"message": "Что-то пошло не так",
		})
		return
	}
	organizationFromDb, _ := GetOrganization(organization.Email)
	claimOrganization.Id = organization.Id
	claimOrganization.Email = organization.Email
	claimOrganization.Schema = organizationFromDb.Key

	token, err := utils.CreateToken(claimOrganization)

	if err != nil {
		c.String(500, err.Error())
		return
	}

	organizationFromDb.Password = ""

	c.JSON(http.StatusCreated, gin.H{
		"organization": organizationFromDb,
		"token":        token,
		"message":      "Организация успешно создана",
	})
}

func LoginOrganization(c *gin.Context) {
	var organization structs.Organization
	var claimOrganization structs.Claims

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

	claimOrganization.Id = result.Id
	claimOrganization.Email = result.Email
	claimOrganization.Schema = result.Key
	token, err := utils.CreateToken(claimOrganization)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Что-то пошло не так",
		})
	}
	result.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"organization": result,
		"message":      "Вы успешно вошли",
		"token":        token,
	})
}

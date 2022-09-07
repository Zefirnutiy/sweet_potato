package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateOrganization(org structs.Organization) *sql.Row {
	result := db.Dbpool.QueryRow(`INSERT INTO "public"."Organization" ("Title", 
	"Password", "Email", "EmailNotifications", "Key", "UserLimit", "Statistics", 
	"ProtectionCheating", "Date", "Time", "ThemeId") 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING "Id";`,
		org.Title, 
		org.Password, 
		org.Email, 
		org.EmailNotifications, 
		org.Key,  
		org.UserLimit, 
		org.Statistics, 
		org.ProtectionCheating, 
		org.Date, 
		org.Time,
		org.ThemeId)

	return result
}

func GetOrganization(email string) (structs.Organization, bool) {
	var organization structs.Organization

	err := db.Dbpool.QueryRow(`SELECT * FROM public."Organization" WHERE "Email" IN ($1)`, email).Scan(
		&organization.Id,
		&organization.Title,
		&organization.Password,
		&organization.Email,
		&organization.EmailNotifications,
		&organization.Key,
		&organization.UserLimit,
		&organization.Statistics,
		&organization.ProtectionCheating,
		&organization.Date,
		&organization.Time,
		&organization.ThemeId,
	)

	if err != nil {
		fmt.Println(err)
		return structs.Organization{}, false
	}

	return organization, true
}

func RegisterOrganization(c *gin.Context) {
	var (
		organization structs.Organization
	 	claimOrganization structs.Claims
		id int
	)

	if err := c.ShouldBindJSON(&organization); err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}

	
	_, isExist := GetOrganization(organization.Email)
	
	if isExist {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Такой пользователь уже существует",
		})
		return
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(organization.Password), bcrypt.DefaultCost)

	res := CreateOrganization(organization)

	organization.Key = string(hashedPassword[len(hashedPassword)-6:])
	res.Scan(&id)
	
	if err := db.CreateTable("./db/organization.sql", string(hashedPassword[len(hashedPassword)-6:])); err != nil {
		c.String(http.StatusNotImplemented, err.Error(), gin.H{
			"message": "Что-то пошло не так",
		})
		return
	}
	claimOrganization.Id = id
	claimOrganization.Email = organization.Email
	claimOrganization.KeySchema = organization.Key

	token, err := utils.CreateToken(claimOrganization)

	if err != nil {
		c.String(500, err.Error())
		return
	}

	organization.Password = ""

	c.JSON(http.StatusCreated, gin.H{
		"organization": organization,
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
	claimOrganization.KeySchema = result.Key
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

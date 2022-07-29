package controllers

import (
	"fmt"
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessing(c gin.Context) structs.Client {
	var data structs.Client
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"message": "Некорректные данные",
		})
		return structs.Client{}
	}
	return data
}


func GetClients(c *gin.Context) {
	schema := c.GetString("schema")
	var clientList []structs.Client
	var client structs.Client

	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	rows, err := db.Dbpool.Query(`SELECT * FROM "` + "%[1]s" + `"."Client"`, schema)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"result": nil,
			"message": "Ничего не найдено",
		})
		return
	}

	for rows.Next() {
		err = rows.Scan(
			&client.Id,
			&client.FirstName,
			&client.LastName,
			&client.Patronymic,
			&client.Password,
			&client.Email,
			&client.Telephone,
			&client.EmailNotifications,
			&client.GroupId,
			&client.Organization)
		clientList = append(clientList, client)
		if err != nil {
			utils.Logger.Println(err)
			c.JSON(500, gin.H{
				"result": nil,
				"message": "Ошибка сервера",
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"result": clientList,
		"message": nil,
	})
}

func GetClientById(c *gin.Context) {
	schema := c.GetString("schema")
	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	id := c.Params.ByName("id")
	var client structs.Client

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+ "%[2]s"+`"."Client" WHERE "Id"=$1`, id, schema).Scan(
		&client.Id,
		&client.FirstName,
		&client.LastName,
		&client.Patronymic,
		&client.Password,
		&client.Email,
		&client.Telephone,
		&client.EmailNotifications,
		&client.GroupId,
		&client.Organization,
	)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"result": nil,
			"message": "Ничего не найдено",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": client,
		"message": nil,
	})
}


func GetClientByGroupId(c *gin.Context) {
	schema := c.GetString("schema")
	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	id := c.Params.ByName("id")
	var client structs.Client

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+ "%[2]s"+`"."Client" WHERE "GroupId"=$1`, id, schema).Scan(
		&client.Id,
		&client.FirstName,
		&client.LastName,
		&client.Patronymic,
		&client.Password,
		&client.Email,
		&client.Telephone,
		&client.EmailNotifications,
		&client.GroupId,
		&client.Organization,
	)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"result": nil,
			"message": "Ничего не найдено",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": client,
		"message": nil,
	})
}

func CreateClient(c *gin.Context) {
	data := DataProcessing(*c)
	encryptedPassword, err := utils.Encrypt(data.Password)
	if err != nil {
		utils.Logger.Println(data.Email, err)
		c.JSON(500, gin.H{
			"message": "Ошибка сервера",
		})
		return
	}

	_, err = db.Dbpool.Exec(`INSERT INTO "`+data.Organization+`"."Client"
							("FirstName", 
							"LastName", 
							"Patronymic", 
							"Password", 
							"Email", 
							"Telephone", 
							"EmailNotifications", 
							"GroupId", 
							"Organization") 
							VALUES( $1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		data.FirstName,
		data.LastName,
		data.Patronymic,
		encryptedPassword,
		data.Email,
		data.Telephone,
		data.EmailNotifications,
		data.GroupId,
		data.Organization)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"message": "Не получилось записать данные",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Запись создана",
	})
}

func UpdateClient(c *gin.Context) {
	schema := c.GetString("title")
	fmt.Println(schema)
	data := DataProcessing(*c)
	encryptedPassword, err := utils.Encrypt(data.Password)
	if err != nil {
		utils.Logger.Println(data.Email, err)
		c.JSON(500, gin.H{
			"message": "Ошибка сервера",
		})
		return
	}
	
	_, err = db.Dbpool.Exec(`UPDATE "`+"%[10]s"+`"."Client" 
		SET "FirstName"=$1 ,
		"LastName"=$2,
		"Patronymic"=$3, 
		"Password"=$4, 
		"Email"=$5, 
		"Telephone"=$6, 
		"EmailNotifications"=$7, 
		"GroupId"=$8
		WHERE "Id"=$9`,
		data.FirstName,
		data.LastName,
		data.Patronymic,
		encryptedPassword,
		data.Email,
		data.Telephone,
		data.EmailNotifications,
		data.GroupId,
		data.Id,
		schema)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"message": "Не получилось обновить данные",
		})
		return
	}
	c.JSON(200, gin.H{
		"result": "Данные обновлены",
	})
}

func DeleteClient(c *gin.Context) {
	schema := c.GetString("schema")
	tokken := c.Params.ByName("tokken")
	id := c.Params.ByName("id") 
	fmt.Println(tokken)
	_, err := db.Dbpool.Exec(`DELETE FROM "`+"%[2]s"+`"."Client" WHERE "Id"=$1`, id, schema)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"message": "Не получилось удалить данные",
		})
		return
	}
	c.JSON(200, gin.H{
		"result": "Данные удалены",
	})
}
	
	
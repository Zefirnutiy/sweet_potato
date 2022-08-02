
package controllers

import (
	"fmt"
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingClient(c gin.Context) structs.Client {
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

	rows, err := db.Dbpool.Query(`SELECT * FROM "` + schema + `"."Client"`, schema)
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
		)
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

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+schema+`"."Client" WHERE "Id"=$1`, id ).Scan(
		&client.Id, 
		&client.FirstName, 
		&client.LastName, 
		&client.Patronymic, 
		&client.Password, 
		&client.Email, 
		&client.Telephone, 
		&client.EmailNotifications, 
		&client.GroupId, 
		
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

	groupId := c.Params.ByName("groupId")
	var client structs.Client

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+schema+`"."Client" WHERE "GroupId"=$1`, groupId ).Scan(
		&client.Id, 
		&client.FirstName, 
		&client.LastName, 
		&client.Patronymic, 
		&client.Password, 
		&client.Email, 
		&client.Telephone, 
		&client.EmailNotifications, 
		&client.GroupId, 
		
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
	schema := c.GetString("schema")
	data := DataProcessingClient(*c)
	var err error
	
	data.Password, err = utils.Encrypt(data.Password)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"message": "Ошибка сервера",
		})
		return
	}
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+schema+`"."Client"
		(
		"FirstName", 
		"LastName", 
		"Patronymic", 
		"Email", 
		"Telephone", 
		"EmailNotifications", 
		"GroupId", 
		
		) 
		VALUES( $1, $2, $3, $4, $5, $6, $7 )`,
		data.FirstName, 
		data.LastName, 
		data.Patronymic, 
		data.Email, 
		data.Telephone, 
		data.EmailNotifications, 
		data.GroupId, 
		)
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

	schema := c.GetString("schema")
	data := DataProcessingClient(*c)
	var err error
	
	data.Password, err = utils.Encrypt(data.Password)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"message": "Ошибка сервера",
		})
		return
	}
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+schema+`"."Client" 
		SET 
		"FirstName"=$2,
		"LastName"=$3,
		"Patronymic"=$4,
		"Password"=$5,
		"Email"=$6,
		"Telephone"=$7,
		"EmailNotifications"=$8,
		"GroupId"=$9
		
		WHERE "Id"=$1`,
		data.Id, 
		data.FirstName, 
		data.LastName, 
		data.Patronymic, 
		data.Password, 
		data.Email, 
		data.Telephone, 
		data.EmailNotifications, 
		data.GroupId, 
		
		)
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
	_, err := db.Dbpool.Exec(`DELETE FROM "`+schema+`"."Client" WHERE "Id"=$1`, id)
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
	
	
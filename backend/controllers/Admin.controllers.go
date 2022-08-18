package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingAdmin(c gin.Context) structs.Admin {
	var data structs.Admin
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"message": "Некорректные данные",
		})
		return structs.Admin{}
	}
	return data
}


func GetAdmins(c *gin.Context) {
	var adminList []structs.Admin
	var admin structs.Admin

	rows, err := db.Dbpool.Query(`SELECT * FROM "Admin"`)
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
		&admin.Id, 
		&admin.FirstName, 
		&admin.LastName, 
		&admin.Email, 
		&admin.Password, 
		)
		adminList = append(adminList, admin)
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
		"result": adminList,
		"message": nil,
	})
}

func GetAdminById(c *gin.Context) {

	id := c.Params.ByName("id")
	var admin structs.Admin

	err := db.Dbpool.QueryRow(`SELECT * FROM "Admin" WHERE "Id"=$1`, id ).Scan(
		&admin.Id, 
		&admin.FirstName, 
		&admin.LastName, 
		&admin.Email, 
		&admin.Password, 
		
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
		"result": admin,
		"message": nil,
	})
}

	


func CreateAdmin(c *gin.Context) {
	data := DataProcessingAdmin(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "Admin"
		(
		"Id", 
		"FirstName", 
		"LastName", 
		"Email", 
		
		) 
		VALUES( $1, $2, $3, $4 )`,
		data.Id, 
		data.FirstName, 
		data.LastName, 
		data.Email, 
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
	

func UpdateAdmin(c *gin.Context) {

	data := DataProcessingAdmin(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "Admin" 
		SET 
		"FirstName"=$2,
		"LastName"=$3,
		"Email"=$4,
		"Password"=$5
		
		WHERE "Id"=$1`,
		data.Id, 
		data.FirstName, 
		data.LastName, 
		data.Email, 
		data.Password, 
		
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
	

func DeleteAdmin(c *gin.Context) {
	id := c.Params.ByName("id") 
	_, err := db.Dbpool.Exec(`DELETE FROM "Admin" WHERE "Id"=$1`, id)
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
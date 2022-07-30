
package controllers

import (
	"fmt"
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingDepartment(c gin.Context) structs.Department {
	var data structs.Department
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"message": "Некорректные данные",
		})
		return structs.Department{}
	}
	return data
}


func GetDepartments(c *gin.Context) {
	schema := c.GetString("schema")
	var departmentList []structs.Department
	var department structs.Department

	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	rows, err := db.Dbpool.Query(`SELECT * FROM "` + schema + `"."Department"`, schema)
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
		&department.Id, 
		&department.Title, 
		)
		departmentList = append(departmentList, department)
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
		"result": departmentList,
		"message": nil,
	})
}

func GetDepartmentById(c *gin.Context) {
	schema := c.GetString("schema")
	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	id := c.Params.ByName("id")
	var department structs.Department

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+schema+`"."Department" WHERE "Id"=$1`, id ).Scan(
		&department.Id, 
		&department.Title, 
		
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
		"result": department,
		"message": nil,
	})
}

	

func CreateDepartment(c *gin.Context) {
	schema := c.GetString("schema")
	data := DataProcessingDepartment(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+schema+`"."Department"
		(
		"Title", 
		
		) 
		VALUES( $1 )`,
		data.Title, 
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
	

func UpdateDepartment(c *gin.Context) {

	schema := c.GetString("schema")
	data := DataProcessingDepartment(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+schema+`"."Department" 
		SET 
		"Title"=$2
		
		WHERE "Id"=$1`,
		data.Id, 
		data.Title, 
		
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
	

func DeleteDepartment(c *gin.Context) {
	schema := c.GetString("schema")
	tokken := c.Params.ByName("tokken")
	id := c.Params.ByName("id") 
	fmt.Println(tokken)
	_, err := db.Dbpool.Exec(`DELETE FROM "`+schema+`"."Department" WHERE "Id"=$1`, id)
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
	
	
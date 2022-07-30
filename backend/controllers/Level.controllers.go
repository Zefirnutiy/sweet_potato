
package controllers

import (
	"fmt"
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingLevel(c gin.Context) structs.Level {
	var data structs.Level
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"message": "Некорректные данные",
		})
		return structs.Level{}
	}
	return data
}


func GetLevels(c *gin.Context) {
	schema := c.GetString("schema")
	var levelList []structs.Level
	var level structs.Level

	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	rows, err := db.Dbpool.Query(`SELECT * FROM "` + schema + `"."Level"`, schema)
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
		&level.Id, 
		&level.Title, 
		&level.Price, 
		&level.Paid, 
		&level.CreateCourse, 
		&level.TakeCourse, 
		&level.AploadFile, 
		&level.ViewYourResult, 
		&level.ViewOtherResults, 
		)
		levelList = append(levelList, level)
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
		"result": levelList,
		"message": nil,
	})
}

func GetLevelById(c *gin.Context) {
	schema := c.GetString("schema")
	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	id := c.Params.ByName("id")
	var level structs.Level

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+schema+`"."Level" WHERE "Id"=$1`, id ).Scan(
		&level.Id, 
		&level.Title, 
		&level.Price, 
		&level.Paid, 
		&level.CreateCourse, 
		&level.TakeCourse, 
		&level.AploadFile, 
		&level.ViewYourResult, 
		&level.ViewOtherResults, 
		
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
		"result": level,
		"message": nil,
	})
}

	

func CreateLevel(c *gin.Context) {
	schema := c.GetString("schema")
	data := DataProcessingLevel(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+schema+`"."Level"
		(
		"Title", 
		"Price", 
		"Paid", 
		"CreateCourse", 
		"TakeCourse", 
		"AploadFile", 
		"ViewYourResult", 
		"ViewOtherResults", 
		
		) 
		VALUES( $1, $2, $3, $4, $5, $6, $7, $8 )`,
		data.Title, 
		data.Price, 
		data.Paid, 
		data.CreateCourse, 
		data.TakeCourse, 
		data.AploadFile, 
		data.ViewYourResult, 
		data.ViewOtherResults, 
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
	

func UpdateLevel(c *gin.Context) {

	schema := c.GetString("schema")
	data := DataProcessingLevel(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+schema+`"."Level" 
		SET 
		"Title"=$2,
		"Price"=$3,
		"Paid"=$4,
		"CreateCourse"=$5,
		"TakeCourse"=$6,
		"AploadFile"=$7,
		"ViewYourResult"=$8,
		"ViewOtherResults"=$9
		
		WHERE "Id"=$1`,
		data.Id, 
		data.Title, 
		data.Price, 
		data.Paid, 
		data.CreateCourse, 
		data.TakeCourse, 
		data.AploadFile, 
		data.ViewYourResult, 
		data.ViewOtherResults, 
		
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
	

func DeleteLevel(c *gin.Context) {
	schema := c.GetString("schema")
	tokken := c.Params.ByName("tokken")
	id := c.Params.ByName("id") 
	fmt.Println(tokken)
	_, err := db.Dbpool.Exec(`DELETE FROM "`+schema+`"."Level" WHERE "Id"=$1`, id)
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
	
	
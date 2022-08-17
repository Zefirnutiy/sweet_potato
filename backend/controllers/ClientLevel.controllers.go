
package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingClientLevel(c gin.Context) structs.ClientLevel {
	var data structs.ClientLevel
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.ClientLevel{}
	}
	return data
}


func GetClientLevels(c *gin.Context) {
	schema := c.Params.ByName("schema")
	var clientLevelList []structs.ClientLevel
	var clientLevel structs.ClientLevel

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+schema+`"."ClientLevel"`)
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
		&clientLevel.Id, 
		&clientLevel.Title, 
		&clientLevel.CreateCourse, 
		&clientLevel.TakeCourse, 
		&clientLevel.AploadFile, 
		&clientLevel.ViewYourResult, 
		&clientLevel.ViewOtherResults, 
		)
		clientLevelList = append(clientLevelList, clientLevel)
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
		"result": clientLevelList,
		"message": nil,
	})
}

func GetClientLevelById(c *gin.Context) {
	schema := c.Params.ByName("schema")
	id := c.Params.ByName("id")
	var clientLevel structs.ClientLevel

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+schema+`"."ClientLevel" WHERE "Id"=$1`, id ).Scan(
		&clientLevel.Id, 
		&clientLevel.Title, 
		&clientLevel.CreateCourse, 
		&clientLevel.TakeCourse, 
		&clientLevel.AploadFile, 
		&clientLevel.ViewYourResult, 
		&clientLevel.ViewOtherResults, 
		
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
		"result": clientLevel,
		"message": nil,
	})
}


func CreateClientLevel(c *gin.Context) {
	schema := c.Params.ByName("schema")
	data := DataProcessingClientLevel(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+schema+`"."ClientLevel"
		(
		"Title", 
		"CreateCourse", 
		"TakeCourse", 
		"AploadFile", 
		"ViewYourResult", 
		"ViewOtherResults", 
		
		) 
		VALUES( $1, $2, $3, $4, $5, $6 )`,
		data.Title, 
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
	

func UpdateClientLevel(c *gin.Context) {

	schema := c.Params.ByName("schema")
	id := c.Params.ByName("id")
	data := DataProcessingClientLevel(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+schema+`"."ClientLevel" 
		SET 
		"Title"=$1,
		"CreateCourse"=$2,
		"TakeCourse"=$3,
		"AploadFile"=$4,
		"ViewYourResult"=$5,
		"ViewOtherResults"=$6
		
		WHERE "Id"=$1`,
		id,
		data.Title, 
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
	

func DeleteClientLevel(c *gin.Context) {
	schema := c.Params.ByName("schema")
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+schema+`"."ClientLevel" WHERE "Id"=$1`, id)
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
	
	
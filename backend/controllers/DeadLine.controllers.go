
package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingDeadLine(c gin.Context) structs.DeadLine {
	var data structs.DeadLine
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.DeadLine{}
	}
	return data
}


func GetDeadLines(c *gin.Context) {
	var deadLineList []structs.DeadLine
	var deadLine structs.DeadLine

	rows, err := db.Dbpool.Query(`SELECT * FROM "DeadLine"`)
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
		&deadLine.Id, 
		&deadLine.Date, 
		&deadLine.LevelId, 
		&deadLine.OrganizationId, 
		)
		deadLineList = append(deadLineList, deadLine)
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
		"result": deadLineList,
		"message": nil,
	})
}

func GetDeadLineById(c *gin.Context) {
	id := c.Params.ByName("id")
	var deadLine structs.DeadLine

	err := db.Dbpool.QueryRow(`SELECT * FROM "DeadLine" WHERE "Id"=$1`, id ).Scan(
		&deadLine.Id, 
		&deadLine.Date, 
		&deadLine.LevelId, 
		&deadLine.OrganizationId, 
		
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
		"result": deadLine,
		"message": nil,
	})
}

	
func GetDeadLineByOrganizationId(c *gin.Context) {
	organizationId := c.Params.ByName("organizationId")
	var deadLine structs.DeadLine

	err := db.Dbpool.QueryRow(`SELECT * FROM "DeadLine" WHERE "OrganizationId"=$1`, organizationId ).Scan(
		&deadLine.Id, 
		&deadLine.Date, 
		&deadLine.LevelId, 
		&deadLine.OrganizationId, 
		
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
		"result": deadLine,
		"message": nil,
	})
}

	


func CreateDeadLine(c *gin.Context) {
	data := DataProcessingDeadLine(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "DeadLine"
		(
		"Date", 
		"LevelId", 
		"OrganizationId", 
		
		) 
		VALUES( $1, $2, $3 )`,
		data.Date, 
		data.LevelId, 
		data.OrganizationId, 
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
	

func UpdateDeadLine(c *gin.Context) {

	id := c.Params.ByName("id")
	data := DataProcessingDeadLine(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "DeadLine" 
		SET 
		"Date"=$1,
		"LevelId"=$2,
		"OrganizationId"=$3
		
		WHERE "Id"=$1`,
		id,
		data.Date, 
		data.LevelId, 
		data.OrganizationId, 
		
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
	

func DeleteDeadLine(c *gin.Context) {
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "DeadLine" WHERE "Id"=$1`, id)
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
	
	
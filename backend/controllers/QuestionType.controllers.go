
package controllers

import (
	"fmt"
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingQuestionType(c gin.Context) structs.QuestionType {
	var data structs.QuestionType
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"message": "Некорректные данные",
		})
		return structs.QuestionType{}
	}
	return data
}


func GetQuestionTypes(c *gin.Context) {
	schema := c.GetString("schema")
	var questionTypeList []structs.QuestionType
	var questionType structs.QuestionType

	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	rows, err := db.Dbpool.Query(`SELECT * FROM "` + schema + `"."QuestionType"`, schema)
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
		&questionType.Id, 
		&questionType.Type, 
		)
		questionTypeList = append(questionTypeList, questionType)
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
		"result": questionTypeList,
		"message": nil,
	})
}

func GetQuestionTypeById(c *gin.Context) {
	schema := c.GetString("schema")
	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	id := c.Params.ByName("id")
	var questionType structs.QuestionType

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+schema+`"."QuestionType" WHERE "Id"=$1`, id ).Scan(
		&questionType.Id, 
		&questionType.Type, 
		
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
		"result": questionType,
		"message": nil,
	})
}

	

func CreateQuestionType(c *gin.Context) {
	schema := c.GetString("schema")
	data := DataProcessingQuestionType(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+schema+`"."QuestionType"
		(
		"Type", 
		
		) 
		VALUES( $1 )`,
		data.Type, 
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
	

func UpdateQuestionType(c *gin.Context) {

	schema := c.GetString("schema")
	data := DataProcessingQuestionType(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+schema+`"."QuestionType" 
		SET 
		"Type"=$2
		
		WHERE "Id"=$1`,
		data.Id, 
		data.Type, 
		
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
	

func DeleteQuestionType(c *gin.Context) {
	schema := c.GetString("schema")
	tokken := c.Params.ByName("tokken")
	id := c.Params.ByName("id") 
	fmt.Println(tokken)
	_, err := db.Dbpool.Exec(`DELETE FROM "`+schema+`"."QuestionType" WHERE "Id"=$1`, id)
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
	
	
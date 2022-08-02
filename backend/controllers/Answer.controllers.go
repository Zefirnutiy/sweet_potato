
package controllers

import (
	"fmt"
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingAnswer(c gin.Context) structs.Answer {
	var data structs.Answer
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"message": "Некорректные данные",
		})
		return structs.Answer{}
	}
	return data
}


func GetAnswers(c *gin.Context) {
	schema := c.GetString("schema")
	var answerList []structs.Answer
	var answer structs.Answer

	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	rows, err := db.Dbpool.Query(`SELECT * FROM "` + schema + `"."Answer"`, schema)
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
		&answer.Id, 
		&answer.Text, 
		&answer.Correct, 
		&answer.QuestionId, 
		)
		answerList = append(answerList, answer)
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
		"result": answerList,
		"message": nil,
	})
}

func GetAnswerById(c *gin.Context) {
	schema := c.GetString("schema")
	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	id := c.Params.ByName("id")
	var answer structs.Answer

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+schema+`"."Answer" WHERE "Id"=$1`, id ).Scan(
		&answer.Id, 
		&answer.Text, 
		&answer.Correct, 
		&answer.QuestionId, 
		
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
		"result": answer,
		"message": nil,
	})
}

	


func CreateAnswer(c *gin.Context) {
	schema := c.GetString("schema")
	data := DataProcessingAnswer(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+schema+`"."Answer"
		(
		"Text", 
		"Correct", 
		"QuestionId", 
		
		) 
		VALUES( $1, $2, $3 )`,
		data.Text, 
		data.Correct, 
		data.QuestionId, 
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
	

func UpdateAnswer(c *gin.Context) {

	schema := c.GetString("schema")
	data := DataProcessingAnswer(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+schema+`"."Answer" 
		SET 
		"Text"=$2,
		"Correct"=$3,
		"QuestionId"=$4
		
		WHERE "Id"=$1`,
		data.Id, 
		data.Text, 
		data.Correct, 
		data.QuestionId, 
		
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
	

func DeleteAnswer(c *gin.Context) {
	schema := c.GetString("schema")
	tokken := c.Params.ByName("tokken")
	id := c.Params.ByName("id") 
	fmt.Println(tokken)
	_, err := db.Dbpool.Exec(`DELETE FROM "`+schema+`"."Answer" WHERE "Id"=$1`, id)
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
	
	
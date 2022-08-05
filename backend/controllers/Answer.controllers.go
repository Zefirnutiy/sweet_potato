
package controllers

import (
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
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.Answer{}
	}
	return data
}


func GetAnswers(c *gin.Context) {
	schema := c.Params.ByName("schema")
	var answerList []structs.Answer
	var answer structs.Answer

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+schema+`"."Answer"`)
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
	schema := c.Params.ByName("schema")
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

	

func GetAnswerByQuestionId(c *gin.Context) {
	schema := c.Params.ByName("schema")
	questionId := c.Params.ByName("questionId")
	var answerList []structs.Answer
	var answer structs.Answer

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+schema+`"."Answer" WHERE "QuestionId"=$1`, questionId )
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

	

func CreateAnswer(c *gin.Context) {
	schema := c.Params.ByName("schema")
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

	schema := c.Params.ByName("schema")
	id := c.Params.ByName("id")
	data := DataProcessingAnswer(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+schema+`"."Answer" 
		SET 
		"Text"=$1,
		"Correct"=$2,
		"QuestionId"=$3
		
		WHERE "Id"=$1`,
		id,
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
	schema := c.Params.ByName("schema")
	id := c.Params.ByName("id")
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
	
	
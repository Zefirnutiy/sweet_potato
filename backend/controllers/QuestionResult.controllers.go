
package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingQuestionResult(c gin.Context) structs.QuestionResult {
	var data structs.QuestionResult
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.QuestionResult{}
	}
	return data
}


func GetQuestionResults(c *gin.Context) {
	schema := c.Params.ByName("schema")
	var questionResultList []structs.QuestionResult
	var questionResult structs.QuestionResult

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+schema+`"."QuestionResult"`)
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
		&questionResult.Id, 
		&questionResult.Date, 
		&questionResult.QuestionId, 
		&questionResult.ClientId, 
		&questionResult.Scores, 
		)
		questionResultList = append(questionResultList, questionResult)
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
		"result": questionResultList,
		"message": nil,
	})
}

func GetQuestionResultById(c *gin.Context) {
	schema := c.Params.ByName("schema")
	id := c.Params.ByName("id")
	var questionResult structs.QuestionResult

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+schema+`"."QuestionResult" WHERE "Id"=$1`, id ).Scan(
		&questionResult.Id, 
		&questionResult.Date, 
		&questionResult.QuestionId, 
		&questionResult.ClientId, 
		&questionResult.Scores, 
		
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
		"result": questionResult,
		"message": nil,
	})
}

	

func GetQuestionResultByQuestionId(c *gin.Context) {
	schema := c.Params.ByName("schema")
	questionId := c.Params.ByName("questionId")
	var questionResultList []structs.QuestionResult
	var questionResult structs.QuestionResult

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+schema+`"."QuestionResult" WHERE "QuestionId"=$1`, questionId )
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
		&questionResult.Id, 
		&questionResult.Date, 
		&questionResult.QuestionId, 
		&questionResult.ClientId, 
		&questionResult.Scores, 
		)
		questionResultList = append(questionResultList, questionResult)
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
		"result": questionResultList,
		"message": nil,
	})
}

	
func GetQuestionResultByClientId(c *gin.Context) {
	schema := c.Params.ByName("schema")
	clientId := c.Params.ByName("clientId")
	var questionResultList []structs.QuestionResult
	var questionResult structs.QuestionResult

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+schema+`"."QuestionResult" WHERE "ClientId"=$1`, clientId )
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
		&questionResult.Id, 
		&questionResult.Date, 
		&questionResult.QuestionId, 
		&questionResult.ClientId, 
		&questionResult.Scores, 
		)
		questionResultList = append(questionResultList, questionResult)
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
		"result": questionResultList,
		"message": nil,
	})
}

	

func CreateQuestionResult(c *gin.Context) {
	schema := c.Params.ByName("schema")
	data := DataProcessingQuestionResult(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+schema+`"."QuestionResult"
		(
		"Date", 
		"QuestionId", 
		"ClientId", 
		"Scores", 
		
		) 
		VALUES( $1, $2, $3, $4 )`,
		data.Date, 
		data.QuestionId, 
		data.ClientId, 
		data.Scores, 
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
	

func UpdateQuestionResult(c *gin.Context) {

	schema := c.Params.ByName("schema")
	id := c.Params.ByName("id")
	data := DataProcessingQuestionResult(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+schema+`"."QuestionResult" 
		SET 
		"Date"=$1,
		"QuestionId"=$2,
		"ClientId"=$3,
		"Scores"=$4
		
		WHERE "Id"=$1`,
		id,
		data.Date, 
		data.QuestionId, 
		data.ClientId, 
		data.Scores, 
		
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
	

func DeleteQuestionResult(c *gin.Context) {
	schema := c.Params.ByName("schema")
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+schema+`"."QuestionResult" WHERE "Id"=$1`, id)
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
	
	
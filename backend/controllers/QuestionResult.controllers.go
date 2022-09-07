
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
	model := c.Value("Model").(structs.Claims)
	var questionResultList []structs.QuestionResult
	var questionResult structs.QuestionResult

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.KeySchema+`"."QuestionResult"`)
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
		&questionResult.Time, 
		&questionResult.Scores, 
		&questionResult.QuestionId, 
		&questionResult.ClientId, 
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
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	var questionResult structs.QuestionResult

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.KeySchema+`"."QuestionResult" WHERE "Id"=$1`, id ).Scan(
		&questionResult.Id, 
		&questionResult.Date, 
		&questionResult.Time, 
		&questionResult.Scores, 
		&questionResult.QuestionId, 
		&questionResult.ClientId, 
		
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
	model := c.Value("Model").(structs.Claims)
	questionId := c.Params.ByName("questionId")
	var questionResult structs.QuestionResult

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.KeySchema+`"."QuestionResult" WHERE "QuestionId"=$1`, questionId ).Scan(
		&questionResult.Id, 
		&questionResult.Date, 
		&questionResult.Time, 
		&questionResult.Scores, 
		&questionResult.QuestionId, 
		&questionResult.ClientId, 
		
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

	


func CreateQuestionResult(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	data := DataProcessingQuestionResult(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+model.KeySchema+`"."QuestionResult"
		(
		"Date", 
		"Time", 
		"Scores", 
		"QuestionId", 
		"ClientId", 
		
		) 
		VALUES( $1, $2, $3, $4, $5 )`,
		data.Date, 
		data.Time, 
		data.Scores, 
		data.QuestionId, 
		data.ClientId, 
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

	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	data := DataProcessingQuestionResult(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+model.KeySchema+`"."QuestionResult" 
		SET 
		"Date"=$1,
		"Time"=$2,
		"Scores"=$3,
		"QuestionId"=$4,
		"ClientId"=$5
		
		WHERE "Id"=$1`,
		id,
		data.Date, 
		data.Time, 
		data.Scores, 
		data.QuestionId, 
		data.ClientId, 
		
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
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+model.KeySchema+`"."QuestionResult" WHERE "Id"=$1`, id)
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
	
	
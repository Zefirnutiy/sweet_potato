
package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingQuestion(c gin.Context) structs.Question {
	var data structs.Question
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.Question{}
	}
	return data
}


func GetQuestions(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	var questionList []structs.Question
	var question structs.Question

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.Schema+`"."Question"`)
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
		&question.Id, 
		&question.Text, 
		&question.Date, 
		&question.Time, 
		&question.DateDel, 
		&question.TimeDel, 
		&question.Hint, 
		&question.AnswerVariant, 
		&question.AnswerCorrect, 
		&question.Files, 
		&question.TestId, 
		&question.QuestionTypeId, 
		)
		questionList = append(questionList, question)
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
		"result": questionList,
		"message": nil,
	})
}

func GetQuestionById(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	var question structs.Question

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.Schema+`"."Question" WHERE "Id"=$1`, id ).Scan(
		&question.Id, 
		&question.Text, 
		&question.Date, 
		&question.Time, 
		&question.DateDel, 
		&question.TimeDel, 
		&question.Hint, 
		&question.AnswerVariant, 
		&question.AnswerCorrect, 
		&question.Files, 
		&question.TestId, 
		&question.QuestionTypeId, 
		
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
		"result": question,
		"message": nil,
	})
}

	

func GetQuestionByTestId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	testId := c.Params.ByName("testId")
	var questionList []structs.Question
	var question structs.Question

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.Schema+`"."Question" WHERE "TestId"=$1`, testId )
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
		&question.Id, 
		&question.Text, 
		&question.Date, 
		&question.Time, 
		&question.DateDel, 
		&question.TimeDel, 
		&question.Hint, 
		&question.AnswerVariant, 
		&question.AnswerCorrect, 
		&question.Files, 
		&question.TestId, 
		&question.QuestionTypeId, 
		)
		questionList = append(questionList, question)
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
		"result": questionList,
		"message": nil,
	})
}

	

func CreateQuestion(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	data := DataProcessingQuestion(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+model.Schema+`"."Question"
		(
		"Text", 
		"Date", 
		"Time", 
		"Hint", 
		"AnswerVariant", 
		"AnswerCorrect", 
		"Files", 
		"TestId", 
		"QuestionTypeId", 
		
		) 
		VALUES( $1, $2, $3, $4, $5, $6, $7, $8, $9 )`,
		data.Text, 
		data.Date, 
		data.Time, 
		data.Hint, 
		data.AnswerVariant, 
		data.AnswerCorrect, 
		data.Files, 
		data.TestId, 
		data.QuestionTypeId, 
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
	

func UpdateQuestion(c *gin.Context) {

	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	data := DataProcessingQuestion(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+model.Schema+`"."Question" 
		SET 
		"Text"=$1,
		"Date"=$2,
		"Time"=$3,
		"DateDel"=$4,
		"TimeDel"=$5,
		"Hint"=$6,
		"AnswerVariant"=$7,
		"AnswerCorrect"=$8,
		"Files"=$9,
		"TestId"=$10,
		"QuestionTypeId"=$11
		
		WHERE "Id"=$1`,
		id,
		data.Text, 
		data.Date, 
		data.Time, 
		data.DateDel, 
		data.TimeDel, 
		data.Hint, 
		data.AnswerVariant, 
		data.AnswerCorrect, 
		data.Files, 
		data.TestId, 
		data.QuestionTypeId, 
		
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
	

func DeleteQuestion(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+model.Schema+`"."Question" WHERE "Id"=$1`, id)
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
	
	
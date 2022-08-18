
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
		&question.DateDel, 
		&question.TestId, 
		&question.QuestionTypeId, 
		&question.Hint, 
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
		&question.DateDel, 
		&question.TestId, 
		&question.QuestionTypeId, 
		&question.Hint, 
		
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
		&question.DateDel, 
		&question.TestId, 
		&question.QuestionTypeId, 
		&question.Hint, 
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

	
func GetQuestionByQuestionTypeId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	questionTypeId := c.Params.ByName("questionTypeId")
	var questionList []structs.Question
	var question structs.Question

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.Schema+`"."Question" WHERE "QuestionTypeId"=$1`, questionTypeId )
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
		&question.DateDel, 
		&question.TestId, 
		&question.QuestionTypeId, 
		&question.Hint, 
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
		"DateDel", 
		"TestId", 
		"QuestionTypeId", 
		"Hint", 
		
		) 
		VALUES( $1, $2, $3, $4, $5, $6 )`,
		data.Text, 
		data.Date, 
		data.DateDel, 
		data.TestId, 
		data.QuestionTypeId, 
		data.Hint, 
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
		"DateDel"=$3,
		"TestId"=$4,
		"QuestionTypeId"=$5,
		"Hint"=$6
		
		WHERE "Id"=$1`,
		id,
		data.Text, 
		data.Date, 
		data.DateDel, 
		data.TestId, 
		data.QuestionTypeId, 
		data.Hint, 
		
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
	
	
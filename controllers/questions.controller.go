package controllers

import (
	"fmt"
	// "github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/functions"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/gin-gonic/gin"
)

func GetQuestions(c *gin.Context) {
	rows, e := db.Dbpool.Query(`SELECT * FROM "Question"`)

	functions.HandlerError(e)

	var question_list []structs.Question
	var question structs.Question

	for rows.Next() {
		e = rows.Scan(&question.Id, &question.Test_id, &question.Text, &question.Answer, &question.Files, &question.Scores)
		fmt.Println(question)
		question_list = append(question_list, question)
		functions.HandlerError(e)
	}

	c.JSON(200, gin.H{
		"Вопросы": question_list,
	})
}

func GetQuestionsById(c *gin.Context) {
	id := c.Param("id")

	rows, e := db.Dbpool.Query(`SELECT * FROM "Question" WHERE id=$1`, id)

	functions.HandlerError(e)

	var question structs.Question

	for rows.Next() {
		e = rows.Scan(&question.Id, &question.Test_id, &question.Text, &question.Answer, &question.Files, &question.Scores)
		functions.HandlerError(e)
	}

	c.JSON(200, gin.H{
		"Вопрос": question,
	})

}

func CreateQuestion(c *gin.Context) {
	idTest := c.Param("idTest")
	text := c.PostForm("text")
	answer := c.PostForm("answer")
	scores := c.PostForm("scores")

	_, err := db.Dbpool.Query(`INSERT INTO "Question"("test_id", "text", "answer", "scores") VALUES ($1, $2, $3, $4, )`, idTest, text, answer, scores)

	functions.HandlerError(err)

	c.JSON(200, gin.H{
		"message": "Вопрос успешно создан",
	})
}

func UpdateQuestion(c *gin.Context) {
	// id = c.Param("id")
}

func DeleteQuestion(c *gin.Context) {
	// id = c.Param("id")
}

func SearchQuestions(c *gin.Context) {

}


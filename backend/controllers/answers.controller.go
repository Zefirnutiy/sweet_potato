package controllers

import (
	"collage_project/backend/db"
	"collage_project/backend/functions"
	"collage_project/backend/structs"
	"github.com/gin-gonic/gin"
)

func GetAnswersByQuestionId(c *gin.Context) {
	idQuestion := c.Param("idQuestion")
	answers, e := db.Dbpool.Query(`SELECT * FROM "Answer_option" WHERE question_id = $1`, idQuestion)

	functions.HandlerError(e)

	var answerOptionList []structs.Answer_option
	var answerOption structs.Answer_option

	for answers.Next() {
		e = answers.Scan(&answerOption.Id, &answerOption.Question_id, &answerOption.Value, &answerOption.True_answer)
		answerOptionList = append(answerOptionList, answerOption)
		functions.HandlerError(e)
	}

	c.JSON(200, gin.H{
		"Ответы": answerOptionList,
	})
}

func CreateAnswer(c *gin.Context) {
	idQuestion := c.Param("idQuestion")
	value := c.PostForm("value")
	trueAnswer := c.PostForm("trueAnswer")

	_, err := db.Dbpool.Query(`INSERT INTO "Answer_option"("question_id", "value", "true_answer") VALUES($1, $2, $3)`, idQuestion, value, trueAnswer)

	functions.HandlerError(err)

	c.JSON(200, gin.H{
		"message": "Вариант ответа создан успешно.",
	})

}

func UpdateAnswer(c *gin.Context) {
	id := c.Param("id")
	value := c.PostForm("value")
	trueAnswer := c.PostForm("trueAnswer")

	_, err := db.Dbpool.Query(`UPDATE "Answer_option" SET "value"=$1, "true_answer"=$2 WHERE "id"=$3`, value, trueAnswer, id)

	functions.HandlerError(err)

	c.JSON(200, gin.H{
		"message": "Вариант ответа изменен успешно.",
	})
}

func DeleteAnswer(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Dbpool.Query(`DELETE FROM "Answer_option" WHERE "id"=$1`, id)

	functions.HandlerError(err)

	c.JSON(200, gin.H{
		"message": "Вариант ответа удален успешно.",
	})
}

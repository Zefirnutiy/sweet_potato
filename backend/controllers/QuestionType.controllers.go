
package controllers

import (
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
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.QuestionType{}
	}
	return data
}


func GetQuestionTypes(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	var questionTypeList []structs.QuestionType
	var questionType structs.QuestionType

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.Schema+`"."QuestionType"`)
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
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	var questionType structs.QuestionType

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.Schema+`"."QuestionType" WHERE "Id"=$1`, id ).Scan(
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
	model := c.Value("Model").(structs.Claims)
	data := DataProcessingQuestionType(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+model.Schema+`"."QuestionType"
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

	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	data := DataProcessingQuestionType(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+model.Schema+`"."QuestionType" 
		SET 
		"Type"=$1
		
		WHERE "Id"=$1`,
		id,
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
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+model.Schema+`"."QuestionType" WHERE "Id"=$1`, id)
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
	
	
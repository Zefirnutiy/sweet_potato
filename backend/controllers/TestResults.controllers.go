
package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingTestResults(c gin.Context) structs.TestResults {
	var data structs.TestResults
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.TestResults{}
	}
	return data
}


func GetTestResultss(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	var testResultsList []structs.TestResults
	var testResults structs.TestResults

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.Schema+`"."TestResults"`)
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
		&testResults.Id, 
		&testResults.Time, 
		&testResults.Date, 
		&testResults.ClientId, 
		&testResults.TestId, 
		&testResults.Assessment, 
		&testResults.PassageTime, 
		&testResults.Scores, 
		&testResults.CourseId, 
		)
		testResultsList = append(testResultsList, testResults)
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
		"result": testResultsList,
		"message": nil,
	})
}

func GetTestResultsById(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	var testResults structs.TestResults

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.Schema+`"."TestResults" WHERE "Id"=$1`, id ).Scan(
		&testResults.Id, 
		&testResults.Time, 
		&testResults.Date, 
		&testResults.ClientId, 
		&testResults.TestId, 
		&testResults.Assessment, 
		&testResults.PassageTime, 
		&testResults.Scores, 
		&testResults.CourseId, 
		
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
		"result": testResults,
		"message": nil,
	})
}

	

func GetTestResultsByClientId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	clientId := c.Params.ByName("clientId")
	var testResultsList []structs.TestResults
	var testResults structs.TestResults

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.Schema+`"."TestResults" WHERE "ClientId"=$1`, clientId )
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
		&testResults.Id, 
		&testResults.Time, 
		&testResults.Date, 
		&testResults.ClientId, 
		&testResults.TestId, 
		&testResults.Assessment, 
		&testResults.PassageTime, 
		&testResults.Scores, 
		&testResults.CourseId, 
		)
		testResultsList = append(testResultsList, testResults)
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
		"result": testResultsList,
		"message": nil,
	})
}

	
func GetTestResultsByCourseId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	courseId := c.Params.ByName("courseId")
	var testResultsList []structs.TestResults
	var testResults structs.TestResults

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.Schema+`"."TestResults" WHERE "CourseId"=$1`, courseId )
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
		&testResults.Id, 
		&testResults.Time, 
		&testResults.Date, 
		&testResults.ClientId, 
		&testResults.TestId, 
		&testResults.Assessment, 
		&testResults.PassageTime, 
		&testResults.Scores, 
		&testResults.CourseId, 
		)
		testResultsList = append(testResultsList, testResults)
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
		"result": testResultsList,
		"message": nil,
	})
}

	

func CreateTestResults(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	data := DataProcessingTestResults(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+model.Schema+`"."TestResults"
		(
		"Time", 
		"Date", 
		"ClientId", 
		"TestId", 
		"Assessment", 
		"PassageTime", 
		"Scores", 
		"CourseId", 
		
		) 
		VALUES( $1, $2, $3, $4, $5, $6, $7, $8 )`,
		data.Time, 
		data.Date, 
		data.ClientId, 
		data.TestId, 
		data.Assessment, 
		data.PassageTime, 
		data.Scores, 
		data.CourseId, 
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
	

func UpdateTestResults(c *gin.Context) {

	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	data := DataProcessingTestResults(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+model.Schema+`"."TestResults" 
		SET 
		"Time"=$1,
		"Date"=$2,
		"ClientId"=$3,
		"TestId"=$4,
		"Assessment"=$5,
		"PassageTime"=$6,
		"Scores"=$7,
		"CourseId"=$8
		
		WHERE "Id"=$1`,
		id,
		data.Time, 
		data.Date, 
		data.ClientId, 
		data.TestId, 
		data.Assessment, 
		data.PassageTime, 
		data.Scores, 
		data.CourseId, 
		
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
	

func DeleteTestResults(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+model.Schema+`"."TestResults" WHERE "Id"=$1`, id)
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
	
	
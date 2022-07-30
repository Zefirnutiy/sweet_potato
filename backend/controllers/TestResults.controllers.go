
package controllers

import (
	"fmt"
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
		c.JSON(500, gin.H{
			"message": "Некорректные данные",
		})
		return structs.TestResults{}
	}
	return data
}


func GetTestResultss(c *gin.Context) {
	schema := c.GetString("schema")
	var testResultsList []structs.TestResults
	var testResults structs.TestResults

	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	rows, err := db.Dbpool.Query(`SELECT * FROM "` + schema + `"."TestResults"`, schema)
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
	schema := c.GetString("schema")
	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	id := c.Params.ByName("id")
	var testResults structs.TestResults

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+schema+`"."TestResults" WHERE "Id"=$1`, id ).Scan(
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
	schema := c.GetString("schema")
	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	clientId := c.Params.ByName("clientId")
	var testResults structs.TestResults

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+schema+`"."TestResults" WHERE "ClientId"=$1`, clientId ).Scan(
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

	
func GetTestResultsByCourseId(c *gin.Context) {
	schema := c.GetString("schema")
	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	courseId := c.Params.ByName("courseId")
	var testResults structs.TestResults

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+schema+`"."TestResults" WHERE "CourseId"=$1`, courseId ).Scan(
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

	

func CreateTestResults(c *gin.Context) {
	schema := c.GetString("schema")
	data := DataProcessingTestResults(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+schema+`"."TestResults"
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

	schema := c.GetString("schema")
	data := DataProcessingTestResults(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+schema+`"."TestResults" 
		SET 
		"Time"=$2,
		"Date"=$3,
		"ClientId"=$4,
		"TestId"=$5,
		"Assessment"=$6,
		"PassageTime"=$7,
		"Scores"=$8,
		"CourseId"=$9
		
		WHERE "Id"=$1`,
		data.Id, 
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
	schema := c.GetString("schema")
	tokken := c.Params.ByName("tokken")
	id := c.Params.ByName("id") 
	fmt.Println(tokken)
	_, err := db.Dbpool.Exec(`DELETE FROM "`+schema+`"."TestResults" WHERE "Id"=$1`, id)
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
	
	
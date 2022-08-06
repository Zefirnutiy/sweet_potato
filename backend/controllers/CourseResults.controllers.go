
package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingCourseResults(c gin.Context) structs.CourseResults {
	var data structs.CourseResults
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.CourseResults{}
	}
	return data
}


func GetCourseResultss(c *gin.Context) {
	schema := c.Params.ByName("schema")
	var courseResultsList []structs.CourseResults
	var courseResults structs.CourseResults

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+schema+`"."CourseResults"`)
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
		&courseResults.Id, 
		&courseResults.Time, 
		&courseResults.Date, 
		&courseResults.Assessment, 
		&courseResults.Scores, 
		&courseResults.ClientId, 
		)
		courseResultsList = append(courseResultsList, courseResults)
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
		"result": courseResultsList,
		"message": nil,
	})
}

func GetCourseResultsById(c *gin.Context) {
	schema := c.Params.ByName("schema")
	id := c.Params.ByName("id")
	var courseResults structs.CourseResults

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+schema+`"."CourseResults" WHERE "Id"=$1`, id ).Scan(
		&courseResults.Id, 
		&courseResults.Time, 
		&courseResults.Date, 
		&courseResults.Assessment, 
		&courseResults.Scores, 
		&courseResults.ClientId, 
		
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
		"result": courseResults,
		"message": nil,
	})
}

	

func GetCourseResultsByClientId(c *gin.Context) {
	schema := c.Params.ByName("schema")
	clientId := c.Params.ByName("clientId")
	var courseResultsList []structs.CourseResults
	var courseResults structs.CourseResults

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+schema+`"."CourseResults" WHERE "ClientId"=$1`, clientId )
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
		&courseResults.Id, 
		&courseResults.Time, 
		&courseResults.Date, 
		&courseResults.Assessment, 
		&courseResults.Scores, 
		&courseResults.ClientId, 
		)
		courseResultsList = append(courseResultsList, courseResults)
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
		"result": courseResultsList,
		"message": nil,
	})
}

	

func CreateCourseResults(c *gin.Context) {
	schema := c.Params.ByName("schema")
	data := DataProcessingCourseResults(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+schema+`"."CourseResults"
		(
		"Time", 
		"Date", 
		"Assessment", 
		"Scores", 
		"ClientId", 
		
		) 
		VALUES( $1, $2, $3, $4, $5 )`,
		data.Time, 
		data.Date, 
		data.Assessment, 
		data.Scores, 
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
	

func UpdateCourseResults(c *gin.Context) {

	schema := c.Params.ByName("schema")
	id := c.Params.ByName("id")
	data := DataProcessingCourseResults(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+schema+`"."CourseResults" 
		SET 
		"Time"=$1,
		"Date"=$2,
		"Assessment"=$3,
		"Scores"=$4,
		"ClientId"=$5
		
		WHERE "Id"=$1`,
		id,
		data.Time, 
		data.Date, 
		data.Assessment, 
		data.Scores, 
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
	

func DeleteCourseResults(c *gin.Context) {
	schema := c.Params.ByName("schema")
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+schema+`"."CourseResults" WHERE "Id"=$1`, id)
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
	
	
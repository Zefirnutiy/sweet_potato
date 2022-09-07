
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
	model := c.Value("Model").(structs.Claims)
	var courseResultsList []structs.CourseResults
	var courseResults structs.CourseResults

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.KeySchema+`"."CourseResults"`)
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
		&courseResults.Assessment, 
		&courseResults.Scores, 
		&courseResults.NumberTests, 
		&courseResults.PassageTime, 
		&courseResults.Date, 
		&courseResults.Time, 
		&courseResults.ClientId, 
		&courseResults.CourseId, 
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
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	var courseResults structs.CourseResults

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.KeySchema+`"."CourseResults" WHERE "Id"=$1`, id ).Scan(
		&courseResults.Id, 
		&courseResults.Assessment, 
		&courseResults.Scores, 
		&courseResults.NumberTests, 
		&courseResults.PassageTime, 
		&courseResults.Date, 
		&courseResults.Time, 
		&courseResults.ClientId, 
		&courseResults.CourseId, 
		
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

	
func GetCourseResultsByCourseId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	courseId := c.Params.ByName("courseId")
	var courseResults structs.CourseResults

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.KeySchema+`"."CourseResults" WHERE "CourseId"=$1`, courseId ).Scan(
		&courseResults.Id, 
		&courseResults.Assessment, 
		&courseResults.Scores, 
		&courseResults.NumberTests, 
		&courseResults.PassageTime, 
		&courseResults.Date, 
		&courseResults.Time, 
		&courseResults.ClientId, 
		&courseResults.CourseId, 
		
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
	model := c.Value("Model").(structs.Claims)
	clientId := c.Params.ByName("clientId")
	var courseResultsList []structs.CourseResults
	var courseResults structs.CourseResults

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.KeySchema+`"."CourseResults" WHERE "ClientId"=$1`, clientId )
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
		&courseResults.Assessment, 
		&courseResults.Scores, 
		&courseResults.NumberTests, 
		&courseResults.PassageTime, 
		&courseResults.Date, 
		&courseResults.Time, 
		&courseResults.ClientId, 
		&courseResults.CourseId, 
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
	model := c.Value("Model").(structs.Claims)
	data := DataProcessingCourseResults(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+model.KeySchema+`"."CourseResults"
		(
		"Assessment", 
		"Scores", 
		"NumberTests", 
		"PassageTime", 
		"Date", 
		"Time", 
		"ClientId", 
		"CourseId", 
		
		) 
		VALUES( $1, $2, $3, $4, $5, $6, $7, $8 )`,
		data.Assessment, 
		data.Scores, 
		data.NumberTests, 
		data.PassageTime, 
		data.Date, 
		data.Time, 
		data.ClientId, 
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
	

func UpdateCourseResults(c *gin.Context) {

	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	data := DataProcessingCourseResults(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+model.KeySchema+`"."CourseResults" 
		SET 
		"Assessment"=$1,
		"Scores"=$2,
		"NumberTests"=$3,
		"PassageTime"=$4,
		"Date"=$5,
		"Time"=$6,
		"ClientId"=$7,
		"CourseId"=$8
		
		WHERE "Id"=$1`,
		id,
		data.Assessment, 
		data.Scores, 
		data.NumberTests, 
		data.PassageTime, 
		data.Date, 
		data.Time, 
		data.ClientId, 
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
	

func DeleteCourseResults(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+model.KeySchema+`"."CourseResults" WHERE "Id"=$1`, id)
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
	
	
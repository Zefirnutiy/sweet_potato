
package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingCourse(c gin.Context) structs.Course {
	var data structs.Course
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.Course{}
	}
	return data
}


func GetCourses(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	var courseList []structs.Course
	var course structs.Course

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.Schema+`"."Course"`)
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
		&course.Id, 
		&course.Title, 
		&course.Text, 
		&course.Files, 
		&course.Date, 
		&course.Time, 
		&course.DateDel, 
		&course.TimeDel, 
		&course.ClientId, 
		)
		courseList = append(courseList, course)
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
		"result": courseList,
		"message": nil,
	})
}

func GetCourseById(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	var course structs.Course

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.Schema+`"."Course" WHERE "Id"=$1`, id ).Scan(
		&course.Id, 
		&course.Title, 
		&course.Text, 
		&course.Files, 
		&course.Date, 
		&course.Time, 
		&course.DateDel, 
		&course.TimeDel, 
		&course.ClientId, 
		
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
		"result": course,
		"message": nil,
	})
}

	
func GetCourseByClientId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	clientId := c.Params.ByName("clientId")
	var course structs.Course

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.Schema+`"."Course" WHERE "ClientId"=$1`, clientId ).Scan(
		&course.Id, 
		&course.Title, 
		&course.Text, 
		&course.Files, 
		&course.Date, 
		&course.Time, 
		&course.DateDel, 
		&course.TimeDel, 
		&course.ClientId, 
		
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
		"result": course,
		"message": nil,
	})
}

	


func CreateCourse(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	data := DataProcessingCourse(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+model.Schema+`"."Course"
		(
		"Title", 
		"Text", 
		"Files", 
		"Date", 
		"Time", 
		"ClientId", 
		
		) 
		VALUES( $1, $2, $3, $4, $5, $6 )`,
		data.Title, 
		data.Text, 
		data.Files, 
		data.Date, 
		data.Time, 
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
	

func UpdateCourse(c *gin.Context) {

	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	data := DataProcessingCourse(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+model.Schema+`"."Course" 
		SET 
		"Title"=$1,
		"Text"=$2,
		"Files"=$3,
		"Date"=$4,
		"Time"=$5,
		"DateDel"=$6,
		"TimeDel"=$7,
		"ClientId"=$8
		
		WHERE "Id"=$1`,
		id,
		data.Title, 
		data.Text, 
		data.Files, 
		data.Date, 
		data.Time, 
		data.DateDel, 
		data.TimeDel, 
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
	

func DeleteCourse(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+model.Schema+`"."Course" WHERE "Id"=$1`, id)
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
	
	
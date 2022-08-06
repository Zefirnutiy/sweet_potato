
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
	schema := c.Params.ByName("schema")
	var courseList []structs.Course
	var course structs.Course

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+schema+`"."Course"`)
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
		&course.Text, 
		&course.Date, 
		&course.DateDel, 
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
	schema := c.Params.ByName("schema")
	id := c.Params.ByName("id")
	var course structs.Course

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+schema+`"."Course" WHERE "Id"=$1`, id ).Scan(
		&course.Id, 
		&course.Text, 
		&course.Date, 
		&course.DateDel, 
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
	schema := c.Params.ByName("schema")
	clientId := c.Params.ByName("clientId")
	var courseList []structs.Course
	var course structs.Course

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+schema+`"."Course" WHERE "ClientId"=$1`, clientId )
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
		&course.Text, 
		&course.Date, 
		&course.DateDel, 
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

	

func CreateCourse(c *gin.Context) {
	schema := c.Params.ByName("schema")
	data := DataProcessingCourse(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+schema+`"."Course"
		(
		"Text", 
		"Date", 
		"DateDel", 
		"ClientId", 
		
		) 
		VALUES( $1, $2, $3, $4 )`,
		data.Text, 
		data.Date, 
		data.DateDel, 
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

	schema := c.Params.ByName("schema")
	id := c.Params.ByName("id")
	data := DataProcessingCourse(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+schema+`"."Course" 
		SET 
		"Text"=$1,
		"Date"=$2,
		"DateDel"=$3,
		"ClientId"=$4
		
		WHERE "Id"=$1`,
		id,
		data.Text, 
		data.Date, 
		data.DateDel, 
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
	schema := c.Params.ByName("schema")
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+schema+`"."Course" WHERE "Id"=$1`, id)
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
	
	

package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingActiveCourse(c gin.Context) structs.ActiveCourse {
	var data structs.ActiveCourse
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.ActiveCourse{}
	}
	return data
}


func GetActiveCourses(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	var activeCourseList []structs.ActiveCourse
	var activeCourse structs.ActiveCourse

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.Schema+`"."ActiveCourse"`)
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
		&activeCourse.Id, 
		&activeCourse.Date, 
		&activeCourse.Time, 
		&activeCourse.DateClose, 
		&activeCourse.TimeClose, 
		&activeCourse.CourseId, 
		&activeCourse.ClientId, 
		)
		activeCourseList = append(activeCourseList, activeCourse)
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
		"result": activeCourseList,
		"message": nil,
	})
}

func GetActiveCourseById(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	var activeCourse structs.ActiveCourse

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.Schema+`"."ActiveCourse" WHERE "Id"=$1`, id ).Scan(
		&activeCourse.Id, 
		&activeCourse.Date, 
		&activeCourse.Time, 
		&activeCourse.DateClose, 
		&activeCourse.TimeClose, 
		&activeCourse.CourseId, 
		&activeCourse.ClientId, 
		
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
		"result": activeCourse,
		"message": nil,
	})
}

	
func GetActiveCourseByCourseId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	courseId := c.Params.ByName("courseId")
	var activeCourse structs.ActiveCourse

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.Schema+`"."ActiveCourse" WHERE "CourseId"=$1`, courseId ).Scan(
		&activeCourse.Id, 
		&activeCourse.Date, 
		&activeCourse.Time, 
		&activeCourse.DateClose, 
		&activeCourse.TimeClose, 
		&activeCourse.CourseId, 
		&activeCourse.ClientId, 
		
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
		"result": activeCourse,
		"message": nil,
	})
}

	

func GetActiveCourseByClientId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	clientId := c.Params.ByName("clientId")
	var activeCourseList []structs.ActiveCourse
	var activeCourse structs.ActiveCourse

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.Schema+`"."ActiveCourse" WHERE "ClientId"=$1`, clientId )
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
		&activeCourse.Id, 
		&activeCourse.Date, 
		&activeCourse.Time, 
		&activeCourse.DateClose, 
		&activeCourse.TimeClose, 
		&activeCourse.CourseId, 
		&activeCourse.ClientId, 
		)
		activeCourseList = append(activeCourseList, activeCourse)
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
		"result": activeCourseList,
		"message": nil,
	})
}

	

func CreateActiveCourse(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	data := DataProcessingActiveCourse(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+model.Schema+`"."ActiveCourse"
		(
		"Date", 
		"Time", 
		"ClientId", 
		
		) 
		VALUES( $1, $2, $3 )`,
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
	

func UpdateActiveCourse(c *gin.Context) {

	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	data := DataProcessingActiveCourse(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+model.Schema+`"."ActiveCourse" 
		SET 
		"Date"=$1,
		"Time"=$2,
		"DateClose"=$3,
		"TimeClose"=$4,
		"CourseId"=$5,
		"ClientId"=$6
		
		WHERE "Id"=$1`,
		id,
		data.Date, 
		data.Time, 
		data.DateClose, 
		data.TimeClose, 
		data.CourseId, 
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
	

func DeleteActiveCourse(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+model.Schema+`"."ActiveCourse" WHERE "Id"=$1`, id)
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
	
	
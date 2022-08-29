
package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingTest(c gin.Context) structs.Test {
	var data structs.Test
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.Test{}
	}
	return data
}


func GetTests(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	var testList []structs.Test
	var test structs.Test

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.Schema+`"."Test"`)
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
		&test.Id, 
		&test.Title, 
		&test.Text, 
		&test.Files, 
		&test.Date, 
		&test.Time, 
		&test.DateDel, 
		&test.TimeDel, 
		)
		testList = append(testList, test)
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
		"result": testList,
		"message": nil,
	})
}

func GetTestById(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	var test structs.Test

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.Schema+`"."Test" WHERE "Id"=$1`, id ).Scan(
		&test.Id, 
		&test.Title, 
		&test.Text, 
		&test.Files, 
		&test.Date, 
		&test.Time, 
		&test.DateDel, 
		&test.TimeDel, 
		
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
		"result": test,
		"message": nil,
	})
}

	


func CreateTest(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	data := DataProcessingTest(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+model.Schema+`"."Test"
		(
		"Title", 
		"Text", 
		"Files", 
		"Date", 
		"Time", 
		
		) 
		VALUES( $1, $2, $3, $4, $5 )`,
		data.Title, 
		data.Text, 
		data.Files, 
		data.Date, 
		data.Time, 
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
	

func UpdateTest(c *gin.Context) {

	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	data := DataProcessingTest(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+model.Schema+`"."Test" 
		SET 
		"Title"=$1,
		"Text"=$2,
		"Files"=$3,
		"Date"=$4,
		"Time"=$5,
		"DateDel"=$6,
		"TimeDel"=$7
		
		WHERE "Id"=$1`,
		id,
		data.Title, 
		data.Text, 
		data.Files, 
		data.Date, 
		data.Time, 
		data.DateDel, 
		data.TimeDel, 
		
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
	

func DeleteTest(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+model.Schema+`"."Test" WHERE "Id"=$1`, id)
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
	
	
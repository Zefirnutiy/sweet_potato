
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
		&test.Date, 
		&test.DateDel, 
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
		&test.Date, 
		&test.DateDel, 
		
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
		"Date", 
		"DateDel", 
		
		) 
		VALUES( $1, $2, $3, $4 )`,
		data.Title, 
		data.Text, 
		data.Date, 
		data.DateDel, 
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
		"Date"=$3,
		"DateDel"=$4
		
		WHERE "Id"=$1`,
		id,
		data.Title, 
		data.Text, 
		data.Date, 
		data.DateDel, 
		
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
	
	
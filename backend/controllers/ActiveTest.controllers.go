
package controllers

import (
	"fmt"
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingActiveTest(c gin.Context) structs.ActiveTest {
	var data structs.ActiveTest
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"message": "Некорректные данные",
		})
		return structs.ActiveTest{}
	}
	return data
}


func GetActiveTests(c *gin.Context) {
	schema := c.GetString("schema")
	var activeTestList []structs.ActiveTest
	var activeTest structs.ActiveTest

	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	rows, err := db.Dbpool.Query(`SELECT * FROM "` + schema + `"."ActiveTest"`, schema)
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
		&activeTest.Id, 
		&activeTest.Start, 
		&activeTest.End, 
		&activeTest.Time, 
		&activeTest.Attempts, 
		&activeTest.TestId, 
		&activeTest.ClientId, 
		&activeTest.TrainingTest, 
		)
		activeTestList = append(activeTestList, activeTest)
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
		"result": activeTestList,
		"message": nil,
	})
}

func GetActiveTestById(c *gin.Context) {
	schema := c.GetString("schema")
	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	id := c.Params.ByName("id")
	var activeTest structs.ActiveTest

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+schema+`"."ActiveTest" WHERE "Id"=$1`, id ).Scan(
		&activeTest.Id, 
		&activeTest.Start, 
		&activeTest.End, 
		&activeTest.Time, 
		&activeTest.Attempts, 
		&activeTest.TestId, 
		&activeTest.ClientId, 
		&activeTest.TrainingTest, 
		
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
		"result": activeTest,
		"message": nil,
	})
}

	
func GetActiveTestByClientId(c *gin.Context) {
	schema := c.GetString("schema")
	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	clientId := c.Params.ByName("clientId")
	var activeTest structs.ActiveTest

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+schema+`"."ActiveTest" WHERE "ClientId"=$1`, clientId ).Scan(
		&activeTest.Id, 
		&activeTest.Start, 
		&activeTest.End, 
		&activeTest.Time, 
		&activeTest.Attempts, 
		&activeTest.TestId, 
		&activeTest.ClientId, 
		&activeTest.TrainingTest, 
		
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
		"result": activeTest,
		"message": nil,
	})
}

	


func CreateActiveTest(c *gin.Context) {
	schema := c.GetString("schema")
	data := DataProcessingActiveTest(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+schema+`"."ActiveTest"
		(
		"Start", 
		"End", 
		"Time", 
		"Attempts", 
		"TestId", 
		"ClientId", 
		"TrainingTest", 
		
		) 
		VALUES( $1, $2, $3, $4, $5, $6, $7 )`,
		data.Start, 
		data.End, 
		data.Time, 
		data.Attempts, 
		data.TestId, 
		data.ClientId, 
		data.TrainingTest, 
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
	

func UpdateActiveTest(c *gin.Context) {

	schema := c.GetString("schema")
	data := DataProcessingActiveTest(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+schema+`"."ActiveTest" 
		SET 
		"Start"=$2,
		"End"=$3,
		"Time"=$4,
		"Attempts"=$5,
		"TestId"=$6,
		"ClientId"=$7,
		"TrainingTest"=$8
		
		WHERE "Id"=$1`,
		data.Id, 
		data.Start, 
		data.End, 
		data.Time, 
		data.Attempts, 
		data.TestId, 
		data.ClientId, 
		data.TrainingTest, 
		
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
	

func DeleteActiveTest(c *gin.Context) {
	schema := c.GetString("schema")
	tokken := c.Params.ByName("tokken")
	id := c.Params.ByName("id") 
	fmt.Println(tokken)
	_, err := db.Dbpool.Exec(`DELETE FROM "`+schema+`"."ActiveTest" WHERE "Id"=$1`, id)
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
	
	
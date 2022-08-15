
package controllers

import (
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
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.ActiveTest{}
	}
	return data
}

func GetActiveTests(c *gin.Context) {
	schema := c.Params.ByName("schema")
	var activeTestList []structs.ActiveTest
	var activeTest structs.ActiveTest

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+schema+`"."ActiveTest"`)
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
	schema := c.Params.ByName("schema")
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
	schema := c.Params.ByName("schema")
	clientId := c.Params.ByName("clientId")
	var activeTestList []structs.ActiveTest
	var activeTest structs.ActiveTest

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+schema+`"."ActiveTest" WHERE "ClientId"=$1`, clientId )
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


func CreateActiveTest(c *gin.Context) {
	schema := c.Params.ByName("schema")
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

	schema := c.Params.ByName("schema")
	id := c.Params.ByName("id")
	data := DataProcessingActiveTest(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+schema+`"."ActiveTest" 
		SET 
		"Start"=$1,
		"End"=$2,
		"Time"=$3,
		"Attempts"=$4,
		"TestId"=$5,
		"ClientId"=$6,
		"TrainingTest"=$7
		
		WHERE "Id"=$1`,
		id,
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
	schema := c.Params.ByName("schema")
	id := c.Params.ByName("id")
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
	
	
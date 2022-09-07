
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
	model := c.Value("Model").(structs.Claims)
	var activeTestList []structs.ActiveTest
	var activeTest structs.ActiveTest

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.KeySchema+`"."ActiveTest"`)
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
		&activeTest.Date, 
		&activeTest.Time, 
		&activeTest.DateClose, 
		&activeTest.TimeClose, 
		&activeTest.Attempts, 
		&activeTest.TestTypeId, 
		&activeTest.TestId, 
		&activeTest.ClientId, 
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
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	var activeTest structs.ActiveTest

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.KeySchema+`"."ActiveTest" WHERE "Id"=$1`, id ).Scan(
		&activeTest.Id, 
		&activeTest.Date, 
		&activeTest.Time, 
		&activeTest.DateClose, 
		&activeTest.TimeClose, 
		&activeTest.Attempts, 
		&activeTest.TestTypeId, 
		&activeTest.TestId, 
		&activeTest.ClientId, 
		
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

	
func GetActiveTestByTestId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	testId := c.Params.ByName("testId")
	var activeTest structs.ActiveTest

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.KeySchema+`"."ActiveTest" WHERE "TestId"=$1`, testId ).Scan(
		&activeTest.Id, 
		&activeTest.Date, 
		&activeTest.Time, 
		&activeTest.DateClose, 
		&activeTest.TimeClose, 
		&activeTest.Attempts, 
		&activeTest.TestTypeId, 
		&activeTest.TestId, 
		&activeTest.ClientId, 
		
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
	model := c.Value("Model").(structs.Claims)
	clientId := c.Params.ByName("clientId")
	var activeTest structs.ActiveTest

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.KeySchema+`"."ActiveTest" WHERE "ClientId"=$1`, clientId ).Scan(
		&activeTest.Id, 
		&activeTest.Date, 
		&activeTest.Time, 
		&activeTest.DateClose, 
		&activeTest.TimeClose, 
		&activeTest.Attempts, 
		&activeTest.TestTypeId, 
		&activeTest.TestId, 
		&activeTest.ClientId, 
		
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

	

func GetActiveTestByTestTypeId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	testTypeId := c.Params.ByName("testTypeId")
	var activeTestList []structs.ActiveTest
	var activeTest structs.ActiveTest

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.KeySchema+`"."ActiveTest" WHERE "TestTypeId"=$1`, testTypeId )
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
		&activeTest.Date, 
		&activeTest.Time, 
		&activeTest.DateClose, 
		&activeTest.TimeClose, 
		&activeTest.Attempts, 
		&activeTest.TestTypeId, 
		&activeTest.TestId, 
		&activeTest.ClientId, 
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
	model := c.Value("Model").(structs.Claims)
	data := DataProcessingActiveTest(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+model.KeySchema+`"."ActiveTest"
		(
		"Date", 
		"Time", 
		"Attempts", 
		"TestTypeId", 
		"TestId", 
		"ClientId", 
		
		) 
		VALUES( $1, $2, $3, $4, $5, $6 )`,
		data.Date, 
		data.Time, 
		data.Attempts, 
		data.TestTypeId, 
		data.TestId, 
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
	

func UpdateActiveTest(c *gin.Context) {

	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	data := DataProcessingActiveTest(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+model.KeySchema+`"."ActiveTest" 
		SET 
		"Date"=$1,
		"Time"=$2,
		"DateClose"=$3,
		"TimeClose"=$4,
		"Attempts"=$5,
		"TestTypeId"=$6,
		"TestId"=$7,
		"ClientId"=$8
		
		WHERE "Id"=$1`,
		id,
		data.Date, 
		data.Time, 
		data.DateClose, 
		data.TimeClose, 
		data.Attempts, 
		data.TestTypeId, 
		data.TestId, 
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
	

func DeleteActiveTest(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+model.KeySchema+`"."ActiveTest" WHERE "Id"=$1`, id)
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
	
	
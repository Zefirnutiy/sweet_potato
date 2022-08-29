
package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingInputType(c gin.Context) structs.InputType {
	var data structs.InputType
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.InputType{}
	}
	return data
}


func GetInputTypes(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	var inputTypeList []structs.InputType
	var inputType structs.InputType

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.Schema+`"."InputType"`)
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
		&inputType.Id, 
		&inputType.Title, 
		&inputType.Type, 
		)
		inputTypeList = append(inputTypeList, inputType)
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
		"result": inputTypeList,
		"message": nil,
	})
}

func GetInputTypeById(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	var inputType structs.InputType

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.Schema+`"."InputType" WHERE "Id"=$1`, id ).Scan(
		&inputType.Id, 
		&inputType.Title, 
		&inputType.Type, 
		
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
		"result": inputType,
		"message": nil,
	})
}

	


func CreateInputType(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	data := DataProcessingInputType(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+model.Schema+`"."InputType"
		(
		"Title", 
		"Type", 
		
		) 
		VALUES( $1, $2 )`,
		data.Title, 
		data.Type, 
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
	

func UpdateInputType(c *gin.Context) {

	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	data := DataProcessingInputType(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+model.Schema+`"."InputType" 
		SET 
		"Title"=$1,
		"Type"=$2
		
		WHERE "Id"=$1`,
		id,
		data.Title, 
		data.Type, 
		
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
	

func DeleteInputType(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+model.Schema+`"."InputType" WHERE "Id"=$1`, id)
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
	
	
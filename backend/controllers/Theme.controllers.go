
package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingTheme(c gin.Context) structs.Theme {
	var data structs.Theme
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.Theme{}
	}
	return data
}


func GetThemes(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	var themeList []structs.Theme
	var theme structs.Theme

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.KeySchema+`"."Theme"`)
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
		&theme.Id, 
		&theme.Title, 
		)
		themeList = append(themeList, theme)
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
		"result": themeList,
		"message": nil,
	})
}

func GetThemeById(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	var theme structs.Theme

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.KeySchema+`"."Theme" WHERE "Id"=$1`, id ).Scan(
		&theme.Id, 
		&theme.Title, 
		
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
		"result": theme,
		"message": nil,
	})
}

	


func CreateTheme(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	data := DataProcessingTheme(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+model.KeySchema+`"."Theme"
		(
		"Title", 
		
		) 
		VALUES( $1 )`,
		data.Title, 
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
	

func UpdateTheme(c *gin.Context) {

	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	data := DataProcessingTheme(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+model.KeySchema+`"."Theme" 
		SET 
		"Title"=$1
		
		WHERE "Id"=$1`,
		id,
		data.Title, 
		
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
	

func DeleteTheme(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+model.KeySchema+`"."Theme" WHERE "Id"=$1`, id)
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
	
	
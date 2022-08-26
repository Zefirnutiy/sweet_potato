
package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingPublicInfo(c gin.Context) structs.PublicInfo {
	var data structs.PublicInfo
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.PublicInfo{}
	}
	return data
}


func GetPublicInfos(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	var publicInfoList []structs.PublicInfo
	var publicInfo structs.PublicInfo

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.Schema+`"."PublicInfo"`)
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
		&publicInfo.Id, 
		&publicInfo.Title, 
		&publicInfo.Text, 
		&publicInfo.Date, 
		&publicInfo.DateDel, 
		&publicInfo.ClientId, 
		)
		publicInfoList = append(publicInfoList, publicInfo)
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
		"result": publicInfoList,
		"message": nil,
	})
}

func GetPublicInfoById(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	var publicInfo structs.PublicInfo

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.Schema+`"."PublicInfo" WHERE "Id"=$1`, id ).Scan(
		&publicInfo.Id, 
		&publicInfo.Title, 
		&publicInfo.Text, 
		&publicInfo.Date, 
		&publicInfo.DateDel, 
		&publicInfo.ClientId, 
		
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
		"result": publicInfo,
		"message": nil,
	})
}

	

func GetPublicInfoByClientId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	clientId := c.Params.ByName("clientId")
	var publicInfoList []structs.PublicInfo
	var publicInfo structs.PublicInfo

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.Schema+`"."PublicInfo" WHERE "ClientId"=$1`, clientId )
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
		&publicInfo.Id, 
		&publicInfo.Title, 
		&publicInfo.Text, 
		&publicInfo.Date, 
		&publicInfo.DateDel, 
		&publicInfo.ClientId, 
		)
		publicInfoList = append(publicInfoList, publicInfo)
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
		"result": publicInfoList,
		"message": nil,
	})
}

	

func CreatePublicInfo(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	data := DataProcessingPublicInfo(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+model.Schema+`"."PublicInfo"
		(
		"Title", 
		"Text", 
		"Date", 
		"DateDel", 
		"ClientId", 
		
		) 
		VALUES( $1, $2, $3, $4, $5 )`,
		data.Title, 
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
	

func UpdatePublicInfo(c *gin.Context) {

	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	data := DataProcessingPublicInfo(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+model.Schema+`"."PublicInfo" 
		SET 
		"Title"=$1,
		"Text"=$2,
		"Date"=$3,
		"DateDel"=$4,
		"ClientId"=$5
		
		WHERE "Id"=$1`,
		id,
		data.Title, 
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
	

func DeletePublicInfo(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+model.Schema+`"."PublicInfo" WHERE "Id"=$1`, id)
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
	
	
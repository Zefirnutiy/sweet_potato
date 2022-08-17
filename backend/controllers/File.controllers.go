package controllers

import (
	"fmt"
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingFile(c gin.Context) structs.File {
	var data structs.File
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.File{}
	}
	return data
}


func GetFiles(c *gin.Context) {
	schema := c.Params.ByName("schema")
	var fileList []structs.File
	var file structs.File

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+schema+`"."File"`)
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
		&file.Id, 
		&file.Date, 
		&file.DateDel, 
		&file.FileName, 
		&file.FileNameTmp, 
		&file.PublicInfoId, 
		&file.TestId, 
		&file.QuestionId, 
		&file.ClientId, 
		)
		fileList = append(fileList, file)
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
		"result": fileList,
		"message": nil,
	})
}

func GetFileById(c *gin.Context) {
	schema := c.Params.ByName("schema")
	id := c.Params.ByName("id")
	var file structs.File

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+schema+`"."File" WHERE "Id"=$1`, id ).Scan(
		&file.Id, 
		&file.Date, 
		&file.DateDel, 
		&file.FileName, 
		&file.FileNameTmp, 
		&file.PublicInfoId, 
		&file.TestId, 
		&file.QuestionId, 
		&file.ClientId, 
		
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
		"result": file,
		"message": nil,
	})
}

	

func GetFileByClientId(c *gin.Context) {
	schema := c.Params.ByName("schema")
	clientId := c.Params.ByName("clientId")
	var fileList []structs.File
	var file structs.File

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+schema+`"."File" WHERE "ClientId"=$1`, clientId )
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
		&file.Id, 
		&file.Date, 
		&file.DateDel, 
		&file.FileName, 
		&file.FileNameTmp, 
		&file.PublicInfoId, 
		&file.TestId, 
		&file.QuestionId, 
		&file.ClientId, 
		)
		fileList = append(fileList, file)
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
		"result": fileList,
		"message": nil,
	})
}


func UploadFile(c *gin.Context) {
	schema := c.Params.ByName("schema")
	data := DataProcessingFile(*c)
	var err error
	

	form, _ := c.MultipartForm()
    files := form.File["upload[]"]

    for _, file := range files {
      fmt.Println(file.Filename)
    }
	if err != nil {
		fmt.Println(err)
		return
	}


	_, err = db.Dbpool.Exec(`INSERT INTO "`+schema+`"."File"
		(
		"Date", 
		"DateDel", 
		"FileName", 
		"FileNameTmp", 
		"PublicInfoId", 
		"TestId", 
		"QuestionId", 
		"ClientId", 
		
		) 
		VALUES( $1, $2, $3, $4, $5, $6, $7, $8 )`,
		data.Date, 
		data.DateDel, 
		data.FileNameTmp, 
		data.PublicInfoId, 
		data.TestId, 
		data.QuestionId, 
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
	

func UpdateFile(c *gin.Context) {

	schema := c.Params.ByName("schema")
	id := c.Params.ByName("id")
	data := DataProcessingFile(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+schema+`"."File" 
		SET 
		"Date"=$1,
		"DateDel"=$2,
		"FileName"=$3,
		"FileNameTmp"=$4,
		"PublicInfoId"=$5,
		"TestId"=$6,
		"QuestionId"=$7,
		"ClientId"=$8
		
		WHERE "Id"=$1`,
		id,
		data.Date, 
		data.DateDel, 
		data.FileName, 
		data.FileNameTmp, 
		data.PublicInfoId, 
		data.TestId, 
		data.QuestionId, 
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
	

func DeleteFile(c *gin.Context) {
	schema := c.Params.ByName("schema")
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+schema+`"."File" WHERE "Id"=$1`, id)
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
	
	
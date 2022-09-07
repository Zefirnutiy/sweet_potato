
package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingFileInformation(c gin.Context) structs.FileInformation {
	var data structs.FileInformation
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.FileInformation{}
	}
	return data
}


func GetFileInformations(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	var fileInformationList []structs.FileInformation
	var fileInformation structs.FileInformation

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.KeySchema+`"."FileInformation"`)
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
		&fileInformation.ClientId, 
		&fileInformation.FileId, 
		&fileInformation.PublicInfoId, 
		&fileInformation.TestId, 
		&fileInformation.QuestionId, 
		)
		fileInformationList = append(fileInformationList, fileInformation)
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
		"result": fileInformationList,
		"message": nil,
	})
}

func GetFileInformationByFileId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	fileId := c.Params.ByName("fileId")
	var fileInformation structs.FileInformation

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.KeySchema+`"."FileInformation" WHERE "FileId"=$1`, fileId ).Scan(
		&fileInformation.ClientId, 
		&fileInformation.FileId, 
		&fileInformation.PublicInfoId, 
		&fileInformation.TestId, 
		&fileInformation.QuestionId, 
		
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
		"result": fileInformation,
		"message": nil,
	})
}

	
func GetFileInformationByTestId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	testId := c.Params.ByName("testId")
	var fileInformation structs.FileInformation

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.KeySchema+`"."FileInformation" WHERE "TestId"=$1`, testId ).Scan(
		&fileInformation.ClientId, 
		&fileInformation.FileId, 
		&fileInformation.PublicInfoId, 
		&fileInformation.TestId, 
		&fileInformation.QuestionId, 
		
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
		"result": fileInformation,
		"message": nil,
	})
}

	
func GetFileInformationByQuestionId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	questionId := c.Params.ByName("questionId")
	var fileInformation structs.FileInformation

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.KeySchema+`"."FileInformation" WHERE "QuestionId"=$1`, questionId ).Scan(
		&fileInformation.ClientId, 
		&fileInformation.FileId, 
		&fileInformation.PublicInfoId, 
		&fileInformation.TestId, 
		&fileInformation.QuestionId, 
		
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
		"result": fileInformation,
		"message": nil,
	})
}

	

func GetFileInformationByClientId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	clientId := c.Params.ByName("clientId")
	var fileInformationList []structs.FileInformation
	var fileInformation structs.FileInformation

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.KeySchema+`"."FileInformation" WHERE "ClientId"=$1`, clientId )
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
		&fileInformation.ClientId, 
		&fileInformation.FileId, 
		&fileInformation.PublicInfoId, 
		&fileInformation.TestId, 
		&fileInformation.QuestionId, 
		)
		fileInformationList = append(fileInformationList, fileInformation)
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
		"result": fileInformationList,
		"message": nil,
	})
}

	
func GetFileInformationByPublicInfoId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	publicInfoId := c.Params.ByName("publicInfoId")
	var fileInformationList []structs.FileInformation
	var fileInformation structs.FileInformation

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.KeySchema+`"."FileInformation" WHERE "PublicInfoId"=$1`, publicInfoId )
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
		&fileInformation.ClientId, 
		&fileInformation.FileId, 
		&fileInformation.PublicInfoId, 
		&fileInformation.TestId, 
		&fileInformation.QuestionId, 
		)
		fileInformationList = append(fileInformationList, fileInformation)
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
		"result": fileInformationList,
		"message": nil,
	})
}
	

func CreateFileInformation(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	data := DataProcessingFileInformation(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+model.KeySchema+`"."FileInformation"
		(
		"ClientId", 
		"FileId", 
		"PublicInfoId", 
		"TestId", 
		"QuestionId", 
		
		) 
		VALUES( $1, $2, $3, $4, $5 )`,
		data.ClientId, 
		data.FileId, 
		data.PublicInfoId, 
		data.TestId, 
		data.QuestionId, 
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
	

func UpdateFileInformation(c *gin.Context) {

	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	data := DataProcessingFileInformation(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+model.KeySchema+`"."FileInformation" 
		SET 
		"FileId"=$1,
		"PublicInfoId"=$2,
		"TestId"=$3,
		"QuestionId"=$4
		
		WHERE "Id"=$1`,
		id,
		data.FileId, 
		data.PublicInfoId, 
		data.TestId, 
		data.QuestionId, 
		
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
	

func DeleteFileInformation(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+model.KeySchema+`"."FileInformation" WHERE "Id"=$1`, id)
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
	
	

package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingSession(c gin.Context) structs.Session {
	var data structs.Session
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.Session{}
	}
	return data
}


func GetSessions(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	var sessionList []structs.Session
	var session structs.Session

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.KeySchema+`"."Session"`)
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
		&session.Id, 
		&session.UserId, 
		&session.IpAddress, 
		&session.BackendStartDateTime, 
		&session.StateChangeDateTime, 
		&session.StateId, 
		)
		sessionList = append(sessionList, session)
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
		"result": sessionList,
		"message": nil,
	})
}

func GetSessionById(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	var session structs.Session

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.KeySchema+`"."Session" WHERE "Id"=$1`, id ).Scan(
		&session.Id, 
		&session.UserId, 
		&session.IpAddress, 
		&session.BackendStartDateTime, 
		&session.StateChangeDateTime, 
		&session.StateId, 
		
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
		"result": session,
		"message": nil,
	})
}

	
func GetSessionByUserId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	userId := c.Params.ByName("userId")
	var session structs.Session

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.KeySchema+`"."Session" WHERE "UserId"=$1`, userId ).Scan(
		&session.Id, 
		&session.UserId, 
		&session.IpAddress, 
		&session.BackendStartDateTime, 
		&session.StateChangeDateTime, 
		&session.StateId, 
		
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
		"result": session,
		"message": nil,
	})
}

	


func CreateSession(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	data := DataProcessingSession(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+model.KeySchema+`"."Session"
		(
		"UserId", 
		"IpAddress", 
		"BackendStartDateTime", 
		"StateChangeDateTime", 
		"StateId", 
		
		) 
		VALUES( $1, $2, $3, $4, $5 )`,
		data.UserId, 
		data.IpAddress, 
		data.BackendStartDateTime, 
		data.StateChangeDateTime, 
		data.StateId, 
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
	

func UpdateSession(c *gin.Context) {

	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	data := DataProcessingSession(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+model.KeySchema+`"."Session" 
		SET 
		"UserId"=$1,
		"IpAddress"=$2,
		"BackendStartDateTime"=$3,
		"StateChangeDateTime"=$4,
		"StateId"=$5
		
		WHERE "Id"=$1`,
		id,
		data.UserId, 
		data.IpAddress, 
		data.BackendStartDateTime, 
		data.StateChangeDateTime, 
		data.StateId, 
		
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
	

func DeleteSession(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+model.KeySchema+`"."Session" WHERE "Id"=$1`, id)
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
	
	
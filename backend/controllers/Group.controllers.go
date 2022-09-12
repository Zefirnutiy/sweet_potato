package controllers

import (

	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingGroup(c gin.Context) structs.Group {
	var data structs.Group
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.Group{}
	}
	return data
}


func GetGroups(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	var groupList []structs.Group
	var group structs.Group

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.KeySchema+`"."Group"`)
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
		&group.Id, 
		&group.Title, 
		&group.TitleSingular, 
		&group.DepartmentId, 
		)
		groupList = append(groupList, group)
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
		"result": groupList,
		"message": nil,
	})
}

func GetGroupById(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	var group structs.Group

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.KeySchema+`"."Group" WHERE "Id"=$1`, id ).Scan(
		&group.Id, 
		&group.Title, 
		&group.TitleSingular, 
		&group.DepartmentId, 
		
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
		"result": group,
		"message": nil,
	})
}

	
func GetGroupByDepartmentId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	departmentId := c.Params.ByName("departmentId")
	var group structs.Group
	var groupList []structs.Group
  
	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.KeySchema+`"."Group" WHERE "DepartmentId"=$1`, departmentId )
  
	if err != nil {
	  utils.Logger.Println(err)
	  c.JSON(500, gin.H{
		"result": nil,
		"message": "Ничего не найдено",
	  })
	  return
	}
  
	for rows.Next(){
	  err = rows.Scan(
		&group.Id, 
		&group.Title, 
		&group.TitleSingular, 
		&group.DepartmentId, 
	  )
	  if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
		  "result": nil,
		  "message": "Ошибка сервера",
		})
		return
	  }
	  groupList = append(groupList, group)
	}

	c.JSON(200, gin.H{
		"result": group,
		"message": nil,
	})
  }

	


func CreateGroup(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	data := DataProcessingGroup(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+model.KeySchema+`"."Group"
		(
		"Title", 
		"TitleSingular", 
		"DepartmentId", 
		
		) 
		VALUES( $1, $2, $3 )`,
		data.Title, 
		data.TitleSingular, 
		data.DepartmentId, 
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
	

func UpdateGroup(c *gin.Context) {

	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	data := DataProcessingGroup(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+model.KeySchema+`"."Group" 
		SET 
		"Title"=$1,
		"TitleSingular"=$2,
		"DepartmentId"=$3
		
		WHERE "Id"=$1`,
		id,
		data.Title, 
		data.TitleSingular, 
		data.DepartmentId, 
		
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
	

func DeleteGroup(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+model.KeySchema+`"."Group" WHERE "Id"=$1`, id)
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
	
	
package controllers

import (
	"fmt"

	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingDepartment(c gin.Context) structs.Department {
	var data structs.Department
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.Department{}
	}
	return data
}


func GetDepartments(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	var department structs.Department
	type customeDepartment struct {
		Id 		int 
    	Title  	string  
		GroupCount int	
	}
	
	var departmentList []customeDepartment

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.KeySchema+`"."Department"`)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"result": nil,
			"message": "Ничего не найдено",
		})
		return
	}
	var groupCount int

	for rows.Next() {
		err = rows.Scan(
		&department.Id, 
		&department.Title, 
		)
		
		if err != nil {
			utils.Logger.Println(err)
			c.JSON(500, gin.H{
				"result": nil,
				"message": "Ошибка сервера",
			})
			return
		}
		query := fmt.Sprintf(`SELECT Count("Title") FROM "`+model.KeySchema+`"."Group" WHERE "DepartmentId"=$1`)
		row := db.Dbpool.QueryRow(query, department.Id)
		row.Scan(&groupCount)

		departmentList = append(departmentList, customeDepartment{department.Id, department.Title, groupCount})
	}

	fmt.Println(departmentList)
	c.JSON(200, gin.H{
		"result": departmentList,
		"message": nil,
	})
}

func GetDepartmentById(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	var department structs.Department

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.KeySchema+`"."Department" WHERE "Id"=$1`, id ).Scan(
		&department.Id, 
		&department.Title, 
		
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
		"result": department,
		"message": nil,
	})
}

	


func CreateDepartment(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	data := DataProcessingDepartment(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+model.KeySchema+`"."Department"
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
	

func UpdateDepartment(c *gin.Context) {

	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	data := DataProcessingDepartment(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+model.KeySchema+`"."Department" 
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
	

func DeleteDepartment(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+model.KeySchema+`"."Department" WHERE "Id"=$1`, id)
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
	
	
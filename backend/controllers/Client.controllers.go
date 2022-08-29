package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)
func DataProcessingClient(c gin.Context) structs.Client {
	var data structs.Client
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.Client{}
	}
	return data
}


func GetClients(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	var clientList []structs.Client
	var client structs.Client

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.Schema+`"."Client"`)
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
		&client.Id, 
		&client.FirstName, 
		&client.LastName, 
		&client.Patronymic, 
		&client.Password, 
		&client.Email, 
		&client.Telephone, 
		&client.EmailNotifications, 
		&client.ManageCourses, 
		&client.ManageUsers, 
		&client.AploadFiles, 
		&client.ViewYourResults, 
		&client.ViewOtherResults, 
		&client.DepartmentId, 
		&client.GroupId, 
		&client.CreatorId, 
		)
		clientList = append(clientList, client)
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
		"result": clientList,
		"message": nil,
	})
}

func GetClientById(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	var client structs.Client

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.Schema+`"."Client" WHERE "Id"=$1`, id ).Scan(
		&client.Id, 
		&client.FirstName, 
		&client.LastName, 
		&client.Patronymic, 
		&client.Password, 
		&client.Email, 
		&client.Telephone, 
		&client.EmailNotifications, 
		&client.ManageCourses, 
		&client.ManageUsers, 
		&client.AploadFiles, 
		&client.ViewYourResults, 
		&client.ViewOtherResults, 
		&client.DepartmentId, 
		&client.GroupId, 
		&client.CreatorId, 
		
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
		"result": client,
		"message": nil,
	})
}

	

func GetClientByDepartmentId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	departmentId := c.Params.ByName("departmentId")
	var clientList []structs.Client
	var client structs.Client

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.Schema+`"."Client" WHERE "DepartmentId"=$1`, departmentId )
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
		&client.Id, 
		&client.FirstName, 
		&client.LastName, 
		&client.Patronymic, 
		&client.Password, 
		&client.Email, 
		&client.Telephone, 
		&client.EmailNotifications, 
		&client.ManageCourses, 
		&client.ManageUsers, 
		&client.AploadFiles, 
		&client.ViewYourResults, 
		&client.ViewOtherResults, 
		&client.DepartmentId, 
		&client.GroupId, 
		&client.CreatorId, 
		)
		clientList = append(clientList, client)
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
		"result": clientList,
		"message": nil,
	})
}

	
func GetClientByGroupId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	groupId := c.Params.ByName("groupId")
	var clientList []structs.Client
	var client structs.Client

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.Schema+`"."Client" WHERE "GroupId"=$1`, groupId )
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
		&client.Id, 
		&client.FirstName, 
		&client.LastName, 
		&client.Patronymic, 
		&client.Password, 
		&client.Email, 
		&client.Telephone, 
		&client.EmailNotifications, 
		&client.ManageCourses, 
		&client.ManageUsers, 
		&client.AploadFiles, 
		&client.ViewYourResults, 
		&client.ViewOtherResults, 
		&client.DepartmentId, 
		&client.GroupId, 
		&client.CreatorId, 
		)
		clientList = append(clientList, client)
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
		"result": clientList,
		"message": nil,
	})
}

	
func GetClientByCreatorId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	creatorId := c.Params.ByName("creatorId")
	var clientList []structs.Client
	var client structs.Client

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.Schema+`"."Client" WHERE "CreatorId"=$1`, creatorId )
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
		&client.Id, 
		&client.FirstName, 
		&client.LastName, 
		&client.Patronymic, 
		&client.Password, 
		&client.Email, 
		&client.Telephone, 
		&client.EmailNotifications, 
		&client.ManageCourses, 
		&client.ManageUsers, 
		&client.AploadFiles, 
		&client.ViewYourResults, 
		&client.ViewOtherResults, 
		&client.DepartmentId, 
		&client.GroupId, 
		&client.CreatorId, 
		)
		clientList = append(clientList, client)
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
		"result": clientList,
		"message": nil,
	})
}


	
func CreateClient(c *gin.Context) {
	data := DataProcessingClient(*c)
	model := c.Value("Model").(structs.Claims)
	var err error

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"message": "Ошибка сервера",
		})
		return
	}

	_, err = db.Dbpool.Exec(`INSERT INTO "`+model.Schema+`"."Client"
		(
		"FirstName", 
		"LastName", 
		"Patronymic", 
		"Password", 
		"Email", 
		"Telephone", 
		"EmailNotifications", 
		"ManageCourses", 
		"ManageUsers", 
		"AploadFiles", 
		"ViewYourResults", 
		"ViewOtherResults", 
		"DepartmentId", 
		"GroupId", 
		"CreatorId", 
		
		) 
		VALUES( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15 )`,
		data.FirstName, 
		data.LastName, 
		data.Patronymic, 
		hashedPassword, 
		data.Email, 
		data.Telephone, 
		data.EmailNotifications, 
		data.ManageCourses, 
		data.ManageUsers, 
		data.AploadFiles, 
		data.ViewYourResults, 
		data.ViewOtherResults, 
		data.DepartmentId, 
		data.GroupId, 
		data.CreatorId, 
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

type ClientLogin struct {
	Schema   string
	Email    string
	Password string
}

func LoginClient(c *gin.Context) {
	var loginData ClientLogin
	err := c.ShouldBindJSON(&loginData)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "некорректные данные",
		})
		return
	}

	var client structs.Client
	err = db.Dbpool.QueryRow(`SELECT * FROM "`+loginData.Schema+`"."Client" WHERE "Email"=$1`, loginData.Email).Scan(
		&client.Id,
		&client.FirstName,
		&client.LastName,
		&client.Patronymic,
		&client.Password,
		&client.Email,
		&client.Telephone,
		&client.EmailNotifications,
		&client.ManageCourses,
		&client.ManageUsers,
		&client.AploadFiles,
		&client.ViewYourResults,
		&client.ViewOtherResults,
		&client.DepartmentId,
		&client.CreatorId,
		&client.GroupId)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "нет такого клиента",
		})
		return
	}

	inputPas := []byte(loginData.Password)
	clientPas := []byte(client.Password)
	err = bcrypt.CompareHashAndPassword(inputPas, clientPas)

	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"error":   err,
			"message": "Вы ввели неправильные данные",
		})
		return
	}

	claims := structs.Claims{
		Id:      client.Id,
		Email:   client.Email,
		Schema:  loginData.Schema}

	token, err := utils.CreateToken(claims)

	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"message": "ошибка создания токена",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "авторизирован",
		"token":   token,
	})

}


func UpdateClient(c *gin.Context) {

	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	data := DataProcessingClient(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+model.Schema+`"."Client" 
		SET 
		"FirstName"=$1,
		"LastName"=$2,
		"Patronymic"=$3,
		"Password"=$4,
		"Email"=$5,
		"Telephone"=$6,
		"EmailNotifications"=$7,
		"ManageCourses"=$8,
		"ManageUsers"=$9,
		"AploadFiles"=$10,
		"ViewYourResults"=$11,
		"ViewOtherResults"=$12,
		"DepartmentId"=$13,
		"GroupId"=$14,
		"CreatorId"=$15
		
		WHERE "Id"=$1`,
		id,
		data.FirstName, 
		data.LastName, 
		data.Patronymic, 
		data.Password, 
		data.Email, 
		data.Telephone, 
		data.EmailNotifications, 
		data.ManageCourses, 
		data.ManageUsers, 
		data.AploadFiles, 
		data.ViewYourResults, 
		data.ViewOtherResults, 
		data.DepartmentId, 
		data.GroupId, 
		data.CreatorId, 
		
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
	

func DeleteClient(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+model.Schema+`"."Client" WHERE "Id"=$1`, id)
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
	
	
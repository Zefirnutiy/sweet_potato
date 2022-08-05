
package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessingPayment(c gin.Context) structs.Payment {
	var data structs.Payment
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.Payment{}
	}
	return data
}


func GetPayments(c *gin.Context) {
	var paymentList []structs.Payment
	var payment structs.Payment

	rows, err := db.Dbpool.Query(`SELECT * FROM "Payment"`)
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
		&payment.Number, 
		&payment.Name, 
		&payment.Money, 
		&payment.Date, 
		&payment.LevelId, 
		&payment.OrganizationId, 
		)
		paymentList = append(paymentList, payment)
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
		"result": paymentList,
		"message": nil,
	})
}

func GetPaymentByNumber(c *gin.Context) {
	number := c.Params.ByName("number")
	var payment structs.Payment

	err := db.Dbpool.QueryRow(`SELECT * FROM "Payment" WHERE "Number"=$1`, number ).Scan(
		&payment.Number, 
		&payment.Name, 
		&payment.Money, 
		&payment.Date, 
		&payment.LevelId, 
		&payment.OrganizationId, 
		
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
		"result": payment,
		"message": nil,
	})
}

	
func GetPaymentByOrganizationId(c *gin.Context) {
	organizationId := c.Params.ByName("organizationId")
	var payment structs.Payment

	err := db.Dbpool.QueryRow(`SELECT * FROM "Payment" WHERE "OrganizationId"=$1`, organizationId ).Scan(
		&payment.Number, 
		&payment.Name, 
		&payment.Money, 
		&payment.Date, 
		&payment.LevelId, 
		&payment.OrganizationId, 
		
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
		"result": payment,
		"message": nil,
	})
}

	


func CreatePayment(c *gin.Context) {
	data := DataProcessingPayment(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "Payment"
		(
		"Number", 
		"Name", 
		"Money", 
		"Date", 
		"LevelId", 
		"OrganizationId", 
		
		) 
		VALUES( $1, $2, $3, $4, $5, $6 )`,
		data.Number, 
		data.Name, 
		data.Money, 
		data.Date, 
		data.LevelId, 
		data.OrganizationId, 
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
	

func UpdatePayment(c *gin.Context) {

	id := c.Params.ByName("id")
	data := DataProcessingPayment(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "Payment" 
		SET 
		"Name"=$1,
		"Money"=$2,
		"Date"=$3,
		"LevelId"=$4,
		"OrganizationId"=$5
		
		WHERE "Id"=$1`,
		id,
		data.Name, 
		data.Money, 
		data.Date, 
		data.LevelId, 
		data.OrganizationId, 
		
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
	

func DeletePayment(c *gin.Context) {
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "Payment" WHERE "Id"=$1`, id)
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
	
	
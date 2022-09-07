
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
	model := c.Value("Model").(structs.Claims)
	var PaymentList []structs.Payment
	var Payment structs.Payment

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.KeySchema+`"."Payment"`)
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
		&Payment.Number, 
		&Payment.Money, 
		&Payment.Date, 
		&Payment.Time, 
		&Payment.Users, 
		&Payment.Statistics, 
		&Payment.ProtectionCheating, 
		&Payment.OrganizationId, 
		)
		PaymentList = append(PaymentList, Payment)
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
		"result": PaymentList,
		"message": nil,
	})
}

func GetPaymentByNumber(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	number := c.Params.ByName("number")
	var Payment structs.Payment

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+model.KeySchema+`"."Payment" WHERE "Number"=$1`, number ).Scan(
		&Payment.Number, 
		&Payment.Money, 
		&Payment.Date, 
		&Payment.Time, 
		&Payment.Users, 
		&Payment.Statistics, 
		&Payment.ProtectionCheating, 
		&Payment.OrganizationId, 
		
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
		"result": Payment,
		"message": nil,
	})
}

	

func GetPaymentByOrganizationId(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	organizationId := c.Params.ByName("organizationId")
	var PaymentList []structs.Payment
	var Payment structs.Payment

	rows, err := db.Dbpool.Query(`SELECT * FROM "`+model.KeySchema+`"."Payment" WHERE "OrganizationId"=$1`, organizationId )
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
		&Payment.Number, 
		&Payment.Money, 
		&Payment.Date, 
		&Payment.Time, 
		&Payment.Users, 
		&Payment.Statistics, 
		&Payment.ProtectionCheating, 
		&Payment.OrganizationId, 
		)
		PaymentList = append(PaymentList, Payment)
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
		"result": PaymentList,
		"message": nil,
	})
}

	

func CreatePayment(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	data := DataProcessingPayment(*c)
	var err error
	

	_, err = db.Dbpool.Exec(`INSERT INTO "`+model.KeySchema+`"."Payment"
		(
		"Money", 
		"Date", 
		"Time", 
		"Users", 
		"Statistics", 
		"ProtectionCheating", 
		"OrganizationId", 
		
		) 
		VALUES( $1, $2, $3, $4, $5, $6, $7 )`,
		data.Money, 
		data.Date, 
		data.Time, 
		data.Users, 
		data.Statistics, 
		data.ProtectionCheating, 
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

	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	data := DataProcessingPayment(*c)
	var err error
	
	
	_, err = db.Dbpool.Exec(`UPDATE "`+model.KeySchema+`"."Payment" 
		SET 
		"Money"=$1,
		"Date"=$2,
		"Time"=$3,
		"Users"=$4,
		"Statistics"=$5,
		"ProtectionCheating"=$6,
		"OrganizationId"=$7
		
		WHERE "Id"=$1`,
		id,
		data.Money, 
		data.Date, 
		data.Time, 
		data.Users, 
		data.Statistics, 
		data.ProtectionCheating, 
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
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(`DELETE FROM "`+model.KeySchema+`"."Payment" WHERE "Id"=$1`, id)
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
	
	
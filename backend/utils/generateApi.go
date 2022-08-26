package utils

import (
	"fmt"
	"os"
	"strings"
)

var routerText string

// принимает строку, возвращает эту же строку но с первой буквой в нижнем регистре
func firstLower(str string) string{
	return strings.ToLower(string(str[0])) + str[1:]
}

// функция очищает слово от спец-символов скрипта
func clean(str string) string{

	str = strings.ReplaceAll(str, "*", "")
	str = strings.ReplaceAll(str, "#", "")
	str = strings.ReplaceAll(str, "$", "")
	str = strings.ReplaceAll(str, "%", "")

	return str
}

// возвращает чистый массив полей, в которых содержался спец-символ separator
func sortColumn(columns []string, separator string)[]string{
	fieldsMas := ""

	for _, column := range columns{
		if strings.Contains(column, separator) {
			fieldsMas += clean(column) + " "
		}
	}
	return strings.Fields(fieldsMas)
}

// принимает массив переменных, возвращает строку этих переменных в обертке
// из заданных символов
func simbAroundVar(mas []string, simbol string) string{
	text := ""

	for _, field := range mas{
		text += fmt.Sprintf(`%[2]s%[1]s%[2]s, 
		`, field, simbol)
	}

	return text
}

// принимает массив переменных, возвращает строку этих переменных с 
// заданным префиксом
func beforeVar(mas []string, prefix string) string{
	text := ""

	for _, field := range mas{
		text += fmt.Sprintf(`%[2]s%[1]s, 
		`, field, prefix)
		
	}

	return text
}

// функция возвращающая строку $1, $2 до введенного 
func dollarAndInteger(length int) string {

	text := ``

	for i := 0; i < length; i++{
		if i == length - 1{
			text += fmt.Sprintf(`$%[1]d `, i+1)
		}else{
			text += fmt.Sprintf(`$%[1]d, `, i+1)
		}
	}

	return text
}

// генерирует строку "fieldName1"=$2, "fieldNama2"=$3...
func generatePatchFieldString(mas []string) string {
	text := ""

	for i, field := range mas{
		if field != "Id"{

			if i != len(mas)-1{
				text += fmt.Sprintf(`"%[1]s"=$%[2]d,
		`, field, i+1) 
				}else{
				text += fmt.Sprintf(`"%[1]s"=$%[2]d
		`, field, i+1) 
				}
		}
	}
	
	return text
}

// принимает массив полей, которые нужно будет зашифровать и возвращает код зашифровки
func getTextEncrypt(fields []string)string{
	text := ``

	for _, field := range fields{
		text += fmt.Sprintf(`
	data.%[1]s, err = utils.Encrypt(data.%[1]s)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"message": "Ошибка сервера",
		})
		return
	}
	`, field)
	}

	return text
}


func getRequestGenerate(tableName string, columns []string) string {

	lowerName := firstLower(tableName)
	routerText += fmt.Sprintf(`%[1]s.GET("/get%[2]ss", controllers.Get%[2]ss)
	`, lowerName, tableName)
	text := fmt.Sprintf(`
func Get%[1]ss(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	var %[2]sList []structs.%[1]s
	var %[2]s structs.%[1]s

	rows, err := db.Dbpool.Query(@SELECT * FROM "@+model.Schema+@"."%[1]s"@)
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
		%[3]s)
		%[2]sList = append(%[2]sList, %[2]s)
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
		"result": %[2]sList,
		"message": nil,
	})
}`, tableName, 
	lowerName, 
	beforeVar(columns, "&" + lowerName + "."))

	return text
}

func getByRequestGenerate(tableName string, fields, allColumns []string) string {

	lowerName := firstLower(tableName)
	text := ``
	for _, field := range fields{

	routerText += fmt.Sprintf(`%[2]s.GET("/get%[1]sBy%[3]s/:%[4]s", controllers.Get%[1]sBy%[3]s)
	`, tableName, lowerName, field, firstLower(field))

	text += fmt.Sprintf(`
func Get%[1]sBy%[3]s(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	%[4]s := c.Params.ByName("%[4]s")
	var %[2]s structs.%[1]s

	err := db.Dbpool.QueryRow(@SELECT * FROM "@+model.Schema+@"."%[1]s" WHERE "%[3]s"=$1@, %[4]s ).Scan(
		%[5]s
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
		"result": %[2]s,
		"message": nil,
	})
}

	`,tableName, lowerName, field, firstLower(field), beforeVar(allColumns, "&" + lowerName + "."))
	}

	return text
}
func getByManyRequestGenerate(tableName string, fields, allColumns []string) string {

	lowerName := firstLower(tableName)
	text := ``
	for _, field := range fields{

	routerText += fmt.Sprintf(`%[2]s.GET("/get%[1]sByMany%[3]s/:%[4]s", controllers.Get%[1]sBy%[3]s)
	`, tableName, lowerName, field, firstLower(field))

	text += fmt.Sprintf(`
func Get%[1]sBy%[3]s(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	%[4]s := c.Params.ByName("%[4]s")
	var %[2]sList []structs.%[1]s
	var %[2]s structs.%[1]s

	rows, err := db.Dbpool.Query(@SELECT * FROM "@+model.Schema+@"."%[1]s" WHERE "%[3]s"=$1@, %[4]s )
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
		%[5]s)
		%[2]sList = append(%[2]sList, %[2]s)
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
		"result": %[2]sList,
		"message": nil,
	})
}

	`,tableName, lowerName, field, firstLower(field), beforeVar(allColumns, "&" + lowerName + "."))
	}

	return text
}

func postRequestGenerate(tableName string, postColumns, encryptColumns []string) string {

	lowerName := firstLower(tableName)
	routerText += fmt.Sprintf(`%[1]s.POST("/create", controllers.Create%[2]s)
	`,lowerName, tableName )
	text := fmt.Sprintf(`
func Create%[1]s(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	data := DataProcessing%[1]s(*c)
	var err error
	%[6]s

	_, err = db.Dbpool.Exec(@INSERT INTO "@+model.Schema+@"."%[1]s"
		(
		%[3]s
		) 
		VALUES( %[4]s)@,
		%[5]s)
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
	`, tableName, lowerName, simbAroundVar(postColumns, `"`), dollarAndInteger(len(postColumns)), beforeVar(postColumns, "data."), getTextEncrypt(encryptColumns) )

	return text
}

func patchRequestGenerate(tableName string, allColumns, encryptColumns []string) string{

	routerText += fmt.Sprintf(`%[1]s.PATCH("/update", controllers.Update%[2]s)
	`, firstLower(tableName), tableName)
	text := fmt.Sprintf(`
func Update%[1]s(c *gin.Context) {

	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	data := DataProcessing%[1]s(*c)
	var err error
	%[2]s
	
	_, err = db.Dbpool.Exec(@UPDATE "@+model.Schema+@"."%[1]s" 
		SET 
		%[3]s
		WHERE "Id"=$1@,
		id,
		%[4]s
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
	`, tableName, getTextEncrypt(encryptColumns), generatePatchFieldString(allColumns[1:]), beforeVar(allColumns[1:], "data."),)

	return text
}

func deleteRequestGenerate(tableName string) string{

	routerText += fmt.Sprintf(`%[1]s.DELETE("/delete/:id", controllers.Delete%[2]s)
	}`, firstLower(tableName), tableName)
	text := fmt.Sprintf(`
func Delete%[1]s(c *gin.Context) {
	model := c.Value("Model").(structs.Claims)
	id := c.Params.ByName("id")
	_, err := db.Dbpool.Exec(@DELETE FROM "@+model.Schema+@"."%[1]s" WHERE "Id"=$1@, id)
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
	`, tableName)

	return text
}


func ControllerFileCreate(tableName string, columns ...string) error {

	routerText = fmt.Sprintf( `
	%[1]s := router.Group("/api/%[1]s")
	{
	`, firstLower(tableName), tableName)
	
	text := fmt.Sprintf(`
package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessing%[1]s(c gin.Context) structs.%[1]s {
	var data structs.%[1]s
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(400, gin.H{
			"message": "Некорректные данные",
		})
		return structs.%[1]s{}
	}
	return data
}

%[2]s
%[3]s
%[4]s
%[5]s
%[6]s
%[7]s
	`,
	tableName, 
	getRequestGenerate(tableName, sortColumn(columns, "")),
	getByRequestGenerate(tableName, sortColumn(columns, "*"), sortColumn(columns, "")),
	getByManyRequestGenerate(tableName, sortColumn(columns, "%"), sortColumn(columns, "")),
	postRequestGenerate(tableName, sortColumn(columns, "#"), sortColumn(columns, "$")),
	patchRequestGenerate(tableName, sortColumn(columns, ""), sortColumn(columns, "$")),
	deleteRequestGenerate(tableName))
	

	fmt.Println(routerText)

	
	file, err := os.Create(fmt.Sprintf(`./controllers/%[1]s.controllers.go`, tableName))
	if err != nil {
		Logger.Println(err)
		return err
	}
	fmt.Fprint(file, strings.ReplaceAll(text, "@", "`"))
	file.Close()
	
	return nil
}


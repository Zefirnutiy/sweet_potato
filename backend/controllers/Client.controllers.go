package controllers

import (
	"fmt"
	"io"
	"strings"

	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)
func DataProcessing(c gin.Context) structs.Client {
	var data structs.Client
	err := c.BindJSON(&data)
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"message": "Некорректные данные",
		})
		return structs.Client{}
	}
	return data
}

//создает строку patch запроса для клиента
func generateSqlForPatch(keys, values []string) string {
	text := ``

	for i := 1; i < len(keys); i++{

		if i != len(keys)-1{
			text += `"` + keys[i] + `"` + "=" + `'` + values[i] + `'` + ", "
		}else{
			text += `"` + keys[i] + `"` + "=" + `'` + values[i] + `'`
		}
	}

	return fmt.Sprintf(`UPDATE bebra."Client" SET %[1]s WHERE "Id"=%[2]s`, text, values[0])
}

// принимает массив и слово, индекс которого нужно найти в этом массиве.
// если такого слева нет возвращает 0
func find(a []string, x string) int {
    for i, n := range a {
        if x == n {
            return i
        }
    }
    return 0
}

// Функция принимающая в себя джин контекст, который представляет из себя JSON
// и возвращает два массива: первый Массив Ключей и второй Массив Значений.
func potomPridumau(c *gin.Context)([]string, []string, error){
	resp := c.Request

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
		return nil,  nil, err
	}

	bodyStr := string(bodyBytes)
	bodyStr = strings.ReplaceAll(bodyStr, " ", "")
	bodyStr = strings.ReplaceAll(bodyStr, ",", "")
	strMas := strings.Split(bodyStr, "\n")
	sliceStr := strMas[1:len(strMas)-1]

	var keys string
	var values string

	for _, val := range sliceStr{
		miniMas := strings.Split(val, ":")
		values += strings.ReplaceAll(miniMas[0], `"`, "") + " "
		keys += strings.ReplaceAll(miniMas[1], `"`, "") + " "
	}

	keysMas := strings.Fields(keys)
	valuesMas := strings.Fields(values)
	
	return  valuesMas, keysMas, nil
}


func GetClients(c *gin.Context) {
	var clientList []structs.Client
	var client structs.Client

	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	rows, err := db.Dbpool.Query(`SELECT * FROM "` + "KTK" + `"."Client"`)
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
			&client.GroupId,
			&client.Organization)
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
	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	id := c.Params.ByName("id")
	var client structs.Client

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+ "KTK"+`"."Client" WHERE "Id"=$1`, id).Scan(
		&client.Id,
		&client.FirstName,
		&client.LastName,
		&client.Patronymic,
		&client.Password,
		&client.Email,
		&client.Telephone,
		&client.EmailNotifications,
		&client.GroupId,
		&client.Organization,
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


func GetClientByGroupId(c *gin.Context) {
	tokken := c.Params.ByName("tokken")
	fmt.Println("tokken:", tokken)

	id := c.Params.ByName("id")
	var client structs.Client

	err := db.Dbpool.QueryRow(`SELECT * FROM "`+ "KTK"+`"."Client" WHERE "GroupId"=$1`, id).Scan(
		&client.Id,
		&client.FirstName,
		&client.LastName,
		&client.Patronymic,
		&client.Password,
		&client.Email,
		&client.Telephone,
		&client.EmailNotifications,
		&client.GroupId,
		&client.Organization,
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

func CreateClient(c *gin.Context) {
	data := DataProcessing(*c)
	encryptedPassword, err := utils.Encrypt(data.Password)
	if err != nil {
		utils.Logger.Println(data.Email, err)
		c.JSON(500, gin.H{
			"message": "Ошибка сервера",
		})
		return
	}

	_, err = db.Dbpool.Exec(`INSERT INTO "`+data.Organization+`"."Client"
							("FirstName", 
							"LastName", 
							"Patronymic", 
							"Password", 
							"Email", 
							"Telephone", 
							"EmailNotifications", 
							"GroupId", 
							"Organization") 
							VALUES( $1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		data.FirstName,
		data.LastName,
		data.Patronymic,
		encryptedPassword,
		data.Email,
		data.Telephone,
		data.EmailNotifications,
		data.GroupId,
		data.Organization)
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

func UpdateClient(c *gin.Context) { // Не готова

	keys, values, err := potomPridumau(c)

	fmt.Println(keys)
	if err != nil{
		utils.Logger.Println(err)
		return
	}

	if i := find(keys, "Password"); i != 0{

		values[i], err = utils.Encrypt(values[i])

		if err != nil {
			utils.Logger.Println(err)
			c.JSON(500, gin.H{
				"message": "Ошибка сервера",
			})
			return
		}

	}
	_, err = db.Dbpool.Exec(generateSqlForPatch(keys, values))
	if err != nil {
		utils.Logger.Println(err)
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"result": "Данные обновлены",
	})
}

func DeleteClient(c *gin.Context) {
	tokken := c.Params.ByName("tokken")
	id := c.Params.ByName("id") 
	fmt.Println(tokken)
	_, err := db.Dbpool.Exec(`DELETE FROM "`+"KTK"+`"."Client" WHERE "Id"=$1`, id)
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
	
	
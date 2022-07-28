package controllers
import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
)
func GetClient(c *gin.Context) error {
	rows, err := db.Dbpool.Query("SELECT * FROM \"Client\"") 
	if err != nil{
		utils.Logger.Println(err)
		return err
	}
	var client_list []structs.Client
	var client structs.Client
	for rows.Next(){
		err = rows.Scan(client)
		client_list = append(client_list, client)
		if err != nil{
			utils.Logger.Println(err)
			return err
		}
	}
	c.JSON(200, gin.H{
		"client's": client_list,
	})
	return nil
}

func GetClientById(c *gin.Context) error {
	Id := c.Param("Id")
	rows, err := db.Dbpool.Query("SELECT * FROM \"Client\" WHERE Id=$1", Id) 
	if err != nil{
		utils.Logger.Println(err)
		return err
	}
	var client structs.Client
	for rows.Next(){
		err = rows.Scan(client)
		if err != nil{
			utils.Logger.Println(err)
			return err
		}
	}
	c.JSON(200, gin.H{
		"Client": client,
	})
	return nil
}
		

func CreateClient(c *gin.Context) error {
	Id := c.PostForm("Id") 
		FirstName := c.PostForm("FirstName") 
		LastName := c.PostForm("LastName") 
		GroupId := c.PostForm("GroupId") 
		
	_, err := db.Dbpool.Query("INSERT INTO \"Client\"( \"Id\", \"FirstName\", \"LastName\", \"GroupId\",  ) VALUES( $1, $2, $3, $4,  )", Id, FirstName, LastName, GroupId, )
	if err != nil{
		utils.Logger.Println(err)
		return err
	}
	c.JSON(200, gin.H{
		"message": "",
	})
	return nil
}	
	

func UpdateClient(c *gin.Context) error {
	var client structs.Client
	Id := c.PostForm("Id") 
	FirstName := c.PostForm("FirstName") 
	LastName := c.PostForm("LastName") 
	GroupId := c.PostForm("GroupId") 
		
	
	if err := c.ShouldBindJSON(&client); err != nil {
		c.String(http.StatusBadGateway, err.Error())
	}
	
	_, err := db.Dbpool.Query("UPDATE Student SET Id=$1 FirstName=$2 LastName=$3 GroupId=$4 , WHERE id=$1", Id, FirstName, LastName, GroupId,  )
	if err != nil {
		utils.Logger.Println(err)
	}
	
	c.JSON(http.StatusOK, gin.H{
		"org": client,
	})
	return nil
}
	
func DeleteClient(c *gin.Context) error{
	id := c.Param("id")
	_, err := db.Dbpool.Query("DELETE FROM \"Client\" WHERE id=$1", id)
	if err != nil{
		utils.Logger.Println(err)
		return err
	}
	c.JSON(200, gin.H{
		"message": "",
	})
	return nil
}
	
	
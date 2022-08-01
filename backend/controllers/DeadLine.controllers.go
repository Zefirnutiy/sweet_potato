package controllers
import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
)
func GetDeadLine(c *gin.Context){

	rows, err := db.Dbpool.Query("SELECT * FROM \"DeadLine\"") 
	if err != nil{
		utils.Logger.Println(err)
		return
	}

	var deadline_list []structs.DeadLine
	var deadline structs.DeadLine

	for rows.Next(){
		err = rows.Scan(deadline)
		deadline_list = append(deadline_list, deadline)
		if err != nil{
			utils.Logger.Println(err)
			return
		}
	}
	c.JSON(200, gin.H{
		"deadline's": deadline_list,
	})

}

func GetDeadLineById(c *gin.Context){
	Id := c.Param("Id")
	rows, err := db.Dbpool.Query("SELECT * FROM \"DeadLine\" WHERE Id=$1", Id) 
	if err != nil{
		utils.Logger.Println(err)
		return
	}
	var deadline structs.DeadLine

	for rows.Next(){
		err = rows.Scan(deadline)
		if err != nil{
			utils.Logger.Println(err)
			return
		}
	}
	c.JSON(200, gin.H{
		"DeadLine": deadline,
	})

}
		

func CreateDeadLine(c *gin.Context){
	Id := c.PostForm("Id") 
	Date := c.PostForm("Date") 
	LevelId := c.PostForm("LevelId") 
	OrganizationId := c.PostForm("OrganizationId") 
	
	_, err := db.Dbpool.Query("INSERT INTO \"DeadLine\"( \"Id\", \"Date\", \"LevelId\", \"OrganizationId\",  ) VALUES( $1, $2, $3, $4,  )", Id, Date, LevelId, OrganizationId, )
	if err != nil{
		utils.Logger.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"message": "all good",
	})
}	
	

func UpdateDeadLine(c *gin.Context){
	var deadline structs.DeadLine
	
	if err := c.ShouldBindJSON(&deadline); err != nil {
		c.String(http.StatusBadGateway, err.Error())
	}
	
	_, err := db.Dbpool.Query("UPDATE Student SET Id=$1 Date=$2 LevelId=$3 OrganizationId=$4 , WHERE id=$1", &deadline.Id, &deadline.Date, &deadline.LevelId, &deadline.OrganizationId,  )
	if err != nil {
		utils.Logger.Println(err)
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": deadline,
	})
}
	
func DeleteDeadLine(c *gin.Context){
	id := c.Param("id")
	_, err := db.Dbpool.Query("DELETE FROM \"DeadLine\" WHERE id=$1", id)
	if err != nil{
		utils.Logger.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"message": "all good",
	})
}
	
	
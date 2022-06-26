package controllers

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/functions"
	"github.com/Zefirnutiy/sweet_potato.git/structs"
	"github.com/gin-gonic/gin"
)



func GetResultByStudentId(c *gin.Context){
	idStudent := c.Param("idStudent")

	row, err := db.Dbpool.Query(`SELECT * FROM "Result" WHERE student_id=$1`, idStudent)
	
	functions.HandlerError(err)
	
	var result structs.Result
	
	for row.Next(){
		e := row.Scan(&result.Id, &result.Student_id,  &result.Test_id, &result.Attempt, &result.Time_spent, &result.Score_gained, )
		
		functions.HandlerError(e)
	}

	c.JSON(200, gin.H{
		"Результат": result,
	})
}
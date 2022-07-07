package controllers

import (
	"fmt"

	"collage_project/backend/db"
	"collage_project/backend/structs"

	"collage_project/backend/functions"
	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {

	rows, e := db.Dbpool.Query(`SELECT * FROM "Student"`)

	functions.HandlerError(e)

	var student_list []structs.Student
	var student structs.Student

	for rows.Next() {
		e = rows.Scan(&student.Id, &student.First_name, &student.Sure_name, &student.Email, &student.Password, &student.Major, &student.Course, &student.Group)
		fmt.Println(student)
		student_list = append(student_list, student)
		functions.HandlerError(e)
	}

	c.JSON(200, gin.H{
		"Студенты": student_list,
	})

}

func GetStudentById(c *gin.Context) {
	id := c.Param("id")

	rows, e := db.Dbpool.Query(`SELECT * FROM "Student" WHERE id=$1`, id)

	functions.HandlerError(e)

	var student structs.Student

	for rows.Next() {
		e = rows.Scan(&student.Id, &student.First_name, &student.Sure_name, &student.Email, &student.Password, &student.Major, &student.Course, &student.Group)
		functions.HandlerError(e)
	}

	c.JSON(200, gin.H{
		"Студент": student,
	})
}

func GetStudentsByGroup(c *gin.Context) {
	group := c.Param("group")

	rows, e := db.Dbpool.Query(`SELECT * FROM "Student" WHERE id=$1`, group)

	functions.HandlerError(e)

	var student_list []structs.Student
	var student structs.Student

	for rows.Next() {
		e = rows.Scan(&student.Id, &student.First_name, &student.Sure_name, &student.Email, &student.Password, &student.Major, &student.Course, &student.Group)
		fmt.Println(student)
		student_list = append(student_list, student)
		functions.HandlerError(e)
	}

	c.JSON(200, gin.H{
		"Студенты": student_list,
	})

}

func CreateStudent(c *gin.Context) {
	first_name := c.PostForm("first_name")
	sure_name := c.PostForm("sure_name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	major := c.PostForm("major")
	course := c.PostForm("course")
	group := c.PostForm("group")

	_, err := db.Dbpool.Query(`INSERT INTO "Student"("first_name", "sure_name", "email", "password", "major", "course", "group") VALUES($1, $2, $3, $4, $5, $6, $7 )`, first_name, sure_name, email, password, major, course, group)

	functions.HandlerError(err)

	c.JSON(200, gin.H{
		"message": "Пользователь создан успешно.",
	})
}

func UpdateStudent(c *gin.Context) {

	id := c.Param("id")

	first_name := c.PostForm("first_name")
	sure_name := c.PostForm("sure_name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	major := c.PostForm("major")
	course := c.PostForm("course")
	group := c.PostForm("group")

	_, err := db.Dbpool.Query(`UPDATE "Student" SET first_name=$1, sure_name=$2, email=$3, password=$4, major=$5, course=$6, "group"=$7 WHERE id=$8`, first_name, sure_name, email, password, major, course, group, id)

	functions.HandlerError(err)

	c.JSON(200, gin.H{
		"message": "Данные пользователя обновлены успешно.",
	})
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Dbpool.Query(`DELETE FROM "Student" WHERE id=$1`, id)

	functions.HandlerError(err)

	c.JSON(200, gin.H{
		"message": "Пользователь удалён.",
	})
}

func SearchStudents(c *gin.Context) {
	// тут нужно придумать что и как делать
}

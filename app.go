package main

import (
	"github.com/Zefirnutiy/sweet_potato.git/controllers"
	"github.com/gin-gonic/gin"
)

func App(port string) {
	router := gin.Default()

	tests := router.Group("/api/test")
	{
		tests.GET("/", controllers.GetTests)
		tests.GET("/:id", controllers.GetTestsById)
		tests.POST("/create", controllers.CreateTest)
		tests.PATCH("/update/:id", controllers.UpdateTest)
		tests.DELETE("/delete/:id", controllers.DeleteTest)
		tests.GET("/search", controllers.SearchTests)

	}

	answers := router.Group("/api/answers")
	{
		answers.GET("/:question_id", controllers.GetAnswersByQuestionId)
		answers.POST("/create/:idQuestion", controllers.CreateAnswer)
		answers.PATCH("/update/:id", controllers.UpdateAnswer)
		answers.DELETE("/delete/:id", controllers.DeleteAnswer)
	}

	files := router.Group("/api/files")
	{
		files.GET("/", controllers.GetFiles)
		files.POST("/upload", controllers.UploadFile)
		files.DELETE("/delete/:id", controllers.DeleteFile)
	}

	questions := router.Group("/api/questions")
	{
		questions.GET("/", controllers.GetQuestions)
		questions.GET("/:id", controllers.GetQuestionsById)
		questions.POST("/create", controllers.CreateQuestion)
		questions.PATCH("/update/:id", controllers.UpdateQuestion)
		questions.DELETE("/delete/:id", controllers.DeleteQuestion)
		questions.GET("/search", controllers.SearchQuestions)
	}

	result := router.Group("/api/result")
	{
		result.GET("/:idStudent", controllers.GetResultByStudentId)
	}

	students := router.Group("/api/students")
	{
		students.GET("/", controllers.GetStudents)
		students.GET("/:id", controllers.GetStudentsById)
		students.GET("/group/:group", controllers.GetStudentsByGroup)
		students.POST("/create", controllers.CreateStudent)
		students.PATCH("/update/:id", controllers.UpdateStudent)
		students.DELETE("/delete/:id", controllers.DeleteStudent)
		students.GET("/search", controllers.SearchStudents)
	}

	router.Run(port)

}

package routs

import (
	controllers2 "github.com/Zefirnutiy/sweet_potato.git/backend/controllers"
	"github.com/gin-gonic/gin"
)

func Routs(port string) {
	router := gin.Default()

	tests := router.Group("/api/test")
	{
		tests.GET("/", controllers2.GetTests)
		tests.GET("/:id", controllers2.GetTestsById)
		tests.POST("/create", controllers2.CreateTest)
		tests.PATCH("/update/:id", controllers2.UpdateTest)
		tests.DELETE("/delete/:id", controllers2.DeleteTest)
		tests.GET("/search", controllers2.SearchTests)

	}

	answers := router.Group("/api/answers")
	{
		answers.GET("/:question_id", controllers2.GetAnswersByQuestionId)
		answers.POST("/create/:idQuestion", controllers2.CreateAnswer)
		answers.PATCH("/update/:id", controllers2.UpdateAnswer)
		answers.DELETE("/delete/:id", controllers2.DeleteAnswer)
	}

	files := router.Group("/api/files")
	{
		files.GET("/", controllers2.GetFiles)
		files.POST("/upload", controllers2.UploadFile)
		files.DELETE("/delete/:id", controllers2.DeleteFile)
	}

	questions := router.Group("/api/questions")
	{
		questions.GET("/", controllers2.GetQuestions)
		questions.GET("/:id", controllers2.GetQuestionsById)
		questions.POST("/create", controllers2.CreateQuestion)
		questions.PATCH("/update/:id", controllers2.UpdateQuestion)
		questions.DELETE("/delete/:id", controllers2.DeleteQuestion)
		questions.GET("/search", controllers2.SearchQuestions)
	}

	result := router.Group("/api/result")
	{
		result.GET("/:idStudent", controllers2.GetResultByStudentId)
	}

	students := router.Group("/api/students")
	{
		students.GET("/", controllers2.GetStudents)
		students.GET("/:id", controllers2.GetStudentById)
		students.GET("/group/:group", controllers2.GetStudentsByGroup)
		students.POST("/create", controllers2.CreateStudent)
		students.PATCH("/update/:id", controllers2.UpdateStudent)
		students.DELETE("/delete/:id", controllers2.DeleteStudent)
		students.GET("/search", controllers2.SearchStudents)
	}

	router.Run(port)

}

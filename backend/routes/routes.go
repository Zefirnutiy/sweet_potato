package routes

import (
	"github.com/Zefirnutiy/sweet_potato.git/controllers"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Routs(port string) {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"DELETE", "POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	organization := router.Group("/api/organization")
	{
		organization.POST("/register", controllers.RegisterOrganization)
		organization.POST("/login", controllers.LoginOrganization)
	}

	admin := router.Group("/api/admin")
	{
		admin.GET("/", utils.TokenCheckedFromHeader, controllers.GetAdmins)
		admin.GET("/:id", utils.TokenCheckedFromHeader, controllers.GetAdminById)
		admin.POST("/create", utils.TokenCheckedFromHeader, controllers.CreateAdmin)
		admin.PATCH("/update/:id", utils.TokenCheckedFromHeader, controllers.UpdateAdmin)
		admin.DELETE("/delete/:id", utils.TokenCheckedFromHeader, controllers.DeleteAdmin)
	}

	level := router.Group("/api/level")
	{
		level.GET("/", utils.TokenCheckedFromHeader, controllers.GetLevels)
		level.GET("/:id", utils.TokenCheckedFromHeader, controllers.GetLevelById)
		level.POST("/create", utils.TokenCheckedFromHeader, controllers.CreateLevel)
		level.PATCH("/update/:id", utils.TokenCheckedFromHeader, controllers.UpdateLevel)
		level.DELETE("/delete/:id", utils.TokenCheckedFromHeader, controllers.DeleteLevel)
	}

	client := router.Group("/api/client")
	{
		client.GET("/getClients/:schema", controllers.GetClients)
		client.GET("/getClientById/:schema/:id", controllers.GetClientById)
		client.GET("/getClientByGroupId/:schema/:id", controllers.GetClientByGroupId)
		client.GET("/getClientByClientLevelId/:schema/:id", controllers.GetClientByClientLevelId)
		client.POST("/create/:schema", controllers.CreateClient)
		client.PATCH("/update/:schema/:id", controllers.UpdateClient)
		client.DELETE("/delete/:schema/:id", controllers.DeleteClient)
		client.POST("/login", controllers.LoginClient)
	}

	// clientLevel := router.Group("/api/clientLevel")
	// {
	// 	clientLevel.GET("/getClientLevels/:schema", controllers.GetClientLevels)
	// 	clientLevel.GET("/getClientLevelById/:schema/:id", controllers.GetClientLevelById)
	// 	clientLevel.POST("/create/:schema", controllers.CreateClientLevel)
	// 	clientLevel.PATCH("/update/:schema", controllers.UpdateClientLevel)
	// 	clientLevel.DELETE("/delete/:schema/:id", controllers.DeleteClientLevel)
	// }

	// deadLine := router.Group("/api/deadLine")
	// {
	// 	deadLine.GET("/getDeadLines/", controllers.GetDeadLines)
	// 	deadLine.GET("/getDeadLineById/:id", controllers.GetDeadLineById)
	// 	deadLine.GET("/getDeadLineByOrganizationId/:organizationId", controllers.GetDeadLineByOrganizationId)
	// 	deadLine.POST("/create/", controllers.CreateDeadLine)
	// 	deadLine.PATCH("/update/", controllers.UpdateDeadLine)
	// 	deadLine.DELETE("/delete/:id", controllers.DeleteDeadLine)
	// }

	// payment := router.Group("/api/payment")
	// {
	// 	payment.GET("/getPayments/", controllers.GetPayments)
	// 	payment.GET("/getPaymentByNumber/:number", controllers.GetPaymentByNumber)
	// 	payment.GET("/getPaymentByOrganizationId/:organizationId", controllers.GetPaymentByOrganizationId)
	// 	payment.POST("/create/", controllers.CreatePayment)
	// 	payment.PATCH("/update/", controllers.UpdatePayment)
	// 	payment.DELETE("/delete/:id", controllers.DeletePayment)
	// }

	// test := router.Group("/api/test")
	// {
	// 	test.GET("/getTests/:schema", controllers.GetTests)
	// 	test.GET("/getTestById/:schema/:id", controllers.GetTestById)
	// 	test.POST("/create/:schema", controllers.CreateTest)
	// 	test.PATCH("/update/:schema", controllers.UpdateTest)
	// 	test.DELETE("/delete/:schema/:id", controllers.DeleteTest)
	// }

	// question := router.Group("/api/question")
	// {
	// 	question.GET("/getQuestions/:schema", controllers.GetQuestions)
	// 	question.GET("/getQuestionById/:schema/:id", controllers.GetQuestionById)
	// 	question.POST("/create/:schema", controllers.CreateQuestion)
	// 	question.PATCH("/update/:schema", controllers.UpdateQuestion)
	// 	question.DELETE("/delete/:schema/:id", controllers.DeleteQuestion)
	// }

	// questionResult := router.Group("/api/questionResult")
	// {
	// 	questionResult.GET("/getQuestionResults/:schema", controllers.GetQuestionResults)
	// 	questionResult.GET("/getQuestionResultById/:schema/:id", controllers.GetQuestionResultById)
	// 	questionResult.POST("/create/:schema", controllers.CreateQuestionResult)
	// 	questionResult.PATCH("/update/:schema", controllers.UpdateQuestionResult)
	// 	questionResult.DELETE("/delete/:schema/:id", controllers.DeleteQuestionResult)
	// }

	// questionType := router.Group("/api/questionType")
	// {
	// 	questionType.GET("/getQuestionTypes/:schema", controllers.GetQuestionTypes)
	// 	questionType.GET("/getQuestionTypeById/:schema/:id", controllers.GetQuestionTypeById)
	// 	questionType.POST("/create/:schema", controllers.CreateQuestionType)
	// 	questionType.PATCH("/update/:schema", controllers.UpdateQuestionType)
	// 	questionType.DELETE("/delete/:schema/:id", controllers.DeleteQuestionType)
	// }

	// answer := router.Group("/api/answer")
	// {
	// 	answer.GET("/getAnswers/:schema", controllers.GetAnswers)
	// 	answer.GET("/getAnswerById/:schema/:id", controllers.GetAnswerById)
	// 	answer.POST("/create/:schema", controllers.CreateAnswer)
	// 	answer.PATCH("/update/:schema", controllers.UpdateAnswer)
	// 	answer.DELETE("/delete/:schema/:id", controllers.DeleteAnswer)
	// }

	// course := router.Group("/api/course")
	// {
	// 	course.GET("/getCourses/:schema", controllers.GetCourses)
	// 	course.GET("/getCourseById/:schema/:id", controllers.GetCourseById)
	// 	course.POST("/create/:schema", controllers.CreateCourse)
	// 	course.PATCH("/update/:schema", controllers.UpdateCourse)
	// 	course.DELETE("/delete/:schema/:id", controllers.DeleteCourse)
	// }

	// publicInfo := router.Group("/api/publicInfo")
	// {
	// 	publicInfo.GET("/getPublicInfos/:schema", controllers.GetPublicInfos)
	// 	publicInfo.GET("/getPublicInfoById/:schema/:id", controllers.GetPublicInfoById)
	// 	publicInfo.POST("/create/:schema", controllers.CreatePublicInfo)
	// 	publicInfo.PATCH("/update/:schema", controllers.UpdatePublicInfo)
	// 	publicInfo.DELETE("/delete/:schema/:id", controllers.DeletePublicInfo)
	// }

	// courseResults := router.Group("/api/courseResults")
	// {
	// 	courseResults.GET("/getCourseResultss/:schema", controllers.GetCourseResultss)
	// 	courseResults.GET("/getCourseResultsById/:schema/:id", controllers.GetCourseResultsById)
	// 	courseResults.POST("/create/:schema", controllers.CreateCourseResults)
	// 	courseResults.PATCH("/update/:schema", controllers.UpdateCourseResults)
	// 	courseResults.DELETE("/delete/:schema/:id", controllers.DeleteCourseResults)
	// }

	// department := router.Group("/api/department")
	// {
	// 	department.GET("/getDepartments/:schema", controllers.GetDepartments)
	// 	department.GET("/getDepartmentById/:schema/:id", controllers.GetDepartmentById)
	// 	department.POST("/create/:schema", controllers.CreateDepartment)
	// 	department.PATCH("/update/:schema", controllers.UpdateDepartment)
	// 	department.DELETE("/delete/:schema/:id", controllers.DeleteDepartment)
	// }

	// group := router.Group("/api/group")
	// {
	// 	group.GET("/getGroups/:schema", controllers.GetGroups)
	// 	group.GET("/getGroupById/:schema/:id", controllers.GetGroupById)
	// 	group.GET("/getGroupByDepartmentId/:schema/:departmentId", controllers.GetGroupByDepartmentId)
	// 	group.POST("/create/:schema", controllers.CreateGroup)
	// 	group.PATCH("/update/:schema", controllers.UpdateGroup)
	// 	group.DELETE("/delete/:schema/:id", controllers.DeleteGroup)
	// }

	// file := router.Group("/api/file")
	// {
	// 	file.GET("/getFiles/:schema", controllers.GetFiles)
	// 	file.GET("/getFileById/:schema/:id", controllers.GetFileById)
	// 	file.POST("/create/:schema", controllers.CreateFile)
	// 	file.PATCH("/update/:schema", controllers.UpdateFile)
	// 	file.DELETE("/delete/:schema/:id", controllers.DeleteFile)
	// }

	// activeTest := router.Group("/api/activeTest")
	// {
	// 	activeTest.GET("/getActiveTests/:schema", controllers.GetActiveTests)
	// 	activeTest.GET("/getActiveTestById/:schema/:id", controllers.GetActiveTestById)
	// 	activeTest.GET("/getActiveTestByClientId/:schema/:clientId", controllers.GetActiveTestByClientId)
	// 	activeTest.POST("/create/:schema", controllers.CreateActiveTest)
	// 	activeTest.PATCH("/update/:schema", controllers.UpdateActiveTest)
	// 	activeTest.DELETE("/delete/:schema/:id", controllers.DeleteActiveTest)
	// }

	// testResults := router.Group("/api/testResults")
	// {
	// 	testResults.GET("/getTestResultss/:schema", controllers.GetTestResultss)
	// 	testResults.GET("/getTestResultsById/:schema/:id", controllers.GetTestResultsById)
	// 	testResults.GET("/getTestResultsByClientId/:schema/:clientId", controllers.GetTestResultsByClientId)
	// 	testResults.GET("/getTestResultsByCourseId/:schema/:courseId", controllers.GetTestResultsByCourseId)
	// 	testResults.POST("/create/:schema", controllers.CreateTestResults)
	// 	testResults.PATCH("/update/:schema", controllers.UpdateTestResults)
	// 	testResults.DELETE("/delete/:schema/:id", controllers.DeleteTestResults)
	// }

	router.Run(port)

}

package routes

import (
	"github.com/Zefirnutiy/sweet_potato.git/controllers"
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	router.MaxMultipartMemory = 8 << 20  
	
	// organization := router.Group("/api/organization")
	// {
	// 	organization.POST("/register", controllers.RegisterOrganization)
	// 	organization.POST("/login", controllers.LoginOrganization) 
	// }

	file := router.Group("/api/file")
        {
        file.GET("/getFilesg", controllers.GetFiles)
        file.GET("/getFileById/:id", controllers.GetFileById)
        file.GET("/getFileByManyClientId/:clientId", controllers.GetFileByClientId)
        file.POST("/upload", controllers.UploadFile)
        file.PATCH("/update", controllers.UpdateFile)
        file.DELETE("/delete/:id", controllers.DeleteFile)
        }

	// admin := router.Group("/api/admin")
	// {
	// 	admin.GET("/", utils.TokenCheckedFromHeader, controllers.GetAdmins)
	// 	admin.GET("/:id", utils.TokenCheckedFromHeader, controllers.GetAdminById)
	// 	admin.POST("/create", utils.TokenCheckedFromHeader, controllers.CreateAdmin)
	// 	admin.PATCH("/update/:id", utils.TokenCheckedFromHeader, controllers.UpdateAdmin)
	// 	admin.DELETE("/delete/:id", utils.TokenCheckedFromHeader, controllers.DeleteAdmin)
	// }

	// level := router.Group("/api/level")
	// {
	// 	level.GET("/", utils.TokenCheckedFromHeader, controllers.GetLevels)
	// 	level.GET("/:id", utils.TokenCheckedFromHeader, controllers.GetLevelById)
	// 	level.POST("/create", utils.TokenCheckedFromHeader, controllers.CreateLevel)
	// 	level.PATCH("/update/:id", utils.TokenCheckedFromHeader, controllers.UpdateLevel)
	// 	level.DELETE("/delete/:id", utils.TokenCheckedFromHeader, controllers.DeleteLevel)
	// }

	// client := router.Group("/api/client")
	// {
	// 	client.GET("/getClients", controllers.GetClients)
	// 	client.GET("/getClientById/:id", controllers.GetClientById)
	// 	client.GET("/getClientByGroupId/:id", controllers.GetClientByGroupId)
	// 	client.GET("/getClientByClientLevelId/:id", controllers.GetClientByClientLevelId)
	// 	client.POST("/create", controllers.CreateClient)
	// 	client.PATCH("/update/:id", controllers.UpdateClient)
	// 	client.DELETE("/delete/:id", controllers.DeleteClient)
	// 	client.POST("/login", controllers.LoginClient)
	// }

	// clientLevel := router.Group("/api/clientLevel")
	// {
	// 	clientLevel.GET("/getClientLevels", controllers.GetClientLevels)
	// 	clientLevel.GET("/getClientLevelById/:id", controllers.GetClientLevelById)
	// 	clientLevel.POST("/create", controllers.CreateClientLevel)
	// 	clientLevel.PATCH("/update", controllers.UpdateClientLevel)
	// 	clientLevel.DELETE("/delete/:id", controllers.DeleteClientLevel)
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
	// 	test.GET("/getTests", controllers.GetTests)
	// 	test.GET("/getTestById/:id", controllers.GetTestById)
	// 	test.POST("/create", controllers.CreateTest)
	// 	test.PATCH("/update", controllers.UpdateTest)
	// 	test.DELETE("/delete/:id", controllers.DeleteTest)
	// }

	// question := router.Group("/api/question")
	// {
	// 	question.GET("/getQuestions", controllers.GetQuestions)
	// 	question.GET("/getQuestionById/:id", controllers.GetQuestionById)
	// 	question.POST("/create", controllers.CreateQuestion)
	// 	question.PATCH("/update", controllers.UpdateQuestion)
	// 	question.DELETE("/delete/:id", controllers.DeleteQuestion)
	// }

	// questionResult := router.Group("/api/questionResult")
	// {
	// 	questionResult.GET("/getQuestionResults", controllers.GetQuestionResults)
	// 	questionResult.GET("/getQuestionResultById/:id", controllers.GetQuestionResultById)
	// 	questionResult.POST("/create", controllers.CreateQuestionResult)
	// 	questionResult.PATCH("/update", controllers.UpdateQuestionResult)
	// 	questionResult.DELETE("/delete/:id", controllers.DeleteQuestionResult)
	// }

	// questionType := router.Group("/api/questionType")
	// {
	// 	questionType.GET("/getQuestionTypes", controllers.GetQuestionTypes)
	// 	questionType.GET("/getQuestionTypeById/:id", controllers.GetQuestionTypeById)
	// 	questionType.POST("/create", controllers.CreateQuestionType)
	// 	questionType.PATCH("/update", controllers.UpdateQuestionType)
	// 	questionType.DELETE("/delete/:id", controllers.DeleteQuestionType)
	// }

	// answer := router.Group("/api/answer")
	// {
	// 	answer.GET("/getAnswers", controllers.GetAnswers)
	// 	answer.GET("/getAnswerById/:id", controllers.GetAnswerById)
	// 	answer.POST("/create", controllers.CreateAnswer)
	// 	answer.PATCH("/update", controllers.UpdateAnswer)
	// 	answer.DELETE("/delete/:id", controllers.DeleteAnswer)
	// }

	// course := router.Group("/api/course")
	// {
	// 	course.GET("/getCourses", controllers.GetCourses)
	// 	course.GET("/getCourseById/:id", controllers.GetCourseById)
	// 	course.POST("/create", controllers.CreateCourse)
	// 	course.PATCH("/update", controllers.UpdateCourse)
	// 	course.DELETE("/delete/:id", controllers.DeleteCourse)
	// }

	// publicInfo := router.Group("/api/publicInfo")
	// {
	// 	publicInfo.GET("/getPublicInfos", controllers.GetPublicInfos)
	// 	publicInfo.GET("/getPublicInfoById/:id", controllers.GetPublicInfoById)
	// 	publicInfo.POST("/create", controllers.CreatePublicInfo)
	// 	publicInfo.PATCH("/update", controllers.UpdatePublicInfo)
	// 	publicInfo.DELETE("/delete/:id", controllers.DeletePublicInfo)
	// }

	// courseResults := router.Group("/api/courseResults")
	// {
	// 	courseResults.GET("/getCourseResultss", controllers.GetCourseResultss)
	// 	courseResults.GET("/getCourseResultsById/:id", controllers.GetCourseResultsById)
	// 	courseResults.POST("/create", controllers.CreateCourseResults)
	// 	courseResults.PATCH("/update", controllers.UpdateCourseResults)
	// 	courseResults.DELETE("/delete/:id", controllers.DeleteCourseResults)
	// }

	// department := router.Group("/api/department")
	// {
	// 	department.GET("/getDepartments", controllers.GetDepartments)
	// 	department.GET("/getDepartmentById/:id", controllers.GetDepartmentById)
	// 	department.POST("/create", controllers.CreateDepartment)
	// 	department.PATCH("/update", controllers.UpdateDepartment)
	// 	department.DELETE("/delete/:id", controllers.DeleteDepartment)
	// }

	// group := router.Group("/api/group")
	// {
	// 	group.GET("/getGroups", controllers.GetGroups)
	// 	group.GET("/getGroupById/:id", controllers.GetGroupById)
	// 	group.GET("/getGroupByDepartmentId/:departmentId", controllers.GetGroupByDepartmentId)
	// 	group.POST("/create", controllers.CreateGroup)
	// 	group.PATCH("/update", controllers.UpdateGroup)
	// 	group.DELETE("/delete/:id", controllers.DeleteGroup)
	// }

	// file := router.Group("/api/file")
	// {
	// 	file.GET("/getFiles", controllers.GetFiles)
	// 	file.GET("/getFileById/:id", controllers.GetFileById)
	// 	file.POST("/create", controllers.CreateFile)
	// 	file.PATCH("/update", controllers.UpdateFile)
	// 	file.DELETE("/delete/:id", controllers.DeleteFile)
	// }

	// activeTest := router.Group("/api/activeTest")
	// {
	// 	activeTest.GET("/getActiveTests", controllers.GetActiveTests)
	// 	activeTest.GET("/getActiveTestById/:id", controllers.GetActiveTestById)
	// 	activeTest.GET("/getActiveTestByClientId/:clientId", controllers.GetActiveTestByClientId)
	// 	activeTest.POST("/create", controllers.CreateActiveTest)
	// 	activeTest.PATCH("/update", controllers.UpdateActiveTest)
	// 	activeTest.DELETE("/delete/:id", controllers.DeleteActiveTest)
	// }

	// testResults := router.Group("/api/testResults")
	// {
	// 	testResults.GET("/getTestResultss", controllers.GetTestResultss)
	// 	testResults.GET("/getTestResultsById/:id", controllers.GetTestResultsById)
	// 	testResults.GET("/getTestResultsByClientId/:clientId", controllers.GetTestResultsByClientId)
	// 	testResults.GET("/getTestResultsByCourseId/:courseId", controllers.GetTestResultsByCourseId)
	// 	testResults.POST("/create", controllers.CreateTestResults)
	// 	testResults.PATCH("/update", controllers.UpdateTestResults)
	// 	testResults.DELETE("/delete/:id", controllers.DeleteTestResults)
	// }

	router.Run(port)

}

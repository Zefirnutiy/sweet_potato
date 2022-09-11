package routes

import (

	"github.com/Zefirnutiy/sweet_potato.git/controllers"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
    // TO allow CORS
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    }
}

func Routs(port string) {
	router := gin.Default()

	router.Use(CORS())
	
	organization := router.Group("/auth")
	{
		go organization.POST("/register", controllers.RegisterOrganization)
		go organization.POST("/login", controllers.LoginOrganization) 
	}

    api := router.Group("/api", utils.TokenCheckedFromHeader)
    {

        file := api.Group("/file")
            {
                file.GET("/getFilesg", controllers.GetFiles)
                file.GET("/getFileById/:id", controllers.GetFileById)
                file.GET("/getFileByManyClientId/:clientId", controllers.GetFileByClientId)
                file.POST("/upload", utils.TokenCheckedFromHeader, controllers.UploadFile)
                file.PATCH("/update", controllers.UpdateFile)
                file.DELETE("/delete/:id", controllers.DeleteFile)
            }
            client := api.Group("/client")
            {
                client.GET("/getClients", controllers.GetClients)
                client.GET("/getClientById/:id", controllers.GetClientById)
                client.GET("/getClientByManyDepartmentId/:departmentId", controllers.GetClientByDepartmentId)
                client.GET("/getClientByManyGroupId/:groupId", controllers.GetClientByGroupId)
                client.GET("/getClientByManyCreatorId/:creatorId", controllers.GetClientByCreatorId)
                client.POST("/create", controllers.CreateClient)
                client.PATCH("/update", controllers.UpdateClient)
                client.DELETE("/delete/:id", controllers.DeleteClient)
            }

            course := api.Group("/course")
            {
                course.GET("/getCourses", controllers.GetCourses)
                course.GET("/getCourseById/:id", controllers.GetCourseById)
                course.GET("/getCourseByClientId/:clientId", controllers.GetCourseByClientId)
                course.POST("/create", controllers.CreateCourse)
                course.PATCH("/update", controllers.UpdateCourse)
                course.DELETE("/delete/:id", controllers.DeleteCourse)
            }

            courseResults := api.Group("/courseResults")
            {
                courseResults.GET("/getCourseResultss", controllers.GetCourseResultss)
                courseResults.GET("/getCourseResultsById/:id", controllers.GetCourseResultsById)
                courseResults.GET("/getCourseResultsByCourseId/:courseId", controllers.GetCourseResultsByCourseId)
                courseResults.GET("/getCourseResultsByManyClientId/:clientId", controllers.GetCourseResultsByClientId)
                courseResults.POST("/create", controllers.CreateCourseResults)
                courseResults.PATCH("/update", controllers.UpdateCourseResults)
                courseResults.DELETE("/delete/:id", controllers.DeleteCourseResults)
            }

            activeCourse := api.Group("/activeCourse")
            {
                activeCourse.GET("/getActiveCourses", controllers.GetActiveCourses)
                activeCourse.GET("/getActiveCourseById/:id", controllers.GetActiveCourseById)
                activeCourse.GET("/getActiveCourseByCourseId/:courseId", controllers.GetActiveCourseByCourseId)
                activeCourse.GET("/getActiveCourseByManyClientId/:clientId", controllers.GetActiveCourseByClientId)
                activeCourse.POST("/create", controllers.CreateActiveCourse)
                activeCourse.PATCH("/update", controllers.UpdateActiveCourse)
                activeCourse.DELETE("/delete/:id", controllers.DeleteActiveCourse)
            }

            department := api.Group("/department")
            {
                department.GET("/getDepartments",  controllers.GetDepartments)
                department.GET("/getDepartmentById/:id", controllers.GetDepartmentById)
                department.POST("/create", controllers.CreateDepartment)
                department.PATCH("/update", controllers.UpdateDepartment)
                department.DELETE("/delete/:id", controllers.DeleteDepartment)
            }

            group := api.Group("/group")
            {
                group.GET("/getGroups", controllers.GetGroups)
                group.GET("/getGroupById/:id", controllers.GetGroupById)
                group.GET("/getGroupByDepartmentId/:departmentId", controllers.GetGroupByDepartmentId)
                group.POST("/create", controllers.CreateGroup)
                group.PATCH("/update", controllers.UpdateGroup)
                group.DELETE("/delete/:id", controllers.DeleteGroup)
            }

            fileInformation := api.Group("/fileInformation")
            {
                fileInformation.GET("/getFileInformations", controllers.GetFileInformations)
                fileInformation.GET("/getFileInformationByFileId/:fileId", controllers.GetFileInformationByFileId)
                fileInformation.GET("/getFileInformationByTestId/:testId", controllers.GetFileInformationByTestId)
                fileInformation.GET("/getFileInformationByQuestionId/:questionId", controllers.GetFileInformationByQuestionId)
                fileInformation.GET("/getFileInformationByManyClientId/:clientId", controllers.GetFileInformationByClientId)
                fileInformation.GET("/getFileInformationByManyPublicInfoId/:publicInfoId", controllers.GetFileInformationByPublicInfoId)
                fileInformation.GET("/getFileInformationByManyTestId/:testId", controllers.GetFileInformationByTestId)
                fileInformation.POST("/create", controllers.CreateFileInformation)
                fileInformation.PATCH("/update", controllers.UpdateFileInformation)
                fileInformation.DELETE("/delete/:id", controllers.DeleteFileInformation)
            }

            publicInfo := api.Group("/publicInfo")
            {
                publicInfo.GET("/getPublicInfos", controllers.GetPublicInfos)
                publicInfo.GET("/getPublicInfoById/:id", controllers.GetPublicInfoById)
                publicInfo.GET("/getPublicInfoByManyClientId/:clientId", controllers.GetPublicInfoByClientId)
                publicInfo.POST("/create", controllers.CreatePublicInfo)
                publicInfo.PATCH("/update", controllers.UpdatePublicInfo)
                publicInfo.DELETE("/delete/:id", controllers.DeletePublicInfo)
            }

            test := api.Group("/test")
            {
                test.GET("/getTests", controllers.GetTests)
                test.GET("/getTestById/:id", controllers.GetTestById)
                test.POST("/create", controllers.CreateTest)
                test.PATCH("/update", controllers.UpdateTest)
                test.DELETE("/delete/:id", controllers.DeleteTest)
            }

            testResults := api.Group("/testResults")
            {
                testResults.GET("/getTestResultss", controllers.GetTestResultss)
                testResults.GET("/getTestResultsById/:id", controllers.GetTestResultsById)
                testResults.GET("/getTestResultsByTestId/:testId", controllers.GetTestResultsByTestId)
                testResults.GET("/getTestResultsByManyCourseId/:courseId", controllers.GetTestResultsByCourseId)
                testResults.GET("/getTestResultsByManyClientId/:clientId", controllers.GetTestResultsByClientId)
                testResults.POST("/create", controllers.CreateTestResults)
                testResults.PATCH("/update", controllers.UpdateTestResults)
                testResults.DELETE("/delete/:id", controllers.DeleteTestResults)
            }

            activeTest := api.Group("/activeTest")
            {
                activeTest.GET("/getActiveTests", controllers.GetActiveTests)
                activeTest.GET("/getActiveTestById/:id", controllers.GetActiveTestById)
                activeTest.GET("/getActiveTestByTestId/:testId", controllers.GetActiveTestByTestId)
                activeTest.GET("/getActiveTestByClientId/:clientId", controllers.GetActiveTestByClientId)
                activeTest.GET("/getActiveTestByManyTestTypeId/:testTypeId", controllers.GetActiveTestByTestTypeId)
                activeTest.POST("/create", controllers.CreateActiveTest)
                activeTest.PATCH("/update", controllers.UpdateActiveTest)
                activeTest.DELETE("/delete/:id", controllers.DeleteActiveTest)
            }

            testType := api.Group("/testType")
            {
                testType.GET("/getTestTypes", controllers.GetTestTypes)
                testType.GET("/getTestTypeById/:id", controllers.GetTestTypeById)
                testType.POST("/create", controllers.CreateTestType)
                testType.PATCH("/update", controllers.UpdateTestType)
                testType.DELETE("/delete/:id", controllers.DeleteTestType)
            }

            question := api.Group("/question")
            {
                question.GET("/getQuestions", controllers.GetQuestions)
                question.GET("/getQuestionById/:id", controllers.GetQuestionById)
                question.GET("/getQuestionByManyTestId/:testId", controllers.GetQuestionByTestId)
                question.POST("/create", controllers.CreateQuestion)
                question.PATCH("/update", controllers.UpdateQuestion)
                question.DELETE("/delete/:id", controllers.DeleteQuestion)
            }

            questionResult := api.Group("/questionResult")
            {
                questionResult.GET("/getQuestionResults", controllers.GetQuestionResults)
                questionResult.GET("/getQuestionResultById/:id", controllers.GetQuestionResultById)
                questionResult.GET("/getQuestionResultByQuestionId/:questionId", controllers.GetQuestionResultByQuestionId)
                questionResult.POST("/create", controllers.CreateQuestionResult)
                questionResult.PATCH("/update", controllers.UpdateQuestionResult)
                questionResult.DELETE("/delete/:id", controllers.DeleteQuestionResult)
            }

            questionType := api.Group("/questionType")
            {
                questionType.GET("/getQuestionTypes", controllers.GetQuestionTypes)
                questionType.GET("/getQuestionTypeById/:id", controllers.GetQuestionTypeById)
                questionType.GET("/getQuestionTypeByInputTypeId/:inputTypeId", controllers.GetQuestionTypeByInputTypeId)
                questionType.POST("/create", controllers.CreateQuestionType)
                questionType.PATCH("/update", controllers.UpdateQuestionType)
                questionType.DELETE("/delete/:id", controllers.DeleteQuestionType)
            }

            inputType := api.Group("/inputType")
            {
                inputType.GET("/getInputTypes", controllers.GetInputTypes)
                inputType.GET("/getInputTypeById/:id", controllers.GetInputTypeById)
                inputType.POST("/create", controllers.CreateInputType)
                inputType.PATCH("/update", controllers.UpdateInputType)
                inputType.DELETE("/delete/:id", controllers.DeleteInputType)
            }

            theme := api.Group("/theme")
            {
                theme.GET("/getThemes", controllers.GetThemes)
                theme.GET("/getThemeById/:id", controllers.GetThemeById)
                theme.POST("/create", controllers.CreateTheme)
                theme.PATCH("/update", controllers.UpdateTheme)
                theme.DELETE("/delete/:id", controllers.DeleteTheme)
            }

            payment := api.Group("/Payment")
            {
                payment.GET("/getPayments", controllers.GetPayments)
                payment.GET("/getPaymentByNumber/:number", controllers.GetPaymentByNumber)
                payment.GET("/getPaymentByManyOrganizationId/:organizationId", controllers.GetPaymentByOrganizationId)
                payment.POST("/create", controllers.CreatePayment)
                payment.PATCH("/update", controllers.UpdatePayment)
                payment.DELETE("/delete/:id", controllers.DeletePayment)
            }

            deadLine := api.Group("/deadLine")
            {
                deadLine.GET("/getDeadLines", controllers.GetDeadLines)
                deadLine.GET("/getDeadLineById/:id", controllers.GetDeadLineById)
                deadLine.GET("/getDeadLineByOrganizationId/:organizationId", controllers.GetDeadLineByOrganizationId)
                deadLine.POST("/create", controllers.CreateDeadLine)
                deadLine.PATCH("/update", controllers.UpdateDeadLine)
                deadLine.DELETE("/delete/:id", controllers.DeleteDeadLine)
            }

            session := api.Group("/session")
            {
                session.GET("/getSessions", controllers.GetSessions)
                session.GET("/getSessionById/:id", controllers.GetSessionById)
                session.GET("/getSessionByUserId/:userId", controllers.GetSessionByUserId)
                session.POST("/create", controllers.CreateSession)
                session.PATCH("/update", controllers.UpdateSession)
                session.DELETE("/delete/:id", controllers.DeleteSession)
            }

            state := api.Group("/state")
            {
                state.GET("/getStates", controllers.GetStates)
                state.GET("/getStateById/:id", controllers.GetStateById)
                state.POST("/create", controllers.CreateState)
                state.PATCH("/update", controllers.UpdateState)
                state.DELETE("/delete/:id", controllers.DeleteState)
            }
    }
	router.Run(port)

}

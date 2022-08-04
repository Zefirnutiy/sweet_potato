package utils


func Generate(){
	ControllerFileCreate("Level",
		"*Id", "#Title", "#Price", "#Paid", "#CreateCourse", "#TakeCourse",
		"#AploadFile", "#ViewYourResult", "#ViewOtherResults")

	ControllerFileCreate("ClientLevel",
		"*Id", "#Title", "#CreateCourse", "#TakeCourse", 
		"#AploadFile", "#ViewYourResult", "#ViewOtherResults")	
		
	// ControllerFileCreate("DeadLine",
	// 	"*Id", "#Date", "#LevelId", "*#OrganizationId")

	// ControllerFileCreate("Payment",
	// 	"*#Number", "#Name", "#Money", "#Date", "#LevelId", "*#OrganizationId")
	
	ControllerFileCreate("Client",
		"*Id", "#FirstName", "#LastName", "#Patronymic", "$Password", "#Email",
		"#Telephone", "#EmailNotifications", "*#GroupId", "*#ClientLevelId")

	// ControllerFileCreate("Test",
	// 	"*Id", "#Title", "#Text", "#Date", "#DateDel")

	// ControllerFileCreate("Admin",
	// 	"*Id", "#FirstName", "#LastName", "$Password", "#Email")

	// ControllerFileCreate("Question",
	// 	"*Id", "#Text", "#Date", "#DateDel", "#TestId", 
	// 	"#QuestionTypeId","#Hint")

	// ControllerFileCreate("QuestionResult",
	// "*Id", "#Date", "#QuestionId", "#ClientId", "#Scores")

	ControllerFileCreate("QuestionType",
		"*Id", "#Type")

	ControllerFileCreate("Answer",
		"*Id", "#Text", "#Correct", "#QuestionId")

	// ControllerFileCreate("Course",
	// 	"*Id", "#Text", "#Date", "#DateDel", "#ClientId")
	
	// ControllerFileCreate("PublicInfo",
	// 	"*Id", "#Title", "#Text", "#Date", "#DateDel", "#ClientId")
	
	// ControllerFileCreate("CourseResults",
	// 	"*Id", "#Time", "#Date", "#Assessment", "#Scores", "#ClientId")
	
	ControllerFileCreate("Department",
		"*Id", "#Title")

	ControllerFileCreate("Group",
		"*Id", "#Title", "*#DepartmentId")

	// ControllerFileCreate("File", Переделать
	// 	"*Id", "#Date", "#DateDel", "#FileName", "#FileNameTmp", "#PublicInfoId",
	// 	"#TestId", "#QuestionId", "#ClientId")

	ControllerFileCreate("ActiveTest",
		"*Id", "#Start", "#End", "#Time", "#Attempts", "#TestId",
		"*#ClientId", "#TrainingTest")

	ControllerFileCreate("TestResults",
		"*Id", "#Time", "#Date", "*#ClientId", "#TestId", "#Assessment",
		"#PassageTime", "#Scores", "*#CourseId")
	
}

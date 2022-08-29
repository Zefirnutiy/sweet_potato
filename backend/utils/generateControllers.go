package utils


func Generate(){
	
	ControllerFileCreate("Client",
		"*Id", "#FirstName", "#LastName", "#Patronymic", "#Password", 
		"#Email", "#Telephone", "#EmailNotifications", "#ManageCourses", 
		"#ManageUsers", "#AploadFiles", "#ViewYourResults", "#ViewOtherResults",
		"%#DepartmentId", "%%#GroupId", "%#CreatorId")

	ControllerFileCreate("Course",
		"*Id", "#Title", "#Text", "#Files", "#Date", "#Time",
		"DateDel", "TimeDel", "#*ClientId")

	ControllerFileCreate("CourseResults",
		"*Id", "#Assessment", "#Scores", "#NumberTests", "#PassageTime",
		"#Date", "#Time", "%#ClientId",	"#*CourseId")

	ControllerFileCreate("ActiveCourse",
		"*Id", "#Date", "#Time", "DateClose", "TimeClose", "*CourseId",
		"%#ClientId")

	ControllerFileCreate("Department",
		"*Id", "#Title")	

	ControllerFileCreate("Group",
		"*Id", "#Title", "#TitleSingular", "*#DepartmentId")	

	ControllerFileCreate("FileInformation",
		"%#ClientId", "*#FileId", "%#PublicInfoId", "%%#*TestId", "*#QuestionId")

	ControllerFileCreate("PublicInfo",
		"*Id", "#Title", "#Text", "#Files", "#Date", "#Time", "DateDel", 
		"TimeDel", "%#ClientId",)	

	ControllerFileCreate("Test",
		"*Id", "#Title", "#Text", "#Files", "#Date", "#Time",
		"DateDel", "TimeDel")

	ControllerFileCreate("TestResults",
		"*Id", "#PassageTime", "#Assessment", "#Scores", "#Attempts", "#Date", 
		"#Time", "%#CourseId", "*#TestId", "%#ClientId")

	ControllerFileCreate("ActiveTest",
		"*Id", "#Date", "#Time", "DateClose", "TimeClose", "#Attempts",
		"%%#TestTypeId", "*#TestId", "*#ClientId")

	ControllerFileCreate("TestType",
		"*Id", "#Title")

	ControllerFileCreate("Question",
		"*Id", "#Text", "#Date", "#Time", "DateDel", "TimeDel", "#Hint",
		"#AnswerVariant", "#AnswerCorrect", "#Files", "%%#TestId", "#QuestionTypeId")

	ControllerFileCreate("QuestionResult",
		"*Id", "#Date", "#Time", "#Scores", "*#QuestionId", "#ClientId")

	ControllerFileCreate("QuestionType",
		"*Id", "#Title", "*#InputTypeId")

	ControllerFileCreate("InputType",
		"*Id", "#Title", "#Type")	

	ControllerFileCreate("Theme",
		"*Id", "#Title")

	ControllerFileCreate("Paymant",
		"*Number", "#Money", "#Date", "#Time", "#Users", "#Statistics", 
		"#ProtectionCheating", "%#OrganizationId")

	ControllerFileCreate("DeadLine",
		"*Id", "#Date", "#Time", "*#OrganizationId")

	ControllerFileCreate("Session",
		"*Id", "*#UserId", "#IpAddress", "#BackendStartDateTime", 
		"#StateChangeDateTime", "#StateId")

	ControllerFileCreate("State",
		"*Id", "#Title")
}

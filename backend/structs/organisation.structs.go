package structs

import "time"

type Client struct {

	Id 					int 
    FirstName   		string
    LastName    		string
    Patronymic    		string
    Password   			string
    Email   			string
    Telephone   		string
    EmailNotifications 	bool  
    ManageCourses 		bool  
    ManageUsers 		bool  
    AploadFiles 		bool  
    ViewYourResults 	bool  
    ViewOtherResults 	bool  
    DepartmentId 		int 
    GroupId 			int 
    CreatorId 			int

}

type Course struct{
	Id 					int 
	Title    			string
	Text   	 			string
	Files 				bool   
    Date  				time.Time 
    Time  				time.Time
    DateDel 			time.Time
    TimeDel 			time.Time
    ClientId 			int 
}

type CourseResults struct{

	Id 			int 
    Assessment  string
    Scores 		int 
    NumberTests int 
    PassageTime time.Time
    Date 		time.Time 
    Time 		time.Time
    ClientId 	int 
    CourseId 	int 
}

type ActiveCourse struct{

	Id	 		int 
    Date 		time.Time 
    Time 		time.Time
    DateClose 	time.Time 
    TimeClose 	time.Time
    CourseId 	int 
    ClientId 	int 
	
}

type Department struct{
	Id 		int 
    Title  	string  
}
	
type Group struct{
	Id 				int 
    Title    		string
    TitleSingular   string  
    DepartmentId 	int 
}

type File struct{
	Id 			int 
    Date 		time.Time 
    Time 		time.Time 
    DateDel 	time.Time
    TimeDel 	time.Time
    FileName  	string 
    FileNameTmp string
}

type FileInformation struct{

	ClientId 	 int 
    FileId 		 int 
    PublicInfoId int
    TestId 		 int
    QuestionId 	 int
}

type PublicInfo struct{

	Id 		 int 
    Title    string 
    Text 	 string 
    Files 	 bool   
    Date 	 time.Time 
    Time 	 time.Time
    DateDel  time.Time
    TimeDel  time.Time
    ClientId int 
}

type Test struct{

	Id 		int 
    Title 	string   
    Text 	string  
    Files 	bool   
    Date 	time.Time 
    Time 	time.Time
    DateDel time.Time
    TimeDel time.Time
}

type TestResults struct{

	Id 			int 
    PassageTime time.Time
    Assessment  string
    Scores 		int 
    Attempts 	int 
    Date 		time.Time 
    Time 		time.Time
    CourseId 	int 
    TestId 		int 
    ClientId 	int 
}

type ActiveTest struct{

	Id 			int 
    Date 		time.Time 
    Time 		time.Time
    DateClose 	time.Time 
    TimeClose 	time.Time
    Attempts 	int 
    TestTypeId 	int
    TestId 		int 
    ClientId 	int 
}

type TestType struct{
	Id 		int 
    Title   int 
}
	
type Question struct{
	Id 				int 
    Text 			string 
    Date 			time.Time 
    Time 			time.Time
    DateDel 		time.Time
    TimeDel 		time.Time
    Hint  			string
    AnswerVariant 	string  //Варианты ответа будут храниться в таком виде - ответ;ответ;ответ;ответ
    AnswerCorrect 	string  //Верные варианты ответа будут храниться в таком виде - ответ;ответ;ответ;ответ
    TestId 			int 
    Files 			bool   
}
	
type QuestionResult struct{
    Id 				int 
	QuestionTypeId	int 
    Date 			time.Time 
    Time 			time.Time
    QuestonId 		int 
    ClientId 		int 
}
	
type QuestionType struct{
	Id 			int 
    Title 		string
    InputTypeId int
}

type InputType struct{
	Id 	  int 
	Title string
	Type  string
}
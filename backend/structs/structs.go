package structs

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Level struct {
	Id               int 	`json:"id"`
	Title            string `json:"title"`
	Price            int	`json:"price"`
	Paid             bool	`json:"paid"`
	CreateCourse     bool	`json:"createCourse"`
	TakeCourse       bool	`json:"takeCourse"`
	AploadFile       bool	`json:"aploadFile"`
	ViewYourResult   bool	`json:"viewYourResult"`
	ViewOtherResults bool	`json:"viewOtherResults"`
}
type ClientLevel struct {
    Id                      int 	`json:"id"`  
    Title                   string  `json:"title"`  
    CreateCourse            bool	`json:"createCourse"`  
    TakeCourse              bool	`json:"takeCourse"`  
    AploadFile              bool	`json:"aploadFile"`  
    ViewYourResult          bool	`json:"viewYourResult"`  
    ViewOtherResults        bool	`json:"viewOtherResults"`  
}

type Organization struct {
	Id                 int 		`json:"id"`
	Title              string 	`json:"title"`
	Password           string	`json:"password"`
	Email              string	`json:"email"`
	EmailNotifications bool		`json:"emailNotifications"`
	Key                string 	`json:"key"`
	LevelId            int		`json:"levelId"`
}

type DeadLine struct {
	Id             int 					`json:"id"`
	Date           timestamp.Timestamp	`json:"date"`
	LevelId        int					`json:"levelId"`
	OrganizationId int8					`json:"organizationId"`
}

type Payment struct {
	Number         int					`json:"number"`
	Name           string				`json:"name"`
	Money          int					`json:"money"`
	Date           timestamp.Timestamp	`json:"date"` 
	LevelId        int					`json:"levelId"`
	OrganizationId int8					`json:"organizationId"`
}

type Client struct {
    Id 						int 	`json:"id"`  
    FirstName               string	`json:"firstName"`   
    LastName                string	`json:"lastName"`   
    Patronymic              string	`json:"patronymic"`          
    Password                string	`json:"password"`   
    Email                   string	`json:"email"`   
    Telephone               string	`json:"telephone"`   
    EmailNotifications      bool	`json:"emailNotifications"`  
    GroupId                 int		`json:"groupId"`  
    ClientLevelId           int		`json:"clientLevelId"`  
}

type Test struct {
	Id      int 					`json:"id"`
	Title   string 					`json:"title"`
	Text    string					`json:"text"`
	Date    timestamp.Timestamp		`json:"date"`
	DateDel timestamp.Timestamp		`json:"dateDel"`
}

type Admin struct {
	Id        int 		`json:"id"`
	FirstName string	`json:"firstName"`
	LastName  string	`json:"lastName"`
	Password  string	`json:"password"`
	Email     string	`json:"email"`
}

type Question struct {
	Id             int 					`json:"id"`
	Text           string				`json:"text"`
	Date           timestamp.Timestamp	`json:"date"`
	DateDel        timestamp.Timestamp	`json:"dateDel"`
	TestId         int					`json:"testId"`
	QuestionTypeId int					`json:"questionTypeId"`
	Hint           string				`json:"hint"`
}

type QuestionResult struct {
	Id         int 					`json:"id"`
	Date       timestamp.Timestamp	`json:"date"`
	QuestionId int					`json:"questionId"`
	ClientId   int					`json:"clientId"`
	Scores     int					`json:"scores"`
}

type QuestionType struct {
	Id   int 	`json:"id"`
	Type int	`json:"type"`
}

type Answer struct {
	Id         int 		`json:"id"`
	Text       string	`json:"text"`
	Correct    bool		`json:"correct"`
	QuestionId int8		`json:"questionId"`
}

type Course struct {
	Id       int 					`json:"id"`
	Title    string 				`json:"title"`
	Text    string					`json:"text"`
	Date     timestamp.Timestamp	`json:"date"`
	DateDel  timestamp.Timestamp	`json:"dateDel"`
	ClientId int					`json:"clientId"`
}

type PublicInfo struct {
	Id       int 					`json:"id"`
	Title    string 				`json:"title"`
	Text     string					`json:"text"`
	Date     timestamp.Timestamp	`json:"date"`
	DateDel  timestamp.Timestamp	`json:"dateDel"`
	ClientId int					`json:"clientId"`
}

type CourseResults struct {
	Id         int 					`json:"id"`
	Time       timestamp.Timestamp	`json:"time"`
	Date       timestamp.Timestamp	`json:"date"`
	Assessment string				`json:"assessment"`
	Scores     float32				`json:"scores"`
	ClientId   int					`json:"clientId"`
	CourseId   int					`json:"courseId"`
}

type Department struct {
	Id    int 		`json:"id"`
	Title string 	`json:"title"`
}

type Group struct {
	Id           int 	`json:"id"`
	Title        string `json:"title"`
	DepartmentId int	`json:"departmentId"`
}

type File struct {
	Id          	int 				`json:"id"` 
	Date        	timestamp.Timestamp	`json:"date"` 	
	DateDel     	timestamp.Timestamp	`json:"dateDel"`  
	FileName    	string 				`json:"fileName"`
	FileNameTmp 	string 				`json:"fileNameTmp"`
	TestId      	int8				`json:"testId"` 	
	QuestionId  	int					`json:"questionId"` 
	PublicInfoId 	int 				`json:"publicInfoId"`
	ClientId    	int8				`json:"clientId"`
}

type ActiveTest struct {
	Id           int 					`json:"id"`
	Start        timestamp.Timestamp	`json:"start"`
	End          timestamp.Timestamp	`json:"end"`
	Time         timestamp.Timestamp	`json:"time"`
	Attempts     int					`json:"attempts"`
	TestId       int					`json:"testId"`
	ClientId     int					`json:"clientId"`
	TrainingTest bool					`json:"trainingTest"`
}

type TestResults struct {
	Id          int 				`json:"id"`
	Time        timestamp.Timestamp	`json:"time"`
	Date        timestamp.Timestamp	`json:"date"`
	ClientId    int					`json:"clientId"`
	TestId      int8				`json:"testId"`
	Assessment  string				`json:"assessment"`
	PassageTime timestamp.Timestamp	`json:"passageTime"`
	Scores      float32				`json:"scores"`
	CourseId    int8				`json:"courseId"`
}

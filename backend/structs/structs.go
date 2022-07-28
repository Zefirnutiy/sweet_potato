package structs

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

<<<<<<< HEAD
type ActiveTest struct {
	Id        int64
	Start     timestamp.Timestamp
	End       timestamp.Timestamp
	Time      timestamp.Timestamp
	Attempts  int8
	AccessKey string
	TestId    int
	ClientId  int64
}

type Admin struct {
	Id        int64
	LastName  string
	FirstName string
	Login     string
	Password  string
	Email     string
}

type Answer struct {
	Id         int64
	Text       string
	Correct    bool
	QuestionId int8
=======


type Level struct {
    Id                      int  
    Title                   string  
    Price                   int  
    Paid                    bool  
    CreateCourse            bool  
    TakeCourse              bool  
    AploadFile              bool  
    ViewYourResult          bool  
    ViewOtherResults        bool  
}

type Organization struct {
    Id                      int  
    Title                   string
    Password                string
    Email                   string
    EmailNotifications      bool  
    LevelId                 int
}

type DeadLine struct {
    Id                      int  
    Date                    timestamp.Timestamp
    LevelId                 int  
    OrganizationId          int8  
}

type Payment struct {
    Number                  int  
    Name                    string
    Money                   int  
    Date                    timestamp.Timestamp
    LevelId                 int  
    OrganizationId          int8  
>>>>>>> 495706b989217b6557a78f9451b7f69b60a1096c
}



type Client struct {
<<<<<<< HEAD
	Id                 int64
	FirstName          string
	LastName           string
	Patronymic         string
	Login              string
	Password           string
	Email              string
	Telephone          string
	EmailNotifications bool
	LevelId            int8
	GroupId            int
	OrganizationId     int8
}

type Course struct {
	Id       int64
	Tittle   string
	Date     timestamp.Timestamp
	DateDel  timestamp.Timestamp
	ClientId int64
}

type CourseResults struct {
	Id         int64
	Time       timestamp.Timestamp
	Date       timestamp.Timestamp
	Assessment string
	Scores     int8
	ClientId   int64
	CourseId   int64
}

type DeadLine struct {
	Id             int64
	Date           timestamp.Timestamp
	LevelId        int8
	OrganizationId int8
}

type Department struct {
	Id             int64
	Title          string
	OrganizationId int8
}

type File struct {
	Id          int64
	Date        timestamp.Timestamp
	DateDel     timestamp.Timestamp
	FileName    string
	FileNameTmp string
	TestId      int8
	QuestionId  int64
	ClientId    int8
}

type Group struct {
	Id           int64
	Title        string
	DepartmentId int
}

type Level struct {
	Id               int64
	Title            string
	Price            int
	Paid             bool
	CreateCourse     bool
	TakeCourse       bool
	AploadFile       bool
	ViewYourResult   bool
	ViewOtherResults bool
}

type Organization struct {
	Id                 int64
	Title              string
	Login              string
	Password           string
	Email              string
	Telephone          string
	EmailNotifications bool
	LevelId            int8
}

type Payment struct {
	Number   int
	Name     string
	Money    int
	Date     timestamp.Timestamp
	LevelId  int8
	ClientId int8
}

type Question struct {
	Id             int64
	Text           string
	Date           timestamp.Timestamp
	DateDel        timestamp.Timestamp
	TestId         int64
	QuestionTypeId int
}

type QuestionType struct {
	Id   int64
	Type int
}

type Test struct {
	Id      int64
	Title   string
	Text    string
	Date    timestamp.Timestamp
	DateDel timestamp.Timestamp
}

type TestResults struct {
	Id          int64
	Time        timestamp.Timestamp
	Date        timestamp.Timestamp
	ClientId    int64
	TestId      int8
	Assessment  string
	PassageTime timestamp.Timestamp
	Scores      int8
	CourseId    int8
}

type Publicinfo struct {
	Id       int64
	Title    string
	Text     string
	File     string
	Date     timestamp.Timestamp
	DateDel  timestamp.Timestamp
	ClientId int8
=======
    Id int  
    FirstName               string   
    LastName                string   
    Patronymic              string      
    Password                string   
    Email                   string   
    Telephone               string   
    EmailNotifications      bool  
    GroupId                 int  
}

type Test struct {
    Id                      int  
    Title                   string   
    Text                    string   
    Date                    timestamp.Timestamp
    DateDel                 timestamp.Timestamp
}

type Admin struct {
    Id                      int  
    FirstName               string   
    LastName                string   
    Password                string   
    Email                   string   
}

type Question struct {
    Id                      int  
    Text                    string  
    Date                    timestamp.Timestamp
    DateDel                 timestamp.Timestamp
    TestId                  int  
    QuestionTypeId          int  
    Hint                    string
}

type QuestionResult struct {
	Id                      int  
	QuestionId              int 
	ClientId                int 
	Scores                  int 
}

type QuestionType struct {
    Id                      int  
    Type                    int  
}

type Answer struct {
    Id                      int  
    Text                    string  
    Correct                 bool  
    QuestionId              int8  
}

type Course struct {
    Id                      int  
    Title                   string
    Date                    timestamp.Timestamp
    DateDel                 timestamp.Timestamp
    ClientId                int  
}

type PublicInfo struct {
    Id                      int  
    Title                   string 
    Text                    string
    File                    string
    Date                    timestamp.Timestamp
    DateDel                 timestamp.Timestamp
    ClientId                int  
}


type CourseResults struct {
    Id                      int  
    Time                    timestamp.Timestamp  
    Date                    timestamp.Timestamp  
    Assessment              string
    Scores                  float32  
    ClientId                int  
    CourseId                int  
}


type Department struct {
    Id                      int  
    Title                   string  
}

type Group struct {
    Id                      int  
    Title                   string 
    DepartmentId            int  
}

type File struct {
    Id                      int  
    Date                    timestamp.Timestamp
    DateDel                 timestamp.Timestamp
    FileName                string  
    FileNameTmp             string   
    TestId                  int8  
    QuestionId              int  
    ClientId                int8  
}

type ActiveTest struct {
    Id                      int  
    TimeStart               timestamp.Timestamp
    TimeEnd                 timestamp.Timestamp
    Time                    timestamp.Timestamp
    Attempts                int  
    TestId                  int  
    ClientId                int  
    Trainingtest            bool 
>>>>>>> 495706b989217b6557a78f9451b7f69b60a1096c
}

type TestResults struct {
    Id                      int  
    Time                    timestamp.Timestamp  
    Date                    timestamp.Timestamp  
    ClientId                int  
    TestId                  int8  
    Assessment              string
    PassageTime             timestamp.Timestamp  
    Scores                  float32  
    CourseId                int8  
}
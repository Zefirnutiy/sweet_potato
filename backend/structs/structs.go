package structs

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

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
}

type Client struct {
    Id 						int  
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
    Date                    timestamp.Timestamp
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
    Start               timestamp.Timestamp
    End                 timestamp.Timestamp
    Time                    timestamp.Timestamp
    Attempts                int  
    TestId                  int  
    ClientId                int  
    TrainingTest            bool 
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
package structs

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type ActiveTest struct {
	Id         int64
	Start      timestamp.Timestamp
	End        timestamp.Timestamp
	Time       timestamp.Timestamp
	Attempts   int8
	AccessKey  string
	TestId     int
	ClientId   int64
}

type Admin struct {
	Id         int64
	LastName   string
	FirstName  string
	Login      string
	Password   string
	Email      string
}

type Answer struct {
	Id          int64
	Text        string
	Correct     bool
	QuestionId  int8
}

type Client struct {
	Id                  int64
	FirstName           string
	LastName            string
	Patronymic          string
	Login               string
	Password            string
	Email               string
	Telephone           string
	EmailNotifications  bool
	LevelId             int8
	GroupId             int
	OrganizationId      int8
}

type Course struct {
	Id        int64
	Tittle    string
	Date      timestamp.Timestamp
	DateDel   timestamp.Timestamp
	ClientId  int64
}

type CourseResults struct {
	Id         int64
	Time       timestamp.Timestamp
	Date       timestamp.Timestamp
	Assessment string
	Scores     int8
	ClientId  int64
	CourseId  int64
}

type DeadLine struct {
	Id              int64
	Date            timestamp.Timestamp
	LevelId         int8
	OrganizationId  int8
}

type Department struct {
	Id              int64
	Title           string
	OrganizationId  int8
}

type File struct {
	Id            int64
	Date          timestamp.Timestamp
	DateDel       timestamp.Timestamp
	FileName      string
	FileNameTmp   string
	TestId        int8
	QuestionId    int64
	ClientId      int8
}

type Group struct {
	Id            int64
	Title         string
	DepartmentId  int
}

type Level struct {
	Id                 int64
	Title              string
	Price              int
	Paid               bool
	CreateCourse       bool
	TakeCourse         bool
	AploadFile         bool
	ViewYourResult     bool
	ViewOtherResults   bool
}

type Organization struct {
	Id                  int64
	Title               int
	Login               string
	Password            string
	Email               string
	Telephone           string
	EmailNotifications  bool
	LevelId             int8
}

type Payment struct {
	Number    int
	Name      string
	Money     int
	Date      timestamp.Timestamp
	LevelId   int8
	ClientId  int8
}

type Question struct {
	Id               int64
	Text             string
	Date             timestamp.Timestamp
	DateDel          timestamp.Timestamp
	TestId           int64
	QuestionTypeId   int
}

type QuestionType struct {
	Id   	int64
	Type 	int
}

type QuestionResult struct {
	Id 		int
	QuestionId 	int
	ClientId 	int
	Scores 		int
}

type Test struct {
	Id       int64
	Title    string
	Text     string
	Date     timestamp.Timestamp
	DateDel  timestamp.Timestamp
}

type TestResults struct {
	Id           int64
	Time         timestamp.Timestamp
	Date         timestamp.Timestamp
	ClientId     int64
	TestId       int8
	Assessment   string
	PassageTime  timestamp.Timestamp
	Scores       int8
	CourseId     int8
}

type Publicinfo struct {
    	Id 		int64
    	Title 		string
    	Text 		string
    	File 		string
    	Date 		timestamp.Timestamp
    	DateDel 	timestamp.Timestamp
    	ClientId 	int8
}


package structs

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Active_Test struct {
	Id         int64
	Start      timestamp.Timestamp
	End        timestamp.Timestamp
	Time       timestamp.Timestamp
	Attempts   int8
	Access_Key string
	Test_id    int
	Client_Id  int64
}

type Admin struct {
	Id         int64
	Last_Name  string
	First_Name string
	Login      string
	Password   string
	Email      string
}

type Answer struct {
	Id          int64
	Text        string
	Correct     bool
	Question_Id int8
}

type Client struct {
	Id                  int64
	First_Name          string
	Last_Name           string
	Patronymic          string
	Login               string
	Password            string
	Email               string
	Telephone           string
	Email_Notifications bool
	Level_Id            int8
	Group_Id            int
	Organization_Id     int8
}

type Course struct {
	Id        int64
	Tittle    string
	Date      timestamp.Timestamp
	Date_Del  timestamp.Timestamp
	Client_Id int64
}

type Course_results struct {
	Id         int64
	Time       timestamp.Timestamp
	Date       timestamp.Timestamp
	Assessment string
	Scores     int8
	Client_Id  int64
	Course_Id  int64
}

type Dead_Line struct {
	Id              int64
	Date            timestamp.Timestamp
	Level_Id        int8
	Organization_Id int8
}

type Department struct {
	Id              int64
	Title           string
	Organization_Id int8
}

type File struct {
	Id            int64
	Date          timestamp.Timestamp
	Date_Del      timestamp.Timestamp
	File_Name     string
	File_Name_Tmp string
	Test_Id       int8
	Question_Id   int64
	Client_Id     int8
}

type Group struct {
	Id            int64
	Title         string
	Department_Id int
}

type Level struct {
	Id                 int64
	Title              string
	Price              int
	Paid               bool
	Create_Course      bool
	Take_Course        bool
	Apload_File        bool
	View_Your_Result   bool
	View_Other_Results bool
}

type Organization struct {
	Id                  int64
	Title               int
	Login               string
	Password            string
	Email               string
	Telephone           string
	Email_Notifications bool
	Level_Id            int8
}

type Payment struct {
	Number    int
	Name      string
	Money     int
	Date      timestamp.Timestamp
	Level_Id  int8
	Client_Id int8
}

type Question struct {
	Id               int64
	Text             string
	Date             timestamp.Timestamp
	Date_Del         timestamp.Timestamp
	Test_Id          int64
	Question_Type_Id int
}

type Question_Type struct {
	Id   int64
	Type int
}

type Session struct {
	Id        int64
	Date      int
	Token     int
	Client_Id int8
}

type Test struct {
	Id       int64
	Title    string
	Text     string
	Date     timestamp.Timestamp
	Date_Del timestamp.Timestamp
}

type Test_Results struct {
	Id           int64
	Time         timestamp.Timestamp
	Date         timestamp.Timestamp
	Client_Id    int64
	Test_Id      int8
	Assessment   string
	Passage_Time timestamp.Timestamp
	Scores       int8
	Course_Id    int8
}

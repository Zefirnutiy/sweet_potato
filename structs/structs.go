package structs

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Student struct {
	Id         int    `json:"Id"`
	First_name string `json:"First_name"`
	Sure_name  string `json:"Sure_name"`
	Email      string `json:"Email"`
	Password   string `json:"Password"`
	Major      string `json:"Major"`
	Course     int    `json:"Course"`
	Group      int    `json:"Group"`
}

type Question struct {
	Id      int    `json:"Id"`
	Test_id int    `json:"Test_id"`
	Text    string `json:"Text"`
	Answer  int    `json:"Answer"`
	Files   int    `json:"Files"`
	Scores  int    `json:"Scores"`
}

type Question_result struct {
	Id          int  `json:"Id"`
	Student_id  int  `json:"Student_id"`
	Question_id int  `json:"Question_id"`
	Accuracy    bool `json:"Accuracy"`
}

type Answer_option struct {
	Id          int    `json:"Id"`
	Question_id int    `json:"Question_id"`
	Value       string `json:"Value"`
	True_answer bool   `json:"True_answer"`
}

type Question_file struct {
	Id          int    `json:"Id"`
	Question_id int    `json:"Question_id"`
	File        string `json:"File"`
}

type Test struct {
	Id           int                 `json:"Id"`
	Professor_id int                 `json:"Professor_id"`
	Title        string              `json:"Title"`
	Start        timestamp.Timestamp `json:"Start"`
	End          timestamp.Timestamp `json:"End"`
	Question_id  int                 `json:"Question_id"`
	Attempt      int                 `json:"Attempt"`
	Time         timestamp.Timestamp `json:"Time"`
}

type Result struct {
	Id           int                 `json:"Id"`
	Student_id   int                 `json:"Student_id"`
	Test_id      int                 `json:"Test_id"`
	Attempt      int                 `json:"Attempt"`
	Time_spent   timestamp.Timestamp `json:"Time_spent"`
	Score_gained int                 `json:"Score_gained"`
}

type Grade struct {
	Id         int `json:"Id"`
	Test_id    int `json:"Test_id"`
	Student_id int `json:"Student_id"`
}

type Additional_materials struct {
	Id           int    `json:"Id"`
	Professor_id int    `json:"Professor_id"`
	File         string `json:"File"`
}

type Professor struct {
	Id         int    `json:"Id"`
	First_name string `json:"First_name"`
	Sure_name  string `json:"Sure_name"`
	Patronymic string `json:"Patronymic"`
	Email      string `json:"Email"`
	Password   string `json:"Password"`
}

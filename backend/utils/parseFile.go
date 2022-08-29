package utils

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var globalId = 0
var globalStr = "%^$"
type InputButton struct{
	Id 				int
	Question 		string
	Answers 		[]string
	CorrectAnswers  []string
	Hint 			string
	Type 			string
}

type InputText struct {
	Id			int
	Question	string
	Answer		string
}

type InpuntConstraint struct{
	Id 			int
	Question	string
	Answers		map[string]string
}

func SortQuestion(path string){
	allQuestion := strings.Split(parseFile(path), "\n\n")

	for _, value := range allQuestion{
		typeAns := getType(value)
		switch typeAns{
		
		case "*":
			fmt.Println(oneChoise(value, typeAns))
		case "#":
			fmt.Println(oneChoise(value, typeAns))
		}
		
	}
}
func oneChoise(str, simbol string) InputButton {
	trueAns := ""
	otherAns := ""
	hint := ""
	strMas := strings.Split(str, "\n")
	for index, val := range strMas{
		if index != 0 {
			if strings.Contains(val, simbol){
				val = strings.ReplaceAll(val, simbol, "")
				trueAns += val 
			}else if strings.Contains(val, "~"){
				val = strings.ReplaceAll(val, "~", "")
				hint = val
			}else{
				otherAns += val + globalStr
			}
		}
	}
	globalId += 1
	return InputButton{globalId, strMas[0], strings.Split(otherAns, globalStr), strings.Split(trueAns, globalStr), hint, simbol}
}


func getType(str string) string{
	
	if strings.Contains(str, "*") {
		return "*"
	}

	if strings.Contains(str, "#") {
		return "#"
	}

	if strings.Contains(str, "[") && strings.Contains(str, "]") {
		return "[]"
	}

	if strings.Contains(str, "=") {
		return "="
	}

	if strings.Contains(str, "1. ") {
		return "1. "
	}

	return "empty"
}



func parseFile(path string) string{
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)
	dataString := ""
	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		dataString += (string(data[:n]))
	}
	return dataString
}


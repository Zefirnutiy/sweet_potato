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
	Type		string
}

type InpuntConstraint struct{
	Id 				int
	Question		string
	LeftAnswers		[]string
	RightAnswers	[]string
	Type			string
}

type InputSort struct {
	Id		int
	Question string
	Answers 	map[string]int
}

//функция принимает в себя строку - путь к файлу, который должен соответствовать стандартам
//которые мы опишем позже.
func SortQuestion(path string){
	allQuestion := strings.Split(parseFile(path), "\n\n")

	for _, value := range allQuestion{
		typeAns := getType(value)
		switch typeAns{
		
		case "*":
			fmt.Println(inputButton(value, typeAns))
		case "#":
			fmt.Println(inputButton(value, typeAns))
		case "[]":
			fmt.Println(inputText(value))
		case "=":
			fmt.Println(inputConstraint(value))
		case "empty":
			fmt.Println(inputSort(value))
		}
		
	}
}

func inputSort(str string) InputSort {
	mapAns := make(map[string]int)
	strMas := strings.Split(str, "\n")
	for index, val := range strMas[1:] {
		mapAns[val] = index + 1
	}
	globalId += 1

	return InputSort{globalId, strMas[0], mapAns}
}

func inputConstraint(str string) InpuntConstraint {
	strMas := strings.Split(str, "\n")
	leftAns := ""
	rightAns := ""
	for _, val := range strMas[1:]{
		mas := strings.Split(val, "=")

		leftAns += mas[0] + globalStr
		rightAns += mas[len(mas)-1] + globalStr
	}
	globalId += 1
	return InpuntConstraint{globalId, strMas[0], strings.Split(leftAns, globalStr), strings.Split(rightAns, globalStr), "="}
}

func inputButton(str, simbol string) InputButton {
	trueAns := ""
	otherAns := ""
	hint := ""
	strMas := strings.Split(str, "\n")
	for _, val := range strMas[1:]{
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
	globalId += 1
	return InputButton{globalId, strMas[0], strings.Split(otherAns, globalStr), strings.Split(trueAns, globalStr), hint, simbol}
}

func inputText(str string) InputText{
	mas := strings.Split(str, "\n")
	clearStr := strings.ReplaceAll(mas[1], "[", "")
	clearStr = strings.ReplaceAll(clearStr, "]", "")
	globalId += 1
	return InputText{globalId, mas[0], clearStr, "[]"}
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


package utils

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type OneChoise struct{
	Question string
	Answers []string
	CorrectAnswers []string
}

func SortQuestion(path string){
	allQuestion := strings.Split(parseFile(path), "\n\n")


	for _, value := range allQuestion {
		getAnswers(value, "#")
		break
	}
}

func getAnswers(str string, simbol string) {
	answers := ""
	correctAnswer := ""
	strMas := strings.Split(str, "\n")
	for index, val := range strMas {
		if index != 0{

			if strings.Contains(val, simbol){
				correctAnswer += strings.ReplaceAll(val, simbol, "") + "^&"
				}else{
					answers += strings.ReplaceAll(val, simbol, "")
				}
			}
	}
	
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


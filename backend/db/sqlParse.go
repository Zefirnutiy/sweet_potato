package db

import (
	"fmt"
	"io"
	"os"
)

func DbCreate(path, schemaName string) string {
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

	text := fmt.Sprintf(`set schema '%[1]s';

%[2]s`, schemaName, dataString)

	return text

}

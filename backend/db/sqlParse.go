package db

import (
	"fmt"
	"io"
	"os"
)

func DbCreate(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)
	data_string := ""
	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		data_string += (string(data[:n]))
	}
	return data_string

}

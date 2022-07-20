package utils

import (
	"fmt"
)

func HandlerError(err error){
	if err != nil{
		Logger.Println(err)
		fmt.Println(err)
		return
	}
}
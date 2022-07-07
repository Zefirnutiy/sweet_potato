package functions

import "fmt"

func HandlerError(err error){
	if err != nil{
		fmt.Println(err)
		return
	}
}
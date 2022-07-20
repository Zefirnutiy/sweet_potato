package functions

import (
	"fmt"

	"github.com/Zefirnutiy/sweet_potato.git/utils"
)

func HandlerError(err error){
	if err != nil{
		utils.Logger.Println(err)
		fmt.Println(err)
		return
	}
}
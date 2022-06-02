package main

import (
	"fmt"
	"github.com/Zefirnutiy/sweet_potato.git/db"
	_ "github.com/lib/pq"
)

var cfg = db.Load("./settings.cfg")

func init() {
	err := db.Connect(*cfg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	
}

func main() {
	err := db.CreateTable()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	App(cfg.Port)

}

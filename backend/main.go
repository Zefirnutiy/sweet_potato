package main

import (
	"fmt"

	// "github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/routes"

	_ "github.com/lib/pq"
)

var cfg = db.Load("./settings.cfg")

func init() {
	err := db.Connect(*cfg)
	if err != nil {
		panic(err)
	}

}

func main() {
<<<<<<< HEAD
	err := db.CreateTable()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
=======
	
	// err := db.CreateTable("./db/main.sql", "Main")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// err = db.CreateTable("./db/organisation.sql", "KTK")
	
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// utils.GetRequestGenerate("DeadLine", "*#Id", "#Date", "#LevelId", "#OrganizationId")
>>>>>>> 495706b989217b6557a78f9451b7f69b60a1096c
	routes.Routs(cfg.Port)

}

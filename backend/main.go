package main

import (
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

	// err := db.CreateTable("./db/main.sql", "main")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// err = db.CreateTable("./db/organization.sql", "KTK")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// utils.Generate()
	// utils.ControllerFileCreate("Client",
	// 	"*Id", "#FirstName", "#LastName", "#Patronymic", "$#Password", "#Email",
	// 	"#Telephone", "#EmailNotifications", "*#GroupId", "*#ClientLevelId")

	routes.Routs(cfg.Port)

}

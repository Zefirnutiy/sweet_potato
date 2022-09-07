package main

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/routes"
	_ "github.com/lib/pq"
)

var cfg = db.Load("./settings.cfg")
// @title Test System
// @version 0.1
// @description Very cool system for testion of people

// @host localhost:8080
// @BasePath /

func init() {
	err := db.Connect(*cfg)
	if err != nil {
		panic(err)
	}
	
}

func main() {


	// err := db.CreateTable("./db/public.sql", "public")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// err = db.CreateTable("./db/organization.sql", "KTK")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// utils.Generate() // создание контроллеров
	// utils.SortQuestion("./text.txt") //Функция работы парсера
	routes.Routs(cfg.Port)

}

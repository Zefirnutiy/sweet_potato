package main

import (
	"fmt"

	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/routes"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func init() {
	if err := InitConfig(); err != nil{
		utils.Logger.Println(err)
		return
	}
	err := db.Connect(viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.user"), viper.GetString("db.name"))
	if err != nil {
		fmt.Println(err)
	}
	
}

func main() {
	// utils.SortQuestion("./text.txt")
	routes.Routs(viper.GetString("port"))
}


func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}


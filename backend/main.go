package main

import (
	"github.com/Zefirnutiy/sweet_potato.git/db"
	"github.com/Zefirnutiy/sweet_potato.git/routes"
	"github.com/Zefirnutiy/sweet_potato.git/utils"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func init() {
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	
}

func main() {

	if err := InitConfig(); err != nil{
		utils.Logger.Println(err)
		return
	}

	routes.Routs(viper.GetString("port"))

}


func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var Dbpool *sql.DB

func Connect() error {
	var err error

	if err := godotenv.Load(); err != nil{
		utils.Logger.Println(err)
		return err
	}
	
	Dbpool, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
	viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.user"), os.Getenv("DB_PASSWORD"), viper.GetString("db.name")))
	if err != nil {
		return err
	}

	return nil
}

//путь к sql файлу указывается относительно корня
//for example: ./db/main.sql
func CreateTable(filePath, schemaName string) error {

	_, err := Dbpool.Exec(fmt.Sprintf(`CREATE SCHEMA "%s";`, schemaName))

	if err != nil {
		fmt.Println("Ошибка создания Схемы:", err)
		return err
	}
	_, err = Dbpool.Exec(DbCreate(filePath, schemaName))

	if err != nil {
		fmt.Println("Ошибка создания таблиц:", err)
		return err
	}

	return nil
}

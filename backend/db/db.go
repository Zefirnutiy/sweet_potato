package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Zefirnutiy/sweet_potato.git/utils"
	"github.com/joho/godotenv"
)

var Dbpool *sql.DB

func Connect(host, port, user, dbName string) error {
	var err error

	if err := godotenv.Load(); err != nil{
		utils.Logger.Println(err)
		return err
	}
	Dbpool, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
	host, port, user, os.Getenv("DB_PASSWORD"), dbName))
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

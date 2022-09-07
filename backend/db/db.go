package db

import (
	"database/sql"
	"fmt"
)

var Dbpool *sql.DB

func Connect(cfg SettingServer) error {
	var err error
	Dbpool, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
	cfg.Database.DbHost, cfg.Database.DbPort, cfg.Database.DbUser, cfg.Database.DbPass, cfg.Database.DbName))
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

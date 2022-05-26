package db

import (
	"database/sql"
	"fmt"
)

var dbpool *sql.DB

func Connect(cfg SettingServer) error {
	var err error
	dbpool, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Database.DbHost, cfg.Database.DbPort, cfg.Database.DbUser, cfg.Database.DbPass, cfg.Database.DbName))
	if err != nil {
		return err
	}

	_, err = dbpool.Exec("SELECT id from Test")

	if err != nil {
		_, err = dbpool.Exec(fmt.Sprintf(`%s`, DbCreate("./db/database.sql")))
	}

	return nil
}

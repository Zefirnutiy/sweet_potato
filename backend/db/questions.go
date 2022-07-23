package db

import (
	"database/sql"
)

func QuestionRegister(title, password, email string, emailnotifications bool) (sql.Result, error) {

	result, err := Dbpool.Exec(`INSERT INTO main."Organization" (title, password, email, emailnotifications) VALUES ($1, $2, $3, $4);`, title, password, email, emailnotifications)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func SelectOrganization(email string) bool {

	err := Dbpool.QueryRow(`SELECT * FROM main."Organization" WHERE email IN ($1)`, email).Scan()

	if err != nil && err == sql.ErrNoRows {
		return false
	}

	return true
}

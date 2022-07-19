package db

import "github.com/golang/protobuf/ptypes/timestamp"

func QuestionRegister(title, password, email string, email_notification bool) {

	Dbpool.Exec(`INSERT INTO public."Service" (title, password, email, email_notification) VALUES ($1, $2, $3, $4);`, title, password, email, email_notification)
}

func QuestionSession(title, date timestamp.Timestamp, password, email string, email_notification bool) {

	Dbpool.Exec(`INSERT INTO public."Service" (title, password, email, email_notification) VALUES ($1, $2, $3, $4);`, password, email, email_notification)
}

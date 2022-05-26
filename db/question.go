package db

func InsertQuest() {

	dbpool.Exec(`INSERT INTO public."Question" (title,answer) VALUES ('где плавают негры?','Черное море');`)

}

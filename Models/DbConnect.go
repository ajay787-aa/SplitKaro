package Models

import (
	"github.com/jmoiron/sqlx"
	"log"
)

func dbConnect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres sslmode=disable password=root host=localhost")
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully connected")
	}
	return db, err
}

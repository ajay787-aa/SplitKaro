package Models

import (
	"fmt"
	"log"
)

type User struct {
	Name   string `db:"username"`
	Email  string `db:"email"`
	UserId int    `db:"userid"`
}

func (user User) ReadAllUsers() []User {
	db, _ := dbConnect()
	var place []User

	rows, err := db.Queryx("SELECT username, email,userid FROM users")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		if err := rows.StructScan(&user); err != nil {
			log.Fatalln(err)
		}
		place = append(place, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatalln(err)
	}

	return place
}

func (user User) InsertUser(user2 User) {
	db, _ := dbConnect()

	stmt, err := db.Preparex("INSERT INTO users (username, email) VALUES ($1, $2)")
	if err != nil {
		panic(err)
	}
	stmt.Exec(user2.Name, user2.Email)

}

func (user User) ReadUserByEmail(email string) (User, error) {
	db, _ := dbConnect()
	var place User
	query := "SELECT username, email,userid FROM users where email=$1"
	rows, err := db.Queryx(query, email)
	if err != nil {
		panic(err)
	}
	if rows.Next() {
		// Scan the result into the 'place' variable
		if err := rows.StructScan(&place); err != nil {
			return place, err
		}
	} else {
		return place, fmt.Errorf("no user found with email: %s", email)
	}

	return place, nil
}

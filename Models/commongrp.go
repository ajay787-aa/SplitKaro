package Models

import (
	"fmt"
	"log"
)

type Group struct {
	Id    int    `db:"grp_id"`
	Name  string `db:"group_name"json:"Name"`
	User  []User
	Email []string `json:"Email"`
	Uid   int      `db:"uid"`
}

func (group Group) CreateGroup(group2 Group) error {
	db, _ := dbConnect()
	user := group2.User

	stmt, err := db.Preparex("INSERT INTO commongrp (group_name, uid) VALUES ($1,$2)")
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(user); i++ {
		fmt.Println(group2.Name, user[i].UserId)
		_, err := stmt.Exec(group2.Name, user[i].UserId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (group Group) FindByGrpName(grpName string) ([]Group, error) {

	db, _ := dbConnect()
	var groups []Group
	query := "SELECT * FROM commongrp where group_name=$1"
	rows, err := db.Queryx(query, grpName)
	if err != nil {
		return groups, err
	}
	for rows.Next() {
		var group Group
		if err := rows.StructScan(&group); err != nil {
			log.Fatalln(err)
		}
		groups = append(groups, group)
	}
	if err == nil {
		return groups, nil
	}
	return nil, err
}

package Models

import (
	"fmt"
)

type Balance struct {
	GrpName      string `json:"UserGroup"`
	AdminEmailId string `json:"UserEmail"`
	Amount       int    `json:"Amount"`
	Sum          int    `db:"sum"`
	EmailB       int    `db:"email_b"`
	EmailA       int    `db:"email_a"`
}

func (Balance Balance) AddBalance(balance2 Balance) error {

	db, _ := dbConnect()
	groups, err := Group{}.FindByGrpName(balance2.GrpName)
	user, err := User{}.ReadUserByEmail(balance2.AdminEmailId)
	if err != nil {
		panic(err)
	}
	fmt.Println("Admin iD", user.UserId)
	balance2.Amount = balance2.Amount / len(groups)
	stmt, err := db.Preparex("INSERT INTO balance (uida,uidb,amount,grp_id) VALUES ($1,$2,$3,$4)")
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(groups); i++ {
		if groups[i].Uid != user.UserId {
			_, err := stmt.Exec(user.UserId, groups[i].Uid, balance2.Amount, groups[i].Id)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

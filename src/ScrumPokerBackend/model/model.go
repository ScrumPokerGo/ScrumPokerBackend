package model

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Account struct {
	ID        uint8     `db:"id" bson:"id"`
	Status    bool      `db:"status" bson:"status"`
	Name      string    `db:"name" bson:"name"`
	CreatedAt time.Time `db:"created_at" bson:"created_at"`
	UpdatedAt time.Time `db:"updated_at" bson:"updated_at"`
	Deleted   bool      `db:"deleted" bson:"deleted"`
}

func NewAccount(name string) Account {
	ac := Account{}
	ac.ID = 0
	ac.Name = name
	return ac
}

func (a Account) save() {
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	query_new := `INSERT INTO account (status,name,deleted) values(true,:name,false)`
	query_update := ``

	if a.ID == 0 {
		_, err := db.NamedExec(query_new, a)
	} else {
		_, err := db.NamedExec(query_update, a)
	}

	if err != nil {
		fmt.Println(err)
		return

	}

	//save the model to the database
}

func (a Account) delete() {
	//delete
}

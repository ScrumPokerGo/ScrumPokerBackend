package model

import (
	"time"
	"ScrumPokerBackend/database"
)

type User struct{
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Deleted   bool      `json:"deleted"`
	Account   Account	`json:"account"`
}


func NewUser(name string) User {
	user := User{}
	user.ID = 0
	user.Name = name
	return user
}

func (u User) Save() {
	if (u.ID==0){
		database.DB.Create(&u)
	}else{
		database.DB.Save(&u)
	}
}

func GetUser(id uint) User{
	u := User{}
	database.DB.First(&u, id)
	return u
}

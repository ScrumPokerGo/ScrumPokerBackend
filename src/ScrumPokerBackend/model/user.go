package model

import (
	"time"
	"ScrumPokerBackend/database"
	"ScrumPokerBackend/error"
)

type User struct{
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Deleted   bool      `json:"deleted"`
	Account   Account	`json:"account"`
}


func NewUser(name string) *User {
	user := User{}
	user.Name = name
	return &user
}

func (u *User) Save() error {
	db := database.DB.Save(&u)
	if (db.Error != nil){
		return db.Error
	}
	return nil
}

func (u *User) Create() error{
	db := database.DB.Create(&u)
	if (db.Error != nil){
		return db.Error
	}
	return nil
}


func GetUser(id uint) (*User,error){
	u := User{}
	db := database.DB.First(&u, id)
	if (db.RowsAffected==0){
		return nil,errors.New("User not found",404)
	}
	return &u,nil
}

func GetUserAll() ([]User, error){
	user_list := []User{}
	db := database.DB.Find(&user_list)
	if (db.RowsAffected==0){
		return nil,errors.New("Users not found",404)
	}
	return user_list,nil
}
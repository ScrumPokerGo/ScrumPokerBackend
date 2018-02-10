package model

import (
	"time"
	"ScrumPokerBackend/database"
)

type Account struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Status    bool      `json:"status"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Deleted   bool      `json:"deleted"`
}

func NewAccount(name string) Account {
	ac := Account{}
	ac.ID = 0
	ac.Name = name
	return ac
}

func GetAccount(id uint) Account{
	a := Account{}
	database.DB.First(&a, id)
	return a
}

func GetAccountAll() []Account{
	account_list := []Account{}
	database.DB.Find(&account_list)
	return account_list
}

func (a Account) Save() {
	if (a.ID==0){
		database.DB.Create(&a)
	}else{
		database.DB.Save(&a)
	}
}

func (a Account) Delete() {
	database.DB.Delete(&a)
}
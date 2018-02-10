package database

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func Init(){
	var err error = nil
	DB, err = gorm.Open("sqlite3", "scrumpoker.db")
	DB.LogMode(true)
	if err != nil {
		panic("failed to connect database")
	}
}


func Close(){
	DB.Close()

}
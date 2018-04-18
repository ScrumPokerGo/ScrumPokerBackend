package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"ScrumPokerBackend/utils"
)


var DB *gorm.DB

func Init() {
	defer utils.Log.Println("Database connection established")
	var err error = nil
	DB, err = gorm.Open("sqlite3", "scrumpoker.db")
	DB.LogMode(true)
	if err != nil {
		panic("failed to connect database")
	}

}

func Close() {
	DB.Close()

}

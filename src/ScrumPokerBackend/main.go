package main

import (
	"ScrumPokerBackend/database"
	"ScrumPokerBackend/model"
	"ScrumPokerBackend/routers"
	"ScrumPokerBackend/configuration"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
	"fmt"

)

func main() {
	configuration.Init()
	database.Init()

	defer database.Close()
	// Migrate the schema
	database.DB.AutoMigrate(&model.Account{})
	database.DB.AutoMigrate(&model.User{})
	database.DB.AutoMigrate(&model.Project{})
	database.DB.AutoMigrate(&model.Milestone{})
	database.DB.AutoMigrate(&model.Issue{})

	//account := model.NewAccount("test")
	//account.Save()
	account,error := model.GetAccount(1)
	fmt.Println(account.Name)
	if error!=nil{
		fmt.Println(account)
	}
	account.GitlabImport(55)

	//user := model.GetUser(1)
	//fmt.Println(user.Name)

	router := mux.NewRouter()
	routers.AddAccountRouters(router)
	routers.AddUserRouters(router)

	// Satisfies the http.Handler interface

	log.Fatal(http.ListenAndServe(":8000", router))
}

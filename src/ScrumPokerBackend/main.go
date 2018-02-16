package main

import (
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"ScrumPokerBackend/database"
	"ScrumPokerBackend/routers"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {

	database.Init()
	defer database.Close()
	// Migrate the schema
	//database.DB.AutoMigrate(&model.Account{})
	//database.DB.AutoMigrate(&model.User{})

	//account := model.GetAccount(1)
	//fmt.Println(account.Name)

	//user := model.GetUser(1)
	//fmt.Println(user.Name)

	router := mux.NewRouter()
	routers.AddAccountRouters(router)



	// Satisfies the http.Handler interface


	log.Fatal(http.ListenAndServe(":8000", handlers.RecoveryHandler()(router)))
}

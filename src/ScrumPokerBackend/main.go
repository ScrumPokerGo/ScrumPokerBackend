package main

import (
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"ScrumPokerBackend/model"
	"fmt"
	"github.com/gorilla/mux"
	"ScrumPokerBackend/handlers"
	"log"
	"ScrumPokerBackend/database"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {


	database.Init()
	// Migrate the schema
	//database.DB.AutoMigrate(&model.Account{})
	//database.DB.AutoMigrate(&model.User{})

	account := model.GetAccount(1)
	fmt.Println(account.Name)

	user := model.GetUser(1)
	fmt.Println(user.Name)


	 router := mux.NewRouter()
	// Routes consist of a path and a handler function.
	 router.HandleFunc("/account", handlers.GetAccounts).Methods("GET")
	 router.HandleFunc("/account/{id}", handlers.GetAccount).Methods("GET")
	 router.HandleFunc("/account/", handlers.CreateAccount).Methods("POST")


	// Bind to a port and pass our router in
	 log.Fatal(http.ListenAndServe(":8000", router))
}

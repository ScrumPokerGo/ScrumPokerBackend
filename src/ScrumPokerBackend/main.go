package main

import (
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/gorilla/mux"
	"log"
	"ScrumPokerBackend/database"
	"ScrumPokerBackend/routers"

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


	 log.Fatal(http.ListenAndServe(":8000", router))
/*
	type StatusError struct {
		Status int
		// Allows you to wrap another error
		Err error
	}

	func (e *StatusError) Error() string {
		return e.Error()
	}

	type AppHandler func(w http.ResponseWriter, r *http.Request) error

	// Satisfies the http.Handler interface
	func (ah AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Centralises your error handling
	err := ah(w, r)
	if err != nil {
	switch e := a.(type) {
	case *StatusError:
	switch e.Status {
	case 400:
	http.Error(w, e.Err.Error(), 400)
	return
	case 404:
	http.NotFound(w, r)
	return
	default:
	http.Error(w, http.StatusText(500), 500)
	return
	}
	default:
	http.Error(w, http.StatusText(500), 500)
	return
	}
	}
*/
	}

package routers

import (
	"ScrumPokerBackend/api"
	"ScrumPokerBackend/middleware"
	"github.com/gorilla/mux"
)

func AddAccountRouters(r *mux.Router) {

	r.HandleFunc("/account", middleware.MiddlewareSet(api.GetAccounts)).Methods("GET")
	r.HandleFunc("/account/{id}", middleware.MiddlewareSet(api.GetAccount)).Methods("GET")
	r.HandleFunc("/account/", middleware.MiddlewareSet(api.CreateAccount)).Methods("POST")

}

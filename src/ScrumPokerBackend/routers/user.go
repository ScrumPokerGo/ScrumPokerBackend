package routers

import (
	"ScrumPokerBackend/api"
	"ScrumPokerBackend/middleware"
	"github.com/gorilla/mux"
)

func AddUserRouters(r *mux.Router) {

	r.HandleFunc("/user", middleware.MiddlewareSet(api.GetUsers)).Methods("GET")
	r.HandleFunc("/user/{id}", middleware.MiddlewareSet(api.GetUser)).Methods("GET")
	r.HandleFunc("/user/", middleware.MiddlewareSet(api.CreateUser)).Methods("POST")
}

package routers


import(
	"github.com/gorilla/mux"
	"ScrumPokerBackend/handlers"
	"ScrumPokerBackend/middleware"
)


func AddAccountRouters(r *mux.Router){

	r.HandleFunc("/account", middleware.MiddlewareSet(handlers.GetAccounts)).Methods("GET")
	r.HandleFunc("/account/{id}", middleware.MiddlewareSet(handlers.GetAccount)).Methods("GET")
	r.HandleFunc("/account/", middleware.MiddlewareSet(handlers.CreateAccount)).Methods("POST")
}
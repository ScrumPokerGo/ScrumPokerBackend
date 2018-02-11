package routers


import(
	"github.com/gorilla/mux"
	"ScrumPokerBackend/handlers"
)

func AddAccountRouters(r *mux.Router){
	r.HandleFunc("/account", handlers.GetAccounts).Methods("GET")
	r.HandleFunc("/account/{id}", handlers.GetAccount).Methods("GET")
	r.HandleFunc("/account/", handlers.CreateAccount).Methods("POST")
}
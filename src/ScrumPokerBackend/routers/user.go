package routers


import(
	"github.com/gorilla/mux"
	"ScrumPokerBackend/handlers"
	"ScrumPokerBackend/middleware"
)


func AddUserRouters(r *mux.Router){

	r.HandleFunc("/user", middleware.MiddlewareSet(handlers.GetUsers)).Methods("GET")
	r.HandleFunc("/user/{id}", middleware.MiddlewareSet(handlers.GetUser)).Methods("GET")
	r.HandleFunc("/user/", middleware.MiddlewareSet(handlers.CreateUser)).Methods("POST")
}

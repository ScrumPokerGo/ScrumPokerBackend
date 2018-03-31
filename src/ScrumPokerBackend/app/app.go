package app

import (
	"ScrumPokerBackend/routers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func (a *App) initializeRoutes() {
	a.Router = mux.NewRouter()
	routers.AddAccountRouters(a.Router)
}

func (a *App) Initialize(user, password, dbname string) {
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

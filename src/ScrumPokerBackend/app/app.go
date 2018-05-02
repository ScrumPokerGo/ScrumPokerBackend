package app

import (
	"ScrumPokerBackend/routers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"ScrumPokerBackend/configuration"
	"ScrumPokerBackend/database"
	"ScrumPokerBackend/model"
)


type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	configuration.Init()
	database.Init()
	a.Router = mux.NewRouter()
	routers.AddAccountRouters(a.Router)
	routers.AddUserRouters(a.Router)
	routers.AddProjectRouters(a.Router)
	routers.AddMilestonesRouters(a.Router)
	routers.AddIssuesRouters(a.Router)
}

func (a *App) Migrate(){

	database.DB.AutoMigrate(&model.Account{})
	database.DB.AutoMigrate(&model.User{})
	database.DB.AutoMigrate(&model.Project{})
	database.DB.AutoMigrate(&model.Milestone{})
	database.DB.AutoMigrate(&model.Issue{})
}


func (a *App) Run(addr string) {
	defer database.Close()
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
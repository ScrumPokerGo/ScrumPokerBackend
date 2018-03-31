package routers

import (
	"ScrumPokerBackend/api"
	"ScrumPokerBackend/middleware"
	"github.com/gorilla/mux"
)

func AddProjectRouters(r *mux.Router) {

	r.HandleFunc("/project", middleware.MiddlewareSet(api.GetProjects)).Methods("GET")
	r.HandleFunc("/project/gitlab/sync", middleware.MiddlewareSet(api.GitlabSyncProject)).Methods("PUT")
	r.HandleFunc("/project/{id}", middleware.MiddlewareSet(api.GetProject)).Methods("GET")
	r.HandleFunc("/project/", middleware.MiddlewareSet(api.CreateProject)).Methods("POST")

}

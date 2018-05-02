package routers

import (
	"ScrumPokerBackend/api"
	"ScrumPokerBackend/middleware"
	"github.com/gorilla/mux"
)

func AddIssuesRouters(r *mux.Router) {

	r.HandleFunc("/issues", middleware.MiddlewareSet(api.GetIssueList)).Methods("GET")
	r.HandleFunc("/project/{project_id}/milestone/{milestone_id}/issue", middleware.MiddlewareSet(api.GetIssuesListByProjectAndMilestone)).Methods("GET")
	r.HandleFunc("/project/{project_id}/milestone/{milestone_id}/issue/{id}", middleware.MiddlewareSet(api.GetIssuesListByProjectAndMilestoneId)).Methods("GET")
	//r.HandleFunc("/milestone/{id}", middleware.MiddlewareSet(api.GetProject)).Methods("GET")

}

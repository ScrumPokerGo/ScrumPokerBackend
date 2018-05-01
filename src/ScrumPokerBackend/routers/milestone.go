package routers

import (
"ScrumPokerBackend/api"
"ScrumPokerBackend/middleware"
"github.com/gorilla/mux"
)

func AddMilestonesRouters(r *mux.Router) {

	r.HandleFunc("/milestone", middleware.MiddlewareSet(api.GetMilestoneList)).Methods("GET")
	r.HandleFunc("/project/{project_id}/milestone", middleware.MiddlewareSet(api.GetMilestoneListByProject)).Methods("GET")
	r.HandleFunc("/project/{project_id}/milestone/{id}", middleware.MiddlewareSet(api.GetMilestoneListByProjectAndId)).Methods("GET")
	//r.HandleFunc("/milestone/{id}", middleware.MiddlewareSet(api.GetProject)).Methods("GET")

}

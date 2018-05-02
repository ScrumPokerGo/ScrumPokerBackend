package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"ScrumPokerBackend/middleware"
	"encoding/json"
	"ScrumPokerBackend/model"
	"strconv"
)

func GetIssueList(w http.ResponseWriter, r *http.Request) error {
	issue_list, err := model.GetIssuesAll()
	if err != nil {
		return middleware.StatusError{500, err}
	}

	err = json.NewEncoder(w).Encode(issue_list)
	if err != nil {
		return middleware.StatusError{500, err}
	}
	return nil
}


func GetIssuesListByProjectAndMilestone(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	var err error = nil
	project_id, err := strconv.Atoi(params["project_id"])
	milestone_id,err := strconv.Atoi(params["milestone_id"])

	issue_list, err := model.GetIssuesAllByProjectMilestone(project_id,milestone_id)
	if err != nil {
		return middleware.StatusError{500, err}
	}

	err = json.NewEncoder(w).Encode(issue_list)
	if err != nil {
		return middleware.StatusError{500, err}
	}
	return nil
}

func GetIssuesListByProjectAndMilestoneId(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	var err error = nil
	project_id, err := strconv.Atoi(params["project_id"])
	milestone_id,err := strconv.Atoi(params["milestone_id"])
	id, err := strconv.Atoi(params["id"])

	issue_list, err := model.GetIssuesAllByProjectMilestoneId(project_id,milestone_id,id)
	if err != nil {
		return middleware.StatusError{500, err}
	}

	err = json.NewEncoder(w).Encode(issue_list)
	if err != nil {
		return middleware.StatusError{500, err}
	}
	return nil
}

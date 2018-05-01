package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"ScrumPokerBackend/middleware"
	"ScrumPokerBackend/model"
	"encoding/json"
	"strconv"
	"ScrumPokerBackend/utils"


)

func GetMilestoneListByProject(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	var err error = nil
	project_id, err := strconv.Atoi(params["project_id"])

	milestone_list, err := model.GetMilestonesAllByProject(project_id)
	if err != nil {
		return middleware.StatusError{500, err}
	}

	err = json.NewEncoder(w).Encode(milestone_list)
	if err != nil {
		return middleware.StatusError{500, err}
	}
	return nil
}


func GetMilestoneListByProjectAndId(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	var err error = nil
	project_id, err := strconv.Atoi(params["project_id"])
	milestone_id, err := strconv.Atoi(params["id"])

	milestone_list, err := model.GetMilestoneByProjectAndId(project_id,milestone_id)
	if err != nil {
		return middleware.StatusError{500, err}
	}

	err = json.NewEncoder(w).Encode(milestone_list)
	if err != nil {
		return middleware.StatusError{500, err}
	}
	return nil
}


func GetMilestoneList(w http.ResponseWriter, r *http.Request) error {
	var err error = nil
	milestone_list, err := model.GetMilestonesAll()
	if err != nil {
		return middleware.StatusError{500, err}
	}

	err = json.NewEncoder(w).Encode(milestone_list)
	if err != nil {
		return middleware.StatusError{500, err}
	}
	return nil
}
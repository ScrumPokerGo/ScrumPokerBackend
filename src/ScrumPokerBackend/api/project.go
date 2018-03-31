
package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"ScrumPokerBackend/middleware"
	"ScrumPokerBackend/model"
	"encoding/json"
	"strconv"
	"github.com/xanzy/go-gitlab"
	"fmt"
	"ScrumPokerBackend/configuration"
)

func GetProjects(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func GitlabSyncProject(w http.ResponseWriter, r *http.Request) error {

	git := gitlab.NewClient(nil, configuration.Configuration.GITLAB_TOKEN)
	git.SetBaseURL(configuration.Configuration.GITLAB_ENDPOINT)
	project_service := git.Projects
	//list_options := gitlab.ListProjectsOptions{}
	project_list, _, err := project_service.ListProjects(nil,nil)
	if (err==nil){
		fmt.Println(project_list)
	}else{
		fmt.Println("ERROR")
	}

	return nil
}

func GetProject(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	var err error = nil
	project_id, err := strconv.ParseUint(params["id"], 10, 8)
	if err != nil {
		return middleware.StatusError{500, err}
	}

	project, err := model.GetProject(int(project_id))
	if err != nil {
		return middleware.StatusError{500, err}
	}

	err = json.NewEncoder(w).Encode(project)
	if err != nil {
		return middleware.StatusError{500, err}
	}
	return nil
}

func CreateProject(w http.ResponseWriter, r *http.Request) error {
	return nil
}

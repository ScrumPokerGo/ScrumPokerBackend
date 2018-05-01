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
	project, err := model.GetProjectsAll()
	if err != nil {
		return middleware.StatusError{500, err}
	}

	err = json.NewEncoder(w).Encode(project)
	if err != nil {
		return middleware.StatusError{500, err}
	}
	return nil
}
//this Function needs to be improved by:
//page limitation
//loop efficiency (Update if exists)
func GitlabSyncProject(w http.ResponseWriter, r *http.Request) error {
	git := gitlab.NewClient(nil, configuration.Config.GITLAB_TOKEN)
	git.SetBaseURL(configuration.Config.GITLAB_ENDPOINT)
	project_service := git.Projects
	milestone := git.Milestones
	issues := git.Issues
	list_options := gitlab.ListProjectsOptions{ListOptions:gitlab.ListOptions{Page:1,PerPage:100}}
	project_list, _, err := project_service.ListProjects(&list_options,nil)

	for _, element := range project_list {
		element_id := element.ID
		_,err := model.GetProjectByGitlabId(element_id)

		if (err != nil){
			project_val := model.NewProject(element.Name,element.ID)
			project_val.Create()
			list_options_milestone := gitlab.ListMilestonesOptions{ListOptions:gitlab.ListOptions{Page:1,PerPage:100}}
			milestones_list,_,err_milestone := milestone.ListMilestones(project_val.GitlabID,&list_options_milestone)
			if err_milestone == nil{
				for _,milestone_element := range milestones_list{
					if (milestone_element.State != "closed"){
						milestone_val := project_val.NewMilestone(milestone_element.Title,milestone_element.ID)
						list_issues_options := gitlab.ListProjectIssuesOptions{Milestone:&milestone_val.Name,ListOptions:gitlab.ListOptions{Page:1,PerPage:100}}
						list_issues_values,_,err_issues := issues.ListProjectIssues(project_val.GitlabID,&list_issues_options)
						if(err_issues == nil){
							for _,issue_val := range list_issues_values{
								issue_element := milestone_val.NewIssue(issue_val.Title,issue_val.ID)
								issue_element.Create()
							}
						}else{
							fmt.Println(err_issues)
						}
					}


				}
			}

		}


	}
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
	project_id, err := strconv.Atoi(params["id"])
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

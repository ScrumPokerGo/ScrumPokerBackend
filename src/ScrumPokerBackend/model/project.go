package model

import (
	"time"
	"ScrumPokerBackend/database"
	"ScrumPokerBackend/error"
)

type Project struct {
	ID           int       `gorm:"primary_key" json:"id"`
	Name         string    `json:"name"`
	GitlabID	 int	   `json:"gitlab_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Deleted      bool      `json:"deleted"`
	Account      Account   `gorm:"foreignkey:AccountRefer" json:"account"`
	AccountRefer int
}

func NewProject(name string) *Project {
	project := Project{}
	project.Name = name
	return &project
}

func GetProject(id int) (*Project, error) {
	p := Project{}
	db := database.DB.First(&p, id)
	if db.RowsAffected == 0 {
		return nil, errors.New("Project not found", 404)
	}
	return &p, nil
}

func GetProjectsAll() ([]Project, error) {
	project_list := []Project{}
	db := database.DB.Find(&project_list)
	if db.RowsAffected == 0 {
		return nil, errors.New("Project not found", 404)
	}
	return project_list, nil
}

func GetProjectByGitlabId(id int) (*Project, error) {
	p := Project{}
	db := database.DB.First(&p, "GitlabID = ?",id)
	if db.RowsAffected == 0 {
		return nil, errors.New("Project not found", 404)
	}
	return &p, nil

}


func (m *Project)GitlabLoadMilestone(name string){


}
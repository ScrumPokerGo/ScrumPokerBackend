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


func (p *Project) Create() error{
	db := database.DB.Save(&p)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func NewProject(name string,gitlab_id int) *Project {
	project := Project{}
	project.Name = name
	project.GitlabID = gitlab_id
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
		return project_list, nil
	}
	return project_list, nil
}

func GetMilestoneByGitlabProjectId(id int) (*Milestone,error){
	m := Milestone{}
	db := database.DB.First(&m, "GitlabID = ?",id)
	if db.RowsAffected == 0 {
		return nil, errors.New("Milestone not found", 404)
	}
	return &m, nil
}

func GetProjectByGitlabId(id int) (*Project, error) {
	p := Project{}
	db := database.DB.First(&p, "GitlabID = ?",id)
	if db.RowsAffected == 0 {
		return nil, errors.New("Project not found", 404)
	}
	return &p, nil

}


func (p *Project)NewMilestone(title string, gitlab_id int) (*Milestone){
	milestone := Milestone{GitlabID:gitlab_id,Name:title,Project:*p}
	milestone.Create()
	return &milestone
}

func (p *Project)GetMilestonesByProjectId() (*Milestone){

	return nil
}
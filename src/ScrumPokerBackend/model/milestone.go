package model

import (
	"time"
	"ScrumPokerBackend/database"
	"ScrumPokerBackend/error"
)

type Milestone struct {
	ID           int       `gorm:"primary_key" json:"id"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Deleted      bool      `json:"deleted"`
	GitlabID	 int	   `json:"gitlab_id"`
	Project      Project   `gorm:"foreignkey:ProjectRefer" json:"project"`
	ProjectRefer int
}

func (m *Milestone) Create() error{
	db := database.DB.Save(&m)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (m *Milestone)NewIssue(title string, gitlab_id int) (*Issue){
	issue := Issue{Name:title,GitlabID:gitlab_id,Milestone:*m}
	issue.Create()
	return &issue
}

func GetMilestoneByProjectAndId(project_id int,id int) (*Milestone, error) {
	m := Milestone{}
	db := database.DB.First(&m, "project_refer = ? and id = ?", project_id,id)
	if db.RowsAffected == 0 {
		return nil, errors.New("Milestone not found", 404)
	}
	return &m, nil
}

func GetMilestonesAll() ([]Milestone, error) {
	milestone_list := []Milestone{}
	db := database.DB.Find(&milestone_list)
	if db.RowsAffected == 0 {
		return milestone_list, nil
	}
	return milestone_list, nil
}


func GetMilestonesAllByProject(project_id int) ([]Milestone, error) {
	milestone_list := []Milestone{}
	db := database.DB.Find(&milestone_list,"project_refer = ?", project_id)
	if db.RowsAffected == 0 {
		return milestone_list, nil
	}
	return milestone_list, nil
}
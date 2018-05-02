package model

import (
	"time"
	"ScrumPokerBackend/database"
)

type Issue struct {
	ID             int       `gorm:"primary_key" json:"id"`
	Name           string    `json:"name"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	GitlabID	   int	     `json:"gitlab_id"`
	TimeEstimated  int		 `json:"time_estimated"`
	Milestone      Milestone `gorm:"foreignkey:MilestoneRefer" json:"milestone"`
	MilestoneRefer int
}

func (i *Issue) Create() error{
	db := database.DB.Save(&i)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func GetIssuesAll() ([]Issue, error) {
	issue_list := []Issue{}
	db := database.DB.Find(&issue_list)
	if db.RowsAffected == 0 {
		return issue_list, nil
	}
	return issue_list, nil
}


func GetIssuesAllByProjectMilestone(project_id int,milestone_id int) ([]Issue, error) {
	issue_list := []Issue{}
	db := database.DB.Joins("left join milestones on milestones.id=milestone_refer and milestones.project_refer = ?",project_id).Find(&issue_list,"milestone_refer = ?", milestone_id)
	if db.RowsAffected == 0 {
		return issue_list, nil
	}
	return issue_list, nil
}


func GetIssuesAllByProjectMilestoneId(project_id int,milestone_id int,id int) ([]Issue, error) {
	issue_list := []Issue{}
	db := database.DB.Joins("left join milestones on milestones.id=milestone_refer and milestones.project_refer = ?",project_id).Find(&issue_list,"milestone_refer = ? AND issues.id = ?", milestone_id,id)
	if db.RowsAffected == 0 {
		return issue_list, nil
	}
	return issue_list, nil
}

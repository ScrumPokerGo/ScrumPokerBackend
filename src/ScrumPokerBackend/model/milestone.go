package model

import (
	"time"
	"ScrumPokerBackend/database"
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
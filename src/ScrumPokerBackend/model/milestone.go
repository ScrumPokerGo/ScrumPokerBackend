package model

import (
	"time"
)

type Milestone struct {
	ID           int       `gorm:"primary_key" json:"id"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Deleted      bool      `json:"deleted"`
	Project      Project   `gorm:"foreignkey:ProjectRefer" json:"project"`
	ProjectRefer int
}

func (m *Milestone)GitlabLoadIssues(){


}
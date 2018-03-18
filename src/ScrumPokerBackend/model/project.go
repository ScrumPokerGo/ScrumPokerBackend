package model

import (
	"time"
)

type Project struct {
	ID           int       `gorm:"primary_key" json:"id"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Deleted      bool      `json:"deleted"`
	Account      Account   `gorm:"foreignkey:AccountRefer" json:"account"`
	AccountRefer int
}

func NewProject(name string) *User {
	user := User{}
	user.Name = name
	return &user
}

func (m *Project)GitlabLoadMilestone(name string){


}
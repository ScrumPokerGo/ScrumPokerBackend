package model

import (
	"time"
)

type Issue struct {
	ID             int       `gorm:"primary_key" json:"id"`
	Name           string    `json:"name"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Milestone      Milestone `gorm:"foreignkey:MilestoneRefer" json:"milestone"`
	MilestoneRefer int
}

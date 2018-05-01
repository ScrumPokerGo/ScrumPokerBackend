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



package model

import (
	"ScrumPokerBackend/database"
	"time"
	"github.com/xanzy/go-gitlab"
	"fmt"
	"ScrumPokerBackend/configuration"
)

type Account struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Status    bool      `json:"status"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Deleted   bool      `json:"deleted"`
}

func NewAccount(name string) *Account {
	ac := Account{}
	ac.ID = 0
	ac.Name = name
	return &ac
}

func GetAccount(id int) (*Account, error) {
	a := Account{}
	db := database.DB.First(&a, id)
	if db.Error != nil {
		return nil, db.Error
	}
	return &a, nil
}

func GetAccountAll() ([]Account, error) {
	account_list := []Account{}
	db := database.DB.Find(&account_list)
	if db.Error != nil {
		return nil, db.Error
	}

	return account_list, nil
}

func (a *Account) Save() error {
	db := database.DB.Save(&a)
	if db.Error != nil {
		return db.Error
	}
	return nil

}

func (a *Account) Create() error {
	db := database.DB.Create(&a)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (a *Account) Delete() {
	database.DB.Delete(&a)
}

func (a *Account) GitlabImport(project_id interface{}){
	git := gitlab.NewClient(nil, configuration.Config.GITLAB_TOKEN)
	git.SetBaseURL(configuration.Config.GITLAB_ENDPOINT)
	project, _, err := git.Projects.GetProject(project_id)
	if (err == nil){
		fmt.Println(project.ID)
	}else{
		fmt.Println(err)
	}
}
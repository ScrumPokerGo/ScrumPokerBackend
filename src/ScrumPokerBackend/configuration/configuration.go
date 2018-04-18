package configuration

import (
	"encoding/json"
	"os"
	"fmt"
	"ScrumPokerBackend/utils"
)

var Config Configuration

func  Init(){
	LoadJson()
	utils.NewLog(Config.LOGGER_PATH)
	utils.Log.Println("Configuration Loaded")
}


type Configuration struct {
	GITLAB_ENDPOINT    string
	GITLAB_TOKEN	string
	DATABASE_CONNECTION   string
	LOGGER_PATH string
}

func LoadJson(){
	file, err := os.Open("conf.json")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	error := decoder.Decode(&Config)
	if error != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println(Config.GITLAB_ENDPOINT) // output: [UserA, UserB]
}


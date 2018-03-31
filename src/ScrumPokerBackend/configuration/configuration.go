package configuration

import (
	"encoding/json"
	"os"
	"fmt"
)

var Config Configuration

func Init(){
	LoadJson()
}


type Configuration struct {
	GITLAB_ENDPOINT    string
	GITLAB_TOKEN	string
	DATABASE_CONNECTION   string
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
	fmt.Println(Config.GITLAB_ENDPOINT) // output: [UserA, UserB]
}

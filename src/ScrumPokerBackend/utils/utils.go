package utils

import (
	"net/http"
	"log"
	"os"

)

type Error struct {
}

func CheckErr(w http.ResponseWriter, err error) {
	if err != nil {
		//json.NewEncoder(w).Encode()
		//log.Fatal("ERROR:", err)
	}
}

var (
	Log      *log.Logger
)

func NewLog(logpath string) {
	println("LogFile: " + logpath)
	file, err := os.Create(logpath)
	if err != nil {
		panic(err)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
}
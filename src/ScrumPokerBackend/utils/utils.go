package utils

import (
	"net/http"
	"log"
	"os"
	"encoding/json"
	"bytes"
)

func PrettyPrintJSON(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")
	return out.Bytes(), err
}

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
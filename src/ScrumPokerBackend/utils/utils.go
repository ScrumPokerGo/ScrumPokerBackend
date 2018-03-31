package utils

import (
	"net/http"
)

type Error struct {
}

func CheckErr(w http.ResponseWriter, err error) {
	if err != nil {
		//json.NewEncoder(w).Encode()
		//log.Fatal("ERROR:", err)
	}
}

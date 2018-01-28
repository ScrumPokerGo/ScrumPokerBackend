package main

import (
	"ScrumPokerBackend/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	m := model.NewAccount("TEST")
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", handler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}

package handlers

import (
	"net/http"
	"encoding/json"
	"ScrumPokerBackend/model"
	"github.com/gorilla/mux"
	"strconv"
	"ScrumPokerBackend/middleware"
)


func GetAccounts(w http.ResponseWriter, r *http.Request) error {
	account :=model.GetAccountAll()
	json.NewEncoder(w).Encode(account)
	return nil
}

func GetAccount(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	account_id, err := strconv.ParseUint(params["id"],10,8)
	if err != nil{
		return middleware.StatusError{500,err}
	}
	account :=model.GetAccount(uint(account_id))
	json.NewEncoder(w).Encode(account)
	return nil
}

func CreateAccount(w http.ResponseWriter, r *http.Request) error {
	var account model.Account
	_ = json.NewDecoder(r.Body).Decode(&account)
	account.Save()
	json.NewEncoder(w).Encode(account)
	return nil
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	account_id, _ := strconv.ParseUint(params["id"],10,8)
	account :=model.GetAccount(uint(account_id))
	json.NewEncoder(w).Encode(account)
	return nil
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) error{
	params := mux.Vars(r)
	account_id, _ := strconv.ParseUint(params["id"],10,8)
	account :=model.GetAccount(uint(account_id))
	json.NewEncoder(w).Encode(account)
	return nil
}


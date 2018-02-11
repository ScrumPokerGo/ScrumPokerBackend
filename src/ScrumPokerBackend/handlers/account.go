package handlers

import (
	"net/http"
	"encoding/json"
	"ScrumPokerBackend/model"
	"github.com/gorilla/mux"
	"strconv"
	"ScrumPokerBackend/utils"
)


func GetAccounts(w http.ResponseWriter, r *http.Request) {
	account :=model.GetAccountAll()
	json.NewEncoder(w).Encode(account)

}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	account_id, err := strconv.ParseUint(params["id"],10,8)
	utils.CheckErr(w,err)
	account :=model.GetAccount(uint(account_id))
	json.NewEncoder(w).Encode(account)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account model.Account
	_ = json.NewDecoder(r.Body).Decode(&account)
	account.Save()
	json.NewEncoder(w).Encode(account)
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	account_id, _ := strconv.ParseUint(params["id"],10,8)
	account :=model.GetAccount(uint(account_id))
	json.NewEncoder(w).Encode(account)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	account_id, _ := strconv.ParseUint(params["id"],10,8)
	account :=model.GetAccount(uint(account_id))
	json.NewEncoder(w).Encode(account)
}


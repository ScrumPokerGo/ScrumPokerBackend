package api

import (
	"ScrumPokerBackend/middleware"
	"ScrumPokerBackend/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetAccounts(w http.ResponseWriter, r *http.Request) error {
	account, err := model.GetAccountAll()
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(account)
	if err != nil {
		return err

	}
	return nil
}

func GetAccount(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	account_id, err := strconv.ParseUint(params["id"], 10, 8)
	if err != nil {
		return middleware.StatusError{500, err}
	}
	account, err := model.GetAccount(int(account_id))
	if err != nil {
		return middleware.StatusError{500, err}
	}
	err = json.NewEncoder(w).Encode(account)
	if err != nil {
		return middleware.StatusError{500, err}
	}
	return nil
}

func CreateAccount(w http.ResponseWriter, r *http.Request) error {
	var account model.Account
	var err error
	err = json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		return middleware.StatusError{500, err}
	}

	err = account.Save()
	if err != nil {
		return middleware.StatusError{500, err}
	}

	err = json.NewEncoder(w).Encode(account)
	if err != nil {
		return middleware.StatusError{500, err}
	}
	return nil
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	var err error
	account_id, err := strconv.ParseUint(params["id"], 10, 8)
	if err != nil {
		return middleware.StatusError{500, err}
	}

	account, err := model.GetAccount(int(account_id))
	if err != nil {
		return middleware.StatusError{500, err}
	}

	err = json.NewEncoder(w).Encode(account)
	if err != nil {
		return middleware.StatusError{500, err}
	}
	return nil
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	var err error

	account_id, err := strconv.ParseUint(params["id"], 10, 8)
	if err != nil {
		return middleware.StatusError{500, err}
	}
	account, err := model.GetAccount(int(account_id))
	if err != nil {
		return middleware.StatusError{500, err}
	}
	err = json.NewEncoder(w).Encode(account)
	if err != nil {
		return middleware.StatusError{500, err}
	}
	return nil
}

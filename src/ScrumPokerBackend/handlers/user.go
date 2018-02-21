package handlers

import (
	"net/http"
	"encoding/json"
	"ScrumPokerBackend/model"
	"github.com/gorilla/mux"
	"strconv"
	"ScrumPokerBackend/middleware"
)


func GetUsers(w http.ResponseWriter, r *http.Request) error {
	user_list,err := model.GetUserAll()
	if (err!=nil){
		return middleware.StatusError{500,err}
	}

	err = json.NewEncoder(w).Encode(user_list)
	if(err!=nil){
		return middleware.StatusError{500,err}
	}

	return nil
}

func GetUser(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	var err error =nil
	user_id, err := strconv.ParseUint(params["id"],10,8)
	if err != nil{
		return middleware.StatusError{500,err}
	}

	user,err := model.GetUser(uint(user_id))
	if err != nil{
		return middleware.StatusError{500,err}
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil{
		return middleware.StatusError{500,err}
	}
	return nil
}

func CreateUser(w http.ResponseWriter, r *http.Request) error {
	var user model.User
	var err error
	err = json.NewDecoder(r.Body).Decode(&user)
	if err!=nil{
		return middleware.StatusError{500,err}
	}
	err = user.Create()
	if err!=nil{
		return middleware.StatusError{500,err}
	}
	err= json.NewEncoder(w).Encode(user)
	if err!=nil{
		return middleware.StatusError{500,err}
	}
	return nil
}

func UpdateUser(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	var err error
	user_id, err := strconv.ParseUint(params["id"],10,8)
	if err!=nil{
		return middleware.StatusError{500,err}
	}
	user,err :=model.GetUser(uint(user_id))
	if err!=nil{
		return middleware.StatusError{500,err}
	}
	err = json.NewEncoder(w).Encode(user)
	if err!=nil{
		return middleware.StatusError{500,err}
	}
	return nil
}

func DeleteUser(w http.ResponseWriter, r *http.Request) error{
	params := mux.Vars(r)
	var err error
	user_id, err := strconv.ParseUint(params["id"],10,8)
	if err!=nil{
		return middleware.StatusError{500,err}
	}
	user,err :=model.GetUser(uint(user_id))
	if err!=nil{
		return middleware.StatusError{500,err}
	}
	err = json.NewEncoder(w).Encode(user)
	if err!=nil{
		return middleware.StatusError{500,err}
	}
	return nil
}


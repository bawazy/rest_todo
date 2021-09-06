package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bawazy/rest_todo/pkg/models"
	"github.com/bawazy/rest_todo/pkg/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	utils.ParseBody(r, user)

	existinguser := models.GetUserbyUsername(user.Username)

	if len(existinguser) == 0 {
		w.WriteHeader(http.StatusConflict)
	} else if user.Password == existinguser[0].Password {

		res, _ := json.Marshal(existinguser)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	utils.ParseBody(r, user)

	existinguser := models.GetUserbyUsername(user.Username)

	if len(existinguser) != 0 {
		w.WriteHeader(http.StatusConflict)
	} else {
		t := user.CreateUser()
		res, _ := json.Marshal(t)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	allusers := models.GetAllUsers()
	res, _ := json.Marshal(allusers)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	users "github.com/Threx-code/go-api/models/user"
)

type ResError struct {
	Error string `json:"error"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &users.CreateRequest{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(body), CreateUser)
	CreateUser.CreateNewUser(w)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	LoginUser := &users.LoginRequest{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(body), LoginUser)
	LoginUser.LoginUser(w)
}

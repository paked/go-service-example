package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/paked/gerrycode/communicator"
	"github.com/paked/pay/models"
	"github.com/paked/restrict"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	coms := communicator.New(w)

	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	u, err := models.NewUser(username, password, email)
	if err != nil {
		coms.Error("Unable to create user")
		return
	}

	coms.OKWithData("user", u)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	coms := communicator.New(w)

	username := r.FormValue("username")
	password := r.FormValue("password")

	u, err := models.FetchUser(username)
	if err != nil {
		fmt.Println(err)
		coms.Error("Could not fetch user")
		return
	}

	if !u.Login(password) {
		coms.Error("Incorrect password")
		return
	}

	claims := make(map[string]interface{})
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	ts, err := restrict.Token(claims)
	if err != nil {
		coms.Fail("Failure signing the token")
		return
	}

	coms.OKWithData("token", ts)
}

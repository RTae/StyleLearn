package handler

import (
	"backend/helper"
	"backend/user"
	"encoding/json"
	"fmt"
	"net/http"
)

//Handler Login Auth CorchroshDB
func Handler(w http.ResponseWriter, r *http.Request) {
	helper.SetupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	if (*r).Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		email := r.FormValue("email")
		pasword := r.FormValue("password")

		// Login
		u := user.User{}
		logs := u.Login(email, pasword)

		json.NewEncoder(w).Encode(logs)
	} else {
		fmt.Fprintf(w, "Please use post medthod")
	}
}

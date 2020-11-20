package handler

import (
	"backend/helper"
	"backend/user"
	"encoding/json"
	"fmt"
	"net/http"
)

//Handler Insert Auth MongoDB
func Handler(w http.ResponseWriter, r *http.Request) {
	helper.SetupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		firstname := r.FormValue("firstname")
		familyname := r.FormValue("familyname")
		brithday := r.FormValue("brithday")
		sex := r.FormValue("sex")
		email := r.FormValue("email")
		password := r.FormValue("password")
		userType := r.FormValue("userType")
		educationType := r.FormValue("educationType")

		// Register
		u := user.User{}
		logs := u.Register(firstname, familyname, brithday, sex, email, password, userType, educationType)

		json.NewEncoder(w).Encode(logs)
	} else {
		fmt.Fprintf(w, "Please use post medthod")
	}
}

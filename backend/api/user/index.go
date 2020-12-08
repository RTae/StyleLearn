package handler

import (
	"backend/helper"
	"backend/user"
	"encoding/json"
	"net/http"
)

//Handler Login Auth CorchroshDB
func Handler(w http.ResponseWriter, r *http.Request) {
	helper.SetupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	} else if (*r).Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		id := r.FormValue("id")
		mode := r.FormValue("mode")

		// create Course
		u := user.User{}

		if mode == "1" {
			logs := u.Read(id)
			json.NewEncoder(w).Encode(logs)
		} else if mode == "2" {
			logs := u.ReadProfile(id)
			json.NewEncoder(w).Encode(logs)
		} else {
			logs := u.GetAllUser()
			json.NewEncoder(w).Encode(logs)
		}
	} else if (*r).Method == "PUT" {
		w.Header().Set("Content-Type", "application/json")

		u := user.User{}
		id := r.FormValue("id")
		firstname := r.FormValue("firstname")
		familyname := r.FormValue("familyname")
		birthday := r.FormValue("birthday")
		linkPic := r.FormValue("linkPic")
		sex := r.FormValue("sex")
		edu := r.FormValue("edu")
		bio := r.FormValue("bio")

		logs := u.Update(id, firstname, familyname, birthday, sex, linkPic, edu, bio)
		json.NewEncoder(w).Encode(logs)
	} else {
		log := make(map[string]interface{})
		log["status"] = "415"
		log["msg"] = "medthod is not provide"
		log["result"] = ""
		json.NewEncoder(w).Encode(log)
	}
}

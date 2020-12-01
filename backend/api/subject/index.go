package handler

import (
	"backend/helper"
	"backend/subject"
	"encoding/json"
	"fmt"
	"net/http"
)

//Handler get all Subject
func Handler(w http.ResponseWriter, r *http.Request) {
	helper.SetupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	if (*r).Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		id := r.FormValue("id")
		mode := r.FormValue("mode")

		// create Subject
		s := subject.Subject{}

		if mode == "1" {
			logs := s.Read(id)
			json.NewEncoder(w).Encode(logs)
		} else {
			logs := s.GetAllSubject()
			json.NewEncoder(w).Encode(logs)
		}
	} else {
		fmt.Fprintf(w, "Please use get medthod")
	}
}

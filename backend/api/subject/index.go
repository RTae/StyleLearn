package handler

import (
	"backend/helper"
	"backend/subject"
	"encoding/json"
	"net/http"
)

//Handler get all Subject
func Handler(w http.ResponseWriter, r *http.Request) {
	helper.SetupResponse(&w, r)

	// create Subject
	s := subject.Subject{}
	if (*r).Method == "OPTIONS" {
		return
	} else if (*r).Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		id := r.FormValue("id")
		mode := r.FormValue("mode")

		if mode == "1" {
			logs := s.Read(id)
			json.NewEncoder(w).Encode(logs)
		} else {
			logs := s.GetAllSubject()
			json.NewEncoder(w).Encode(logs)
		}
	}
}

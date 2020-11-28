package handler

import (
	"backend/helper"
	"backend/progressLesson"
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
	if (*r).Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		uid := r.FormValue("uid")
		cid := r.FormValue("cid")
		mode := r.FormValue("mode")

		// create Subject
		pl := progressLesson.ProgressLesson{}

		if mode == "1" {
			logs := pl.GetAllCourseFronProgressLesson(uid)
			json.NewEncoder(w).Encode(logs)
		} else if mode == "2" {
			logs := pl.GetAllProgressLessonFromCourse(uid, cid)
			json.NewEncoder(w).Encode(logs)
		} else {
			logs := pl.GetAllProgressLesson()
			json.NewEncoder(w).Encode(logs)
		}
	} else {
		fmt.Fprintf(w, "Please use get medthod")
	}
}

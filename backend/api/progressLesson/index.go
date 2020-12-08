package handler

import (
	"backend/helper"
	"backend/progressLesson"
	"encoding/json"
	"net/http"
)

//Handler Login Auth CorchroshDB
func Handler(w http.ResponseWriter, r *http.Request) {
	helper.SetupResponse(&w, r)
	// create progressLesson
	pl := progressLesson.ProgressLesson{}
	if (*r).Method == "OPTIONS" {
		return
	}
	if (*r).Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		uid := r.FormValue("uid")
		lid := r.FormValue("lid")
		mode := r.FormValue("mode")

		if mode == "1" {
			logs := pl.GetAllCourseFronProgressLesson(uid)
			json.NewEncoder(w).Encode(logs)
		} else if mode == "2" {
			logs := pl.GetAllProgressLessonFromCourse(uid, lid)
			json.NewEncoder(w).Encode(logs)
		} else {
			logs := pl.GetAllProgressLesson()
			json.NewEncoder(w).Encode(logs)
		}
	} else if (*r).Method == "POST" {
		uid := r.FormValue("uid")
		lid := r.FormValue("lid")
		quantityDay := r.FormValue("quantityDay")

		logs := pl.Create(uid, lid, quantityDay)
		json.NewEncoder(w).Encode(logs)
	}
}

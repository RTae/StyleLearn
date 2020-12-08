package handler

import (
	"backend/helper"
	"backend/lesson"
	"encoding/json"
	"fmt"
	"net/http"
)

//Handler Login Auth CorchroshDB
func Handler(w http.ResponseWriter, r *http.Request) {
	helper.SetupResponse(&w, r)
	l := lesson.Lesson{}
	if (*r).Method == "OPTIONS" {
		return
	}
	if (*r).Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		id := r.FormValue("id")
		mode := r.FormValue("mode")
		courseName := r.FormValue("courseName")

		if mode == "1" {
			logs := l.Read(id)
			json.NewEncoder(w).Encode(logs)
		} else if mode == "2" {
			logs := l.ReadLessonByCourse(id)
			json.NewEncoder(w).Encode(logs)
		} else if mode == "3" {
			logs := l.GetAllLessonByCourseName(courseName)
			json.NewEncoder(w).Encode(logs)
		} else {
			logs := l.GetAllLesson()
			json.NewEncoder(w).Encode(logs)
		}
	} else {
		fmt.Fprintf(w, "Please use get medthod")
	}
}

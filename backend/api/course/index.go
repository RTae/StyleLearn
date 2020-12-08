package handler

import (
	"backend/course"
	"backend/helper"
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
		id := r.FormValue("id")
		subjectName := r.FormValue("subjectName")
		mode := r.FormValue("mode")

		// create Course
		c := course.Course{}

		if mode == "1" {
			logs := c.Read(id)
			json.NewEncoder(w).Encode(logs)
		} else if mode == "2" {
			logs := c.ReadCourseBySubject(id)
			json.NewEncoder(w).Encode(logs)
		} else if {
			logs := c.GetAllCourseBySubjectName(subjectName)
			json.NewEncoder(w).Encode(logs)
		} else {
			logs := c.GetAllCourse()
			json.NewEncoder(w).Encode(logs)
		}
	} else {
		fmt.Fprintf(w, "Please use get medthod")
	}
}

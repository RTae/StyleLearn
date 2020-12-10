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
	// create Course
	c := course.Course{}
	if (*r).Method == "OPTIONS" {
		return
	}
	if (*r).Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		id := r.FormValue("id")
		limit := r.FormValue("limit")
		subjectName := r.FormValue("subjectName")
		mode := r.FormValue("mode")

		if mode == "1" {
			logs := c.Read(id)
			json.NewEncoder(w).Encode(logs)
		} else if mode == "2" {
			logs := c.ReadCourseBySubject(id)
			json.NewEncoder(w).Encode(logs)
		} else if mode == "3" {
			logs := c.GetAllCourseBySubjectName(subjectName)
			json.NewEncoder(w).Encode(logs)
		} else if mode == "4" {
			if limit == "1" {
				logs := c.GetAllPoppularCourseLimt()
				json.NewEncoder(w).Encode(logs)
			} else {
				logs := c.GetAllPoppularCourse()
				json.NewEncoder(w).Encode(logs)
			}
		} else if mode == "5" {
			if limit == "1" {
				logs := c.GetAllNewestCourseLimt()
				json.NewEncoder(w).Encode(logs)
			} else {
				logs := c.GetAllNewestCourse()
				json.NewEncoder(w).Encode(logs)
			}
		} else {
			logs := c.GetAllCourse()
			json.NewEncoder(w).Encode(logs)
		}
	} else {
		fmt.Fprintf(w, "Please use get medthod")
	}
}

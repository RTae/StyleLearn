package handler

import (
	"backend/helper"
	"backend/video"
	"encoding/json"
	"net/http"
)

//Handler get all Subject
func Handler(w http.ResponseWriter, r *http.Request) {

	// create Subject
	v := video.Video{}
	helper.SetupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	} else if (*r).Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		videoID := r.FormValue("videoID")
		lessonID := r.FormValue("lessonID")
		userID := r.FormValue("userID")
		mode := r.FormValue("mode")

		if mode == "1" {
			logs := v.Read(videoID)
			json.NewEncoder(w).Encode(logs)
		} else if mode == "2" {
			logs := v.GetVideoByLessonID(lessonID)
			json.NewEncoder(w).Encode(logs)
		} else if mode == "3" {
			logs := v.GetAllVideoByUserID(userID)
			json.NewEncoder(w).Encode(logs)
		} else if mode == "4" {
			logs := v.GetVideoShow(videoID)
			json.NewEncoder(w).Encode(logs)
		} else {
			logs := v.GetAllVideo()
			json.NewEncoder(w).Encode(logs)
		}
	} else if (*r).Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		userID := r.FormValue("userID")
		lessonID := r.FormValue("lessonID")
		description := r.FormValue("description")
		videoFile := r.FormValue("videoFile")
		uploadDate := r.FormValue("uploadDate")
		status := r.FormValue("status")
		// mode := r.FormValue("mode")

		logs := v.Create(userID, lessonID, description, videoFile, uploadDate, status)
		json.NewEncoder(w).Encode(logs)
	}
}

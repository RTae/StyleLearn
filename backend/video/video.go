package video

import (
	"strconv"
	"time"

	"backend/entities"
	"backend/helper"

	// Import GORM-related packages.

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Video struct {
}

// Create new subject into platfrom
func (v *Video) Create(userID, lessonID, description, videoFile, uploadDate, status string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var maxUID string
	rowU := db.Table("tbl_videos").Select("max(video_id)").Row()
	rowU.Scan(&maxUID)
	newVID, _ := helper.IncreaseID(maxUID, "v", 1)

	t, err := time.Parse("2006-01-02", uploadDate)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	statusBool, err := strconv.ParseBool(status)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	video := entities.TBL_Video{
		VideoID:     newVID,
		UserID:      userID,
		LessonID:    lessonID,
		Description: description,
		Video:       videoFile,
		UploadDate:  t,
		Status:      &statusBool,
	}

	err = db.Create(&video).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	logs = make(map[string]interface{})
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = ""
	return logs
}

func (v *Video) Read(vid string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var videos []entities.TBL_Video
	err := db.Find(&videos, entities.TBL_Video{VideoID: vid}).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = videos
	return log
}

func (v *Video) Update(vid, userID, lessonID, description, videoFile, uploadDate, status string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	logs = v.Read(vid)
	if logs["status"] != "1" {
		return logs
	}

	t, err := time.Parse("2006-01-02T15:04:05Z07:00", uploadDate)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	statusBool, err := strconv.ParseBool(status)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	err = db.Model(&entities.TBL_Video{VideoID: vid}).Updates(entities.TBL_Video{
		UserID:      userID,
		LessonID:    lessonID,
		Description: description,
		Video:       videoFile,
		UploadDate:  t,
		Status:      &statusBool,
	}).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	logs = make(map[string]interface{})
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = ""
	return logs
}

func (v *Video) Delete(vid string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	err := db.Where("video_id = ?", vid).Delete(&entities.TBL_Video{}).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = ""
	return log
}

func (v *Video) GetAllVideo() map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var videos []entities.TBL_Video
	err := db.Find(&videos).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = videos

	return log
}

type result_video struct {
	VideoID  string
	lessonID string
	Name     string
	UserID   string
}

//GetVideoByLessonID get all video by lessonID
func (v *Video) GetVideoByLessonID(lessonID string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var result []result_video
	err := db.Raw(`	SELECT v.video_id, v.lesson_id, u.firstname AS "name", v.user_id
					FROM tbl_videos v
					INNER JOIN tbl_users u
						ON v.user_id = u.user_id
					WHERE v.lesson_id = ? `, lessonID).Scan(&result).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = result

	return log
}

// GetAllVideoByUserID get all video by userID
func (v *Video) GetAllVideoByUserID(userID string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var result []result_video
	err := db.Raw(`	SELECT v.video_id, l.lesson_id, l.name, v.user_id
					FROM tbl_videos v
					INNER JOIN tbl_lesson_types l
						ON v.lesson_id = l.lesson_id
					WHERE v.user_id = ? `, userID).Scan(&result).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = result

	return log
}

type result_video_show struct {
	VideoID     string
	Name        string
	LessonName  string
	Description string
	Video       string
}

func (v *Video) GetVideoShow(videoID string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var result []result_video_show
	err := db.Raw(`	SELECT v.video_id, u.firstname as "name", l.name as "lesson_name", v.description, v.video
					FROM tbl_videos v
					INNER JOIN tbl_users u
						ON v.user_id = u.user_id
					INNER JOIN tbl_lesson_types l
						ON v.lesson_id = l.lesson_id
					WHERE video_id = ? `, videoID).Scan(&result).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = result

	return log
}

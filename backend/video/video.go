package video

import (
	"os"
	"strconv"
	"strings"
	"time"

	"backend/entities"

	// Import GORM-related packages.
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Video struct {
}

// Create new subject into platfrom
func (v *Video) Create(userID, lessonID, description, videoFile, uploadDate, status string) map[string]interface{} {
	db, logs := v.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var maxUID string
	rowU := db.Table("tbl_videos").Select("max(video_id)").Row()
	rowU.Scan(&maxUID)
	newVID, _ := increaseID(maxUID, "v", 1)

	t, err := time.Parse("2006-01-02", uploadDate)
	if err != nil {
		log := v.errorHandle(err)
		return log
	}

	statusBool, err := strconv.ParseBool(status)
	if err != nil {
		log := v.errorHandle(err)
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
		log := v.errorHandle(err)
		return log
	}

	logs = make(map[string]interface{})
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = ""
	return logs
}

func (v *Video) Read(vid string) map[string]interface{} {
	db, logs := v.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var videos []entities.TBL_Video
	err := db.Find(&videos, entities.TBL_Video{VideoID: vid}).Error
	if err != nil {
		log := v.errorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = videos
	return log
}

func (v *Video) Update(vid, userID, lessonID, description, videoFile, uploadDate, status string) map[string]interface{} {
	db, logs := v.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	logs = v.Read(vid)
	if logs["status"] != "1" {
		return logs
	}

	t, err := time.Parse("2006-01-02T15:04:05Z07:00", uploadDate)
	if err != nil {
		log := v.errorHandle(err)
		return log
	}

	statusBool, err := strconv.ParseBool(status)
	if err != nil {
		log := v.errorHandle(err)
		return log
	}

	err = db.Model(&entities.TBL_Video{}).Where(entities.TBL_Video{VideoID: vid}).Update(entities.TBL_Video{
		UserID:      userID,
		LessonID:    lessonID,
		Description: description,
		Video:       videoFile,
		UploadDate:  t,
		Status:      &statusBool,
	}).Error

	if err != nil {
		log := v.errorHandle(err)
		return log
	}

	logs = make(map[string]interface{})
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = ""
	return logs
}

func (v *Video) Delete(vid string) map[string]interface{} {
	db, logs := v.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	err := db.Where("video_id = ?", vid).Delete(&entities.TBL_Video{}).Error
	if err != nil {
		log := v.errorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = ""
	return log
}

func (v *Video) GetAllVideo() map[string]interface{} {
	db, logs := v.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var videos []entities.TBL_Video
	err := db.Find(&videos).Error
	if err != nil {
		log := v.errorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = videos

	return log
}

func (v *Video) GetVideoByLessonID(lessonID string) map[string]interface{} {
	db, logs := v.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var videos []entities.TBL_Video
	err := db.Find(&videos, entities.TBL_Video{LessonID: lessonID}).Error
	if err != nil {
		log := v.errorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = videos

	return log
}

// Helper Function
func increaseID(id, name string, length int) (string, error) {
	digit, err := strconv.Atoi(id[length:])
	if err != nil {
		return name, err
	}
	digit++
	s := strconv.Itoa(digit)

	newID := name + strings.Repeat("0", 10-length-len(s)) + s
	return newID, nil
}

func (v *Video) initDB() (*gorm.DB, map[string]interface{}) {
	var addr = os.Getenv("COCKROACHDB_URL")
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log := v.errorHandle(err)
		return nil, log
	}
	db.LogMode(true)
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = ""

	return db, log
}

func (v *Video) errorHandle(err error) map[string]interface{} {
	var textError string
	var textStatus string
	if err.Error() == "mongo: no documents in result" {
		textError = "User not found"
		textStatus = "215"
	} else {
		textError = err.Error()
		textStatus = "415"
	}

	logs := make(map[string]interface{})
	logs["status"] = textStatus
	logs["msg"] = textError
	logs["result"] = ""
	return logs
}

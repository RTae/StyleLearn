package lesson

import (
	"os"
	"strconv"
	"strings"

	"backend/entities"

	// Import GORM-related packages.
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Lesson struct {
}

// Create new subject into platfrom
func (l *Lesson) Create(courseID, name, description string) map[string]interface{} {
	db, logs := l.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var maxUID string
	rowU := db.Table("tbl_lesson_types").Select("max(lesson_id)").Row()
	rowU.Scan(&maxUID)
	newLID, _ := increaseID(maxUID, "l", 1)

	lesson := entities.TBL_LessonTypes{
		LessonID:    newLID,
		CourseID:    courseID,
		Name:        name,
		Description: description,
	}

	err := db.Create(&lesson).Error
	if err != nil {
		log := l.errorHandle(err)
		return log
	}

	logs = make(map[string]interface{})
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = ""
	return logs
}

func (l *Lesson) Read(lid string) map[string]interface{} {
	db, logs := l.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var lesson []entities.TBL_LessonTypes
	err := db.Find(&lesson, entities.TBL_LessonTypes{LessonID: lid}).Error
	if err != nil {
		log := l.errorHandle(err)
		return log
	}
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = lesson
	return log
}

func (l *Lesson) Update(lid, courseID, name, description string) map[string]interface{} {
	db, logs := l.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	logs = l.Read(lid)
	if logs["status"] != "1" {
		return logs
	}

	db.Model(&entities.TBL_LessonTypes{}).Where(entities.TBL_LessonTypes{LessonID: lid}).Update(entities.TBL_LessonTypes{CourseID: courseID, Name: name, Description: description})

	logs = make(map[string]interface{})
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = ""
	return logs
}

func (l *Lesson) Delete(lid string) map[string]interface{} {
	db, logs := l.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	err := db.Delete(&entities.TBL_LessonTypes{}, entities.TBL_LessonTypes{LessonID: lid}).Error
	if err != nil {
		log := l.errorHandle(err)
		return log
	}
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = ""
	return log
}

func (l *Lesson) GetAllLesson() map[string]interface{} {
	db, logs := l.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var lessons []entities.TBL_LessonTypes
	err := db.Find(&lessons).Error
	if err != nil {
		log := l.errorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = lessons

	return log
}

func (l *Lesson) ReadLessonByCourse(cid string) map[string]interface{} {
	db, logs := l.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var lesson []entities.TBL_LessonTypes
	err := db.Find(&lesson, entities.TBL_LessonTypes{CourseID: cid}).Error
	if err != nil {
		log := l.errorHandle(err)
		return log
	}
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = lesson
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

func (l *Lesson) initDB() (*gorm.DB, map[string]interface{}) {
	var addr = os.Getenv("COCKROACHDB_URL")
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log := l.errorHandle(err)
		return nil, log
	}
	db.LogMode(true)
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = ""

	return db, log
}

func (l *Lesson) errorHandle(err error) map[string]interface{} {
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

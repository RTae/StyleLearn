package lesson

import (
	"os"
	"strconv"
	"strings"

	"backend/entities"
	"backend/helper"

	// Import GORM-related packages.
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Lesson struct {
}

// Create new subject into platfrom
func (l *Lesson) Create(courseID, name, description string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var maxUID string
	rowU := db.Table("tbl_lesson_types").Select("max(lesson_id)").Row()
	rowU.Scan(&maxUID)
	newLID, _ := helper.IncreaseID(maxUID, "l", 1)

	lesson := entities.TBL_LessonTypes{
		LessonID:    newLID,
		CourseID:    courseID,
		Name:        name,
		Description: description,
	}

	err := db.Create(&lesson).Error
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

func (l *Lesson) Read(lid string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var lesson []entities.TBL_LessonTypes
	err := db.Find(&lesson, entities.TBL_LessonTypes{LessonID: lid}).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = lesson
	return log
}

func (l *Lesson) Update(lid, courseID, name, description string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

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
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	err := db.Delete(&entities.TBL_LessonTypes{}, entities.TBL_LessonTypes{LessonID: lid}).Error
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

func (l *Lesson) GetAllLesson() map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var lessons []entities.TBL_LessonTypes
	err := db.Find(&lessons).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = lessons

	return log
}

func (l *Lesson) ReadLessonByCourse(cid string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var lesson []entities.TBL_LessonTypes
	err := db.Find(&lesson, entities.TBL_LessonTypes{CourseID: cid}).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = lesson
	return log
}

type result_lesson struct {
	LessonID string
	Name     string
}

func (l *Lesson) GetAllLessonByCourseName(courseName string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var result []result_lesson
	err := db.Raw(`	SELECT lesson_id, name 
					FROM tbl_lesson_types
					WHERE course_id = ( SELECT course_id
										FROM tbl_course_types
										WHERE name = ? ) `, courseName).Scan(&result).Error
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

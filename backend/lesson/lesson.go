package lesson

import (
	"backend/entities"
	"backend/helper"
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

	err := db.Model(&entities.TBL_LessonTypes{LessonID: lid}).Updates(entities.TBL_LessonTypes{
		CourseID:    courseID,
		Name:        name,
		Description: description,
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
	CourseName        string
	LessonID          string
	LessonName        string
	LessonDescription string
}

func (l *Lesson) GetAllLessonByCourseID(courseID string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var result []result_lesson
	err := db.Raw(`	SELECT l.lesson_id, l.name as "lesson_name", l.description as "lesson_description",
						   c.name as "course_name"
					FROM tbl_lesson_types l
					INNER JOIN tbl_course_types c
						ON l.course_id = c.course_id
					WHERE l.course_id = ?`, courseID).Scan(&result).Error
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

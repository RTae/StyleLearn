package progressLesson

import (
	"strconv"

	"backend/entities"
	"backend/helper"

	// Import GORM-related packages.

	"gorm.io/gorm/clause"
)

type ProgressLesson struct {
}

// Create new subject into platfrom
func (pl *ProgressLesson) Create(userID, lessonID, quantityDay string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	quantityDayInt, err := strconv.ParseInt(quantityDay, 10, 64)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	progressLesson := entities.TBL_ProgressLesson{
		UserID:      userID,
		LessonID:    lessonID,
		QuantityDay: quantityDayInt,
	}

	err = db.Clauses(clause.OnConflict{DoNothing: true}).Create(&progressLesson).Error
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

func (pl *ProgressLesson) Read(uid, lid string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var lesson []entities.TBL_ProgressLesson
	err := db.Find(&lesson, entities.TBL_ProgressLesson{UserID: uid, LessonID: lid}).Error
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

func (pl *ProgressLesson) Update(uid, lid string, quantityDay int64) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	logs = pl.Read(uid, lid)
	if logs["status"] != "1" {
		return logs
	}
	err := db.Model(&entities.TBL_ProgressLesson{UserID: uid, LessonID: lid}).Updates(entities.TBL_ProgressLesson{
		QuantityDay: quantityDay,
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

func (pl *ProgressLesson) Delete(uid, lid string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	err := db.Delete(&entities.TBL_ProgressLesson{}, entities.TBL_ProgressLesson{UserID: uid, LessonID: lid}).Error
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

func (pl *ProgressLesson) GetAllProgressLesson() map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var lessons []entities.TBL_ProgressLesson
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

type resultCourse struct {
	CourseID    string
	CourseName  string
	SubjectName string
}

func (pl *ProgressLesson) GetAllCourseFronProgressLesson(uid string) map[string]interface{} {
	db, log := helper.InitDB()
	if log["status"] != "1" {
		return log
	}

	var result []resultCourse
	err := db.Raw(`	SELECT DISTINCT(c.course_id) AS "course_id", c.name AS "course_name", 
					s.name AS "subject_name"
					FROM tbl_progress_lessons pl
					INNER JOIN tbl_lesson_types l
						ON pl.lesson_id = l.lesson_id
					INNER JOIN tbl_course_types c
						ON l.course_id = c.course_id
					INNER JOIN tbl_subject_types s
						ON c.subject_id = s.subject_id
					WHERE pl.user_id =  ? `, uid).Scan(&result).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	log = make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = result
	return log
}

type resultLesson struct {
	LessonID   string
	CourseName string
	LessonName string
}

func (pl *ProgressLesson) GetAllProgressLessonFromCourse(uid, cid string) map[string]interface{} {
	db, log := helper.InitDB()
	if log["status"] != "1" {
		return log
	}

	var result []resultLesson
	err := db.Raw(`	SELECT c.name as course_name, pl.lesson_id , l.name as lesson_name
					FROM tbl_progress_lessons pl
					INNER JOIN tbl_lesson_types l
						ON pl.lesson_id = l.lesson_id
					INNER JOIN tbl_course_types c
						ON l.course_id = c.course_id
					WHERE pl.user_id = ? AND c.course_id = ?`, uid, cid).Scan(&result).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	log = make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = result
	return log
}

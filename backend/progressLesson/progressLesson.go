package progressLesson

import (
	"os"
	"strconv"
	"strings"

	"backend/entities"

	// Import GORM-related packages.
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ProgressLesson struct {
}

// Create new subject into platfrom
func (pl *ProgressLesson) Create(userID, lessonID string, quantityDay int) map[string]interface{} {
	db, logs := pl.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	lesson := entities.TBL_ProgressLessoon{
		UserID:      userID,
		LessonID:    lessonID,
		QuantityDay: quantityDay,
	}

	err := db.Create(&lesson).Error
	if err != nil {
		log := pl.errorHandle(err)
		return log
	}

	logs = make(map[string]interface{})
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = ""
	return logs
}

func (pl *ProgressLesson) Read(uid, lid string) map[string]interface{} {
	db, logs := pl.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var lesson []entities.TBL_ProgressLessoon
	err := db.Find(&lesson, entities.TBL_ProgressLessoon{UserID: uid, LessonID: lid}).Error
	if err != nil {
		log := pl.errorHandle(err)
		return log
	}
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = lesson
	return log
}

func (pl *ProgressLesson) Update(uid, lid string, quantityDay int) map[string]interface{} {
	db, logs := pl.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	logs = pl.Read(uid, lid)
	if logs["status"] != "1" {
		return logs
	}

	db.Model(&entities.TBL_ProgressLessoon{}).Where(entities.TBL_ProgressLessoon{UserID: uid, LessonID: lid}).Update(entities.TBL_ProgressLessoon{QuantityDay: quantityDay})

	logs = make(map[string]interface{})
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = ""
	return logs
}

func (pl *ProgressLesson) Delete(uid, lid string) map[string]interface{} {
	db, logs := pl.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	err := db.Delete(&entities.TBL_ProgressLessoon{}, entities.TBL_ProgressLessoon{UserID: uid, LessonID: lid}).Error
	if err != nil {
		log := pl.errorHandle(err)
		return log
	}
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = ""
	return log
}

func (pl *ProgressLesson) GetAllProgressLesson() map[string]interface{} {
	db, logs := pl.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var lessons []entities.TBL_ProgressLessoon
	err := db.Find(&lessons).Error
	if err != nil {
		log := pl.errorHandle(err)
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
	db, log := pl.initDB()
	if log["status"] != "1" {
		return log
	}
	defer db.Close()

	var result []resultCourse
	err := db.Raw(`	SELECT DISTINCT(c.course_id) AS "course_id", c.name AS "course_name", 
					s.name AS "subject_name"
					FROM tbl_progress_lesson pl
					INNER JOIN tbl_lesson_types l
						ON pl.lesson_id = l.lesson_id
					INNER JOIN tbl_course_types c
						ON l.course_id = c.course_id
					INNER JOIN tbl_subject_types s
						ON c.subject_id = s.subject_id
					WHERE pl.user_id =  ? `, uid).Scan(&result).Error
	if err != nil {
		log := pl.errorHandle(err)
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
	LessonName string
}

func (pl *ProgressLesson) GetAllProgressLessonFromCourse(uid, cid string) map[string]interface{} {
	db, log := pl.initDB()
	if log["status"] != "1" {
		return log
	}
	defer db.Close()

	var result []resultLesson
	err := db.Raw(`	SELECT pl.lesson_id , l.name as lesson_name
					FROM tbl_progress_lesson pl
					INNER JOIN tbl_lesson_types l
						ON pl.lesson_id = l.lesson_id
					INNER JOIN tbl_course_types c
						ON l.course_id = c.course_id
					WHERE pl.user_id = ? AND c.course_id = ?`, uid, cid).Scan(&result).Error
	if err != nil {
		log := pl.errorHandle(err)
		return log
	}

	log = make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = result
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

func (pl *ProgressLesson) initDB() (*gorm.DB, map[string]interface{}) {
	var addr = os.Getenv("COCKROACHDB_URL")
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log := pl.errorHandle(err)
		return nil, log
	}
	db.LogMode(true)
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = ""

	return db, log
}

func (pl *ProgressLesson) errorHandle(err error) map[string]interface{} {
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

package course

import (
	"os"
	"strconv"
	"strings"

	"backend/entities"

	// Import GORM-related packages.
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Course struct {
}

// Create new subject into platfrom
func (c *Course) Create(subjectID, name, description string) map[string]interface{} {
	db, logs := c.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var maxUID string
	rowU := db.Table("tbl_course_types").Select("max(course_id)").Row()
	rowU.Scan(&maxUID)
	newCID, _ := increaseID(maxUID, "c", 1)

	course := entities.TBL_CourseTypes{
		CourseID:    newCID,
		SubjectID:   subjectID,
		Name:        name,
		Description: description,
	}

	err := db.Create(&course).Error
	if err != nil {
		log := c.errorHandle(err)
		return log
	}

	logs = make(map[string]interface{})
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = ""
	return logs
}

func (c *Course) Read(cid string) map[string]interface{} {
	db, logs := c.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var course []entities.TBL_CourseTypes
	err := db.Find(&course, entities.TBL_CourseTypes{CourseID: cid}).Error
	if err != nil {
		log := c.errorHandle(err)
		return log
	}
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = course
	return log
}

func (c *Course) Update(cid, subjectID, name, description string) map[string]interface{} {
	db, logs := c.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	logs = c.Read(cid)
	if logs["status"] != "1" {
		return logs
	}

	db.Model(&entities.TBL_CourseTypes{}).Where(entities.TBL_CourseTypes{CourseID: cid}).Update(entities.TBL_CourseTypes{SubjectID: subjectID, Name: name, Description: description})

	logs = make(map[string]interface{})
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = ""
	return logs
}

func (c *Course) Delete(cid string) map[string]interface{} {
	db, logs := c.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	err := db.Delete(&entities.TBL_CourseTypes{}, entities.TBL_CourseTypes{CourseID: cid}).Error
	if err != nil {
		log := c.errorHandle(err)
		return log
	}
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = ""
	return log
}

func (c *Course) GetAllCourse() map[string]interface{} {
	db, logs := c.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var courses []entities.TBL_CourseTypes
	err := db.Find(&courses).Error
	if err != nil {
		log := c.errorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = courses

	return log
}

func (c *Course) ReadCourseBySubject(sid string) map[string]interface{} {
	db, logs := c.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var courses []entities.TBL_CourseTypes
	err := db.Find(&courses, entities.TBL_CourseTypes{SubjectID: sid}).Error
	if err != nil {
		log := c.errorHandle(err)
		return log
	}
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = courses
	return log
}

type result_course struct {
	CourseID string
	Name     string
}

func (c *Course) GetAllCourseBySubjectName(subjectName string) map[string]interface{} {
	db, logs := c.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var result []result_course
	err := db.Raw(`	SELECT course_id, name 
					FROM tbl_course_types
					WHERE subject_id = ( SELECT subject_id
										 FROM tbl_subject_types
										 WHERE name = ? ) `, subjectName).Scan(&result).Error
	if err != nil {
		log := c.errorHandle(err)
		return log
	}

	log := make(map[string]interface{})
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

func (c *Course) initDB() (*gorm.DB, map[string]interface{}) {
	var addr = os.Getenv("COCKROACHDB_URL")
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log := c.errorHandle(err)
		return nil, log
	}
	db.LogMode(true)
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = ""

	return db, log
}

func (c *Course) errorHandle(err error) map[string]interface{} {
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

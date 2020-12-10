package course

import (
	"backend/entities"
	"backend/helper"
)

type Course struct {
}

// Create new subject into platfrom
func (c *Course) Create(subjectID, name, description string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var maxUID string
	rowU := db.Table("tbl_course_types").Select("max(course_id)").Row()
	rowU.Scan(&maxUID)
	newCID, _ := helper.IncreaseID(maxUID, "c", 1)

	course := entities.TBL_CourseTypes{
		CourseID:    newCID,
		SubjectID:   subjectID,
		Name:        name,
		Description: description,
	}

	err := db.Create(&course).Error
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

func (c *Course) Read(cid string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var course []entities.TBL_CourseTypes
	err := db.Find(&course, entities.TBL_CourseTypes{CourseID: cid}).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = course
	return log
}

func (c *Course) Update(cid, subjectID, name, description string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	logs = c.Read(cid)
	if logs["status"] != "1" {
		return logs
	}

	err := db.Model(&entities.TBL_CourseTypes{CourseID: cid}).Updates(entities.TBL_CourseTypes{
		SubjectID:   subjectID,
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

func (c *Course) Delete(cid string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	err := db.Delete(&entities.TBL_CourseTypes{}, entities.TBL_CourseTypes{CourseID: cid}).Error
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

func (c *Course) GetAllCourse() map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var courses []entities.TBL_CourseTypes
	err := db.Find(&courses).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = courses

	return log
}

func (c *Course) ReadCourseBySubject(sid string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var courses []entities.TBL_CourseTypes
	err := db.Find(&courses, entities.TBL_CourseTypes{SubjectID: sid}).Error
	if err != nil {
		log := helper.ErrorHandle(err)
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
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var result []result_course
	err := db.Raw(`	SELECT course_id, name 
					FROM tbl_course_types
					WHERE subject_id = ( SELECT subject_id
										 FROM tbl_subject_types
										 WHERE name = ? ) `, subjectName).Scan(&result).Error
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

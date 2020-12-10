package subject

import (
	"backend/entities"
	"backend/helper"
)

type Subject struct {
}

// Create new subject into platfrom
func (s *Subject) Create(name, description string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var maxUID string
	rowU := db.Table("tbl_subject_types").Select("max(subject_id)").Row()
	rowU.Scan(&maxUID)
	newSID, _ := helper.IncreaseID(maxUID, "s", 1)

	subject := entities.TBL_SubjectTypes{
		SubjectID:   newSID,
		Name:        name,
		Description: description,
	}

	err := db.Create(&subject).Error
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

func (s *Subject) Read(sid string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var subjects []entities.TBL_SubjectTypes
	err := db.Find(&subjects, entities.TBL_SubjectTypes{SubjectID: sid}).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = subjects
	return log
}

func (s *Subject) Update(sid, name, description string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	logs = s.Read(sid)
	if logs["status"] != "1" {
		return logs
	}

	err := db.Model(&entities.TBL_SubjectTypes{SubjectID: sid}).Updates(entities.TBL_SubjectTypes{
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

func (s *Subject) Delete(sid string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	err := db.Delete(&entities.TBL_SubjectTypes{}, entities.TBL_SubjectTypes{SubjectID: sid}).Error
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

func (s *Subject) GetAllSubject() map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var subjects []entities.TBL_SubjectTypes
	err := db.Find(&subjects).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = subjects

	return log
}

package subject

import (
	"os"
	"strconv"
	"strings"

	"backend/entities"

	// Import GORM-related packages.
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Subject struct {
}

// Create new subject into platfrom
func (s *Subject) Create(name, description string) map[string]interface{} {
	db, logs := s.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var maxUID string
	rowU := db.Table("tbl_subject_types").Select("max(subject_id)").Row()
	rowU.Scan(&maxUID)
	newSID, _ := increaseID(maxUID, "s", 1)

	subject := entities.TBL_SubjectTypes{
		SubjectID:   newSID,
		Name:        name,
		Description: description,
	}

	err := db.Create(&subject).Error
	if err != nil {
		log := s.errorHandle(err)
		return log
	}

	logs = make(map[string]interface{})
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = ""
	return logs
}

func (s *Subject) Read(sid string) map[string]interface{} {
	db, logs := s.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var subjects []entities.TBL_SubjectTypes
	err := db.Find(&subjects, entities.TBL_SubjectTypes{SubjectID: sid}).Error
	if err != nil {
		log := s.errorHandle(err)
		return log
	}
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = subjects
	return log
}

func (s *Subject) Update(sid, name, description string) map[string]interface{} {
	db, logs := s.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	logs = s.Read(sid)
	if logs["status"] != "1" {
		return logs
	}

	db.Model(&entities.TBL_SubjectTypes{}).Where(entities.TBL_SubjectTypes{SubjectID: sid}).Update(entities.TBL_SubjectTypes{Name: name, Description: description})

	logs = make(map[string]interface{})
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = ""
	return logs
}

func (s *Subject) Delete(sid string) map[string]interface{} {
	db, logs := s.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	err := db.Delete(&entities.TBL_SubjectTypes{}, entities.TBL_SubjectTypes{SubjectID: sid}).Error
	if err != nil {
		log := s.errorHandle(err)
		return log
	}
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = ""
	return log
}

func (s *Subject) GetAllSubject() map[string]interface{} {
	db, logs := s.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var subjects []entities.TBL_SubjectTypes
	err := db.Find(&subjects).Error
	if err != nil {
		log := s.errorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = subjects

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

func (s *Subject) initDB() (*gorm.DB, map[string]interface{}) {
	var addr = os.Getenv("COCKROACHDB_URL")
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log := s.errorHandle(err)
		return nil, log
	}
	db.LogMode(true)
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = ""

	return db, log
}

func (s *Subject) errorHandle(err error) map[string]interface{} {
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

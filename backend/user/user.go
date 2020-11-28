package user

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

type User struct {
}

type result struct {
	UserID   string
	Name     string
	Email    string
	Password string
}

// Login user to authication
func (u *User) Login(email, password string) map[string]interface{} {
	db, logs := u.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var results []result
	err := db.Raw(`	SELECT u.user_id, tu.name, u.email, a.password 
					FROM tbl_users u
					INNER JOIN tbl_auths a
						ON u.user_id = a.user_id
					INNER JOIN tbl_type_users tu
						ON u.user_type = tu.user_type_id
					WHERE u.email = ? `, email).Scan(&results).Error
	if err != nil {
		log := u.errorHandle(err)
		return log
	}

	logs = make(map[string]interface{})

	if len(results) == 1 {
		if results[0].Password == password {
			returnResult := make(map[string]string)
			returnResult["userID"] = results[0].UserID
			returnResult["userType"] = results[0].Name

			logs["status"] = "1"
			logs["msg"] = ""
			logs["result"] = returnResult
			return logs
		} else {
			logs["status"] = "215"
			logs["msg"] = "Wrong password"
			logs["result"] = ""
			return logs
		}
	} else {
		logs["status"] = "215"
		logs["msg"] = "User not found"
		logs["result"] = ""
		return logs
	}

	logs["status"] = "415"
	logs["msg"] = "Something When wrong"
	logs["result"] = ""
	return logs
}

// Register user use it to Register to platform
func (u *User) Register(firstname, familyname, brithday, sex, email, password, userType, educationType string) map[string]interface{} {
	db, logs := u.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var r []string
	err := db.Raw(`	SELECT email
					FROM tbl_users
					WHERE tbl_users.email = ?`, email).Scan(&r).Error
	if err != nil {
		log := u.errorHandle(err)
		return log
	}

	if len(r) == 0 {
		logs := u.Create(firstname, familyname, brithday, sex, email, password, userType, educationType)
		if logs["status"] == "1" {
			return logs
		}
		return logs
	}

	logs = make(map[string]interface{})
	logs["status"] = "215"
	logs["msg"] = "Email already exist"
	logs["result"] = ""
	return logs
}

// ChangePassword for user
func (u *User) ChangePassword(uid, oldPassword, newPassword string) map[string]interface{} {
	db, logs := u.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var auth []entities.TBL_Auth
	err := db.Find(&auth, entities.TBL_Auth{UserID: uid}).Error
	if err != nil {
		log := u.errorHandle(err)
		return log
	}
	log := make(map[string]interface{})

	if oldPassword == auth[0].Password {

		err = db.Model(&entities.TBL_Auth{}).Where(entities.TBL_Auth{UserID: uid}).Update(entities.TBL_Auth{Password: newPassword}).Error
		if err != nil {
			log := u.errorHandle(err)
			return log
		}

		log["status"] = "1"
		log["msg"] = "Password Change Successful"
		log["result"] = ""
		return log
	} else {
		log["status"] = "215"
		log["msg"] = "Wrong Password"
		log["result"] = ""
		return log
	}
}

// Create new user into platfrom
func (u *User) Create(firstname, familyname, brithday, sex, email, password, userType, educationType string) map[string]interface{} {
	db, logs := u.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var maxUID string
	var maxAID string
	rowU := db.Table("tbl_users").Select("max(user_id)").Row()
	rowU.Scan(&maxUID)
	rowA := db.Table("tbl_auth").Select("max(auth_id)").Row()
	rowA.Scan(&maxAID)
	newUID, _ := increaseID(maxUID, "u", 1)
	newAID, _ := increaseID(maxUID, "ay", 2)

	t, err := time.Parse("2006-01-02", brithday)

	if err != nil {
		log := u.errorHandle(err)
		return log
	}

	user := entities.TBL_Users{
		UserID:        newUID,
		Firstname:     firstname,
		Familyname:    familyname,
		Birthday:      t,
		Sex:           sex,
		Email:         email,
		ProfilePic:    newUID + ".jpg",
		UserType:      userType,
		EducationType: educationType,
		Bio:           "",
	}
	auth := entities.TBL_Auth{
		AuthID:   newAID,
		UserID:   newUID,
		Password: password,
	}

	err = db.Create(&user).Error
	if err != nil {
		log := u.errorHandle(err)
		return log
	}
	err = db.Create(&auth).Error
	if err != nil {
		log := u.errorHandle(err)
		return log
	}

	logs = make(map[string]interface{})
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = ""
	return logs
}

func (u *User) Read(uid string) map[string]interface{} {
	db, logs := u.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var users []entities.TBL_Users
	err := db.Find(&users, entities.TBL_Users{UserID: uid}).Error
	if err != nil {
		log := u.errorHandle(err)
		return log
	}
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = users
	return log
}

type result_profile struct {
	Firstname  string
	Familyname string
	Birthday   time.Time
	Email      string
	Sex        string
	ProfilePic string
	EduName    string
	Bio        string
}

func (u *User) ReadProfile(uid string) map[string]interface{} {
	db, log := u.initDB()
	if log["status"] != "1" {
		return log
	}
	defer db.Close()

	var result []result_profile
	err := db.Raw(`	SELECT u.firstname, u.familyname, u.birthday, 
					u.email, u.sex, u.profile_pic, et.name AS edu_name, u.bio
					FROM tbl_users u
					INNER JOIN tbl_education_types et
						ON u.education_type = et.education_type_id
					WHERE u.user_id =  ? `, uid).Scan(&result).Error
	if err != nil {
		log := u.errorHandle(err)
		return log
	}

	log = make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = result
	return log
}

func (u *User) Update(uid, firstname, familyname, birthday, sex, edu, bio string) map[string]interface{} {
	db, logs := u.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	logs = u.Read(uid)
	if logs["status"] != "1" {
		return logs
	}
	t, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		log := u.errorHandle(err)
		return log
	}
	err = db.Model(&entities.TBL_Users{}).Where(entities.TBL_Users{UserID: uid}).Update(entities.TBL_Users{Firstname: firstname, Familyname: familyname, Birthday: t, Sex: sex, EducationType: edu, Bio: bio}).Error
	if err != nil {
		log := u.errorHandle(err)
		return log
	}
	logs = make(map[string]interface{})
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = ""
	return logs
}

func (u *User) Delete(uid string) map[string]interface{} {
	db, logs := u.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	err := db.Delete(&entities.TBL_Users{}, entities.TBL_Users{UserID: uid}).Error
	if err != nil {
		log := u.errorHandle(err)
		return log
	}
	err = db.Delete(&entities.TBL_Auth{}, entities.TBL_Auth{UserID: uid}).Error
	if err != nil {
		log := u.errorHandle(err)
		return log
	}
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = ""
	return log
}

func (u *User) GetAllUser() map[string]interface{} {
	db, logs := u.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var users []entities.TBL_Users
	err := db.Find(&users).Error
	if err != nil {
		log := u.errorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = users

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

func (u *User) initDB() (*gorm.DB, map[string]interface{}) {
	var addr = os.Getenv("COCKROACHDB_URL")
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log := u.errorHandle(err)
		return nil, log
	}
	db.LogMode(true)
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = ""

	return db, log
}

func (u *User) errorHandle(err error) map[string]interface{} {
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

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	// Import GORM-related packages.

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Account is our model, which corresponds to the "accounts" database
// table.

type TBL_Users struct {
	UserID        string `gorm:"primary_key"`
	Firstname     string
	Familyname    string
	Birthday      time.Time
	Email         string
	Sex           string
	ProfilePic    string
	UserType      string
	EducationType string
	Bio           string
}

type TBL_TypeUsers struct {
	UserTypeID string `gorm:"primary_key"`
	Name       string
}

type TBL_EducationTypes struct {
	EducationTypeID string `gorm:"primary_key"`
	Name            string
}

func main() {
	read("u000000002")
	//db.Delete(&TBL_Users{}, "u000000002")
	//printAll(db)
	//create(db)
	//db.Delete(&TBL_Users{}, "user_id = ?", "u000000010")
}

func initDB() *gorm.DB {
	var addr = "postgres://rtae:123qweasdzxc@stylelearn-65t.gcp-asia-southeast1.cockroachlabs.cloud:26257/defaultdb"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)
	return db
}

func create() {
	db := initDB()
	var maxID string
	row := db.Table("tbl_users").Select("max(user_id)").Row()
	row.Scan(&maxID)
	newUID, _ := increaseID(maxID, "u", 1)
	user := TBL_Users{
		UserID:        newUID,
		Firstname:     "Test123qwe",
		Familyname:    "Asdwc",
		Birthday:      time.Date(1999, 8, 5, 0, 0, 0, 0, time.UTC),
		Sex:           "Male",
		Email:         "potae02@gmail.com",
		ProfilePic:    "u000000001.jpg",
		UserType:      "a",
		EducationType: "a",
		Bio:           "Test bio eiming",
	}
	db.Create(&user)
}

func read(uid string) TBL_Users {
	db := initDB()

	var user TBL_Users
	db.Where(&TBL_Users{UserID: uid}).Find(&user)
	fmt.Print(user)
	return user
}

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

func printAll(db *gorm.DB) {
	var users []TBL_Users
	err := db.Find(&users, TBL_Users{Email: "potae02@gmail.com"}).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users[0].UserID)
}

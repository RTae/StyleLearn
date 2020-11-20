package entities

import "time"

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

type TBL_Auth struct {
	AuthID   string `gorm:"primary_key"`
	UserID   string
	Password string
}

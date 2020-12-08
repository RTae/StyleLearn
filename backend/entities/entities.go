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

type TBL_SubjectTypes struct {
	SubjectID   string `gorm:"primary_key"`
	Name        string
	Description string
}

type TBL_CourseTypes struct {
	CourseID    string `gorm:"primary_key"`
	SubjectID   string
	Name        string
	Description string
}

type TBL_LessonTypes struct {
	LessonID    string `gorm:"primary_key"`
	CourseID    string
	Name        string
	Description string
}

type TBL_ProgressLesson struct {
	UserID      string `gorm:"primary_key"`
	LessonID    string `gorm:"primary_key"`
	QuantityDay int64
}

type TBL_Invoice struct {
	InvoiceID  string `gorm:"primary_key"`
	UserID     string
	CreateDate time.Time
	Total      float64
	Detail     string
	Status     *bool
}

type TBL_InvoiceLineItem struct {
	InvoiceID   string `gorm:"primary_key"`
	LessonID    string `gorm:"primary_key"`
	QuantityDay int64
	AmountTotal float64
}

type TBL_Payment struct {
	PaymentID      string `gorm:"primary_key"`
	InvoiceID      string
	UserID         string
	PaymentTypeID  string
	Status         *bool
	CreateDate     time.Time
	DateTransfer   time.Time
	AmountTransfer float64
	Total          float64
	TransferFrom   string
	TransferTo     string
}

package main

import (

	// Import GORM-related packages.

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TBL_ProgressLesson struct {
	UserID      string
	LessonID    string
	QuantityDay int64
}

func main() {
	dsn := "postgres://rtae:123qweasdzxc@stylelearn-65t.gcp-asia-southeast1.cockroachlabs.cloud:26257/defaultdb"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	}

	progressLesson := TBL_ProgressLesson{
		UserID:      "u000000001",
		LessonID:    "l000000001",
		QuantityDay: 4,
	}

	err = db.Clauses(clause.OnConflict{DoNothing: true}).Create(&progressLesson).Error
	if err != nil {
	}
}

package helper

import (

	// Import GORM-related packages.

	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func InitDB() (*gorm.DB, map[string]interface{}) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 400 * time.Millisecond, // Slow SQL threshold
			LogLevel:      logger.Info,            // Log level
			Colorful:      true,                   // Disable color
		},
	)
	dsn := os.Getenv("COCKROACHDB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log := ErrorHandle(err)
		return nil, log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = ""

	return db, log
}

func ErrorHandle(err error) map[string]interface{} {
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

func IncreaseID(id, name string, length int) (string, error) {
	digit, err := strconv.Atoi(id[length:])
	if err != nil {
		return name, err
	}
	digit++
	s := strconv.Itoa(digit)

	newID := name + strings.Repeat("0", 10-length-len(s)) + s
	return newID, nil
}

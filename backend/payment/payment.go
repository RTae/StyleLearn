package payment

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

// Invoice function
type Payment struct {
}

// Create Create payment
func (p *Payment) Create(invoiceID, userID, paymentTypeID, status, createDate, dateTransfer, amountTranfer, total, transferFrom, tranfrerTo string) map[string]interface{} {
	db, logs := p.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var maxUID string
	rowU := db.Table("tbl_payments").Select("max(payment_id)").Row()
	rowU.Scan(&maxUID)
	newPID, _ := increaseID(maxUID, "p", 1)

	t, err := time.Parse("2006-01-02", createDate)
	if err != nil {
		log := p.errorHandle(err)
		return log
	}

	t2, err := time.Parse("2006-01-02T15:04:05Z07:00", dateTransfer)
	if err != nil {
		log := p.errorHandle(err)
		return log
	}

	totalFloat, err := strconv.ParseFloat(total, 64)
	if err != nil {
		log := p.errorHandle(err)
		return log
	}

	amountFloat, err := strconv.ParseFloat(amountTranfer, 64)
	if err != nil {
		log := p.errorHandle(err)
		return log
	}

	statusBool, err := strconv.ParseBool(status)
	if err != nil {
		log := p.errorHandle(err)
		return log
	}

	Payment := entities.TBL_Payment{
		PaymentID:      newPID,
		InvoiceID:      invoiceID,
		UserID:         userID,
		PaymentTypeID:  paymentTypeID,
		Status:         statusBool,
		CreateDate:     t,
		DateTransfer:   t2,
		AmountTransfer: amountFloat,
		Total:          totalFloat,
		TransferFrom:   transferFrom,
		TransferTo:     tranfrerTo,
	}

	err = db.Create(&Payment).Error
	if err != nil {
		log := p.errorHandle(err)
		return log
	}

	logs = make(map[string]interface{})
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = ""
	return logs
}

// Read read payment
func (p *Payment) Read(pid string) map[string]interface{} {
	db, logs := p.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var Payment []entities.TBL_Payment
	err := db.Find(&Payment, entities.TBL_Payment{PaymentID: pid}).Error
	if err != nil {
		log := p.errorHandle(err)
		return log
	}
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = Payment
	return log
}

//Update update payment
func (p *Payment) Update(pid, invoiceID, userID, paymentTypeID, status, createDate, dateTransfer, amountTranfer, total, transferFrom, tranfrerTo string) map[string]interface{} {
	db, logs := p.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	logs = p.Read(pid)
	if logs["status"] != "1" {
		return logs
	}

	t, err := time.Parse("2006-01-02", createDate)
	if err != nil {
		log := p.errorHandle(err)
		return log
	}

	t2, err := time.Parse("2006-01-02T15:04:05Z07:00", dateTransfer)
	if err != nil {
		log := p.errorHandle(err)
		return log
	}

	totalFloat, err := strconv.ParseFloat(total, 64)
	if err != nil {
		log := p.errorHandle(err)
		return log
	}

	amountFloat, err := strconv.ParseFloat(amountTranfer, 64)
	if err != nil {
		log := p.errorHandle(err)
		return log
	}

	statusBool, err := strconv.ParseBool(status)
	if err != nil {
		log := p.errorHandle(err)
		return log
	}

	err = db.Model(&entities.TBL_Payment{}).Where(entities.TBL_Payment{PaymentID: pid}).Update(entities.TBL_Payment{
		InvoiceID:      invoiceID,
		UserID:         userID,
		PaymentTypeID:  paymentTypeID,
		Status:         statusBool,
		CreateDate:     t,
		DateTransfer:   t2,
		AmountTransfer: amountFloat,
		Total:          totalFloat,
		TransferFrom:   transferFrom,
		TransferTo:     tranfrerTo,
	}).Error

	if err != nil {
		log := p.errorHandle(err)
		return log
	}

	logs = make(map[string]interface{})
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = ""
	return logs
}

// Delete delete payment
func (p *Payment) Delete(pid string) map[string]interface{} {
	db, logs := p.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	err := db.Where("payment_id = ?", pid).Delete(&entities.TBL_Payment{}).Error
	if err != nil {
		log := p.errorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = ""
	return log
}

// GetAllPayment get all payment
func (p *Payment) GetAllPayment() map[string]interface{} {
	db, logs := p.initDB()
	if logs["status"] != "1" {
		return logs
	}
	defer db.Close()

	var Payments []entities.TBL_Payment
	err := db.Find(&Payments).Error
	if err != nil {
		log := p.errorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = Payments

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

func (p *Payment) initDB() (*gorm.DB, map[string]interface{}) {
	var addr = os.Getenv("COCKROACHDB_URL")
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log := p.errorHandle(err)
		return nil, log
	}
	db.LogMode(true)
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = ""

	return db, log
}

func (p *Payment) errorHandle(err error) map[string]interface{} {
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

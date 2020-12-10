package payment

import (
	"strconv"
	"time"

	"backend/entities"
	"backend/helper"
)

// Invoice function
type Payment struct {
}

// Create Create payment
func (p *Payment) Create(invoiceID, userID, paymentTypeID, status, createDate, dateTransfer, amountTranfer, total, transferFrom, tranfrerTo string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var maxUID string
	rowU := db.Table("tbl_payments").Select("max(payment_id)").Row()
	rowU.Scan(&maxUID)
	newPID, _ := helper.IncreaseID(maxUID, "p", 1)

	t, err := time.Parse("2006-01-02", createDate)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	t2, err := time.Parse("2006-01-02T15:04:05Z07:00", dateTransfer)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	totalFloat, err := strconv.ParseFloat(total, 64)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	amountFloat, err := strconv.ParseFloat(amountTranfer, 64)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	statusBool, err := strconv.ParseBool(status)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	Payment := entities.TBL_Payment{
		PaymentID:      newPID,
		InvoiceID:      invoiceID,
		UserID:         userID,
		PaymentTypeID:  paymentTypeID,
		Status:         &statusBool,
		CreateDate:     t,
		DateTransfer:   t2,
		AmountTransfer: amountFloat,
		Total:          totalFloat,
		TransferFrom:   transferFrom,
		TransferTo:     tranfrerTo,
	}

	err = db.Create(&Payment).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	logs = make(map[string]interface{})
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = newPID
	return logs
}

// Read read payment
func (p *Payment) Read(pid string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var Payment []entities.TBL_Payment
	err := db.Find(&Payment, entities.TBL_Payment{PaymentID: pid}).Error
	if err != nil {
		log := helper.ErrorHandle(err)
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
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	logs = p.Read(pid)
	if logs["status"] != "1" {
		return logs
	}

	t, err := time.Parse("2006-01-02", createDate)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	t2, err := time.Parse("2006-01-02T15:04:05Z07:00", dateTransfer)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	totalFloat, err := strconv.ParseFloat(total, 64)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	amountFloat, err := strconv.ParseFloat(amountTranfer, 64)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	statusBool, err := strconv.ParseBool(status)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	err = db.Model(&entities.TBL_Payment{PaymentID: pid}).Updates(entities.TBL_Payment{
		InvoiceID:      invoiceID,
		UserID:         userID,
		PaymentTypeID:  paymentTypeID,
		Status:         &statusBool,
		CreateDate:     t,
		DateTransfer:   t2,
		AmountTransfer: amountFloat,
		Total:          totalFloat,
		TransferFrom:   transferFrom,
		TransferTo:     tranfrerTo,
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

// Delete delete payment
func (p *Payment) Delete(pid string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	err := db.Where("payment_id = ?", pid).Delete(&entities.TBL_Payment{}).Error
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

// GetAllPayment get all payment
func (p *Payment) GetAllPayment() map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var Payments []entities.TBL_Payment
	err := db.Find(&Payments).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = Payments

	return log
}

package invoice

import (
	"strconv"
	"time"

	"backend/entities"
	"backend/helper"
)

// Invoice function
type Invoice struct {
}

// Create Invoice
func (i *Invoice) Create(userID, createDate, total, detail, status string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var maxUID string
	rowU := db.Table("tbl_invoices").Select("max(invoice_id)").Row()
	rowU.Scan(&maxUID)
	newIID, _ := helper.IncreaseID(maxUID, "i", 1)

	t, err := time.Parse("2006-01-02", createDate)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	total_float, err := strconv.ParseFloat(total, 64)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	status_bool, err := strconv.ParseBool(status)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	Invoice := entities.TBL_Invoice{
		InvoiceID:  newIID,
		UserID:     userID,
		CreateDate: t,
		Total:      total_float,
		Detail:     detail,
		Status:     &status_bool,
	}

	err = db.Create(&Invoice).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	logs = make(map[string]interface{})
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = newIID
	return logs
}

func (i *Invoice) Read(iid string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var Invoices []entities.TBL_Invoice
	err := db.Find(&Invoices, entities.TBL_Invoice{InvoiceID: iid}).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}
	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = Invoices
	return log
}

func (i *Invoice) Update(iid, userID, createDate, total, detail, status string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	logs = i.Read(iid)
	if logs["status"] != "1" {
		return logs
	}

	t, err := time.Parse("2006-01-02", createDate)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	total_float, err := strconv.ParseFloat(total, 64)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	status_bool, err := strconv.ParseBool(status)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	err = db.Model(&entities.TBL_Invoice{InvoiceID: iid}).Updates(entities.TBL_Invoice{
		UserID:     userID,
		CreateDate: t,
		Total:      total_float,
		Detail:     detail,
		Status:     &status_bool,
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

func (i *Invoice) Delete(iid string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	err := db.Where("invoice_id = ?", iid).Delete(&entities.TBL_Invoice{}).Error
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

func (i *Invoice) GetAllInvoice() map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var Invoices []entities.TBL_Invoice
	err := db.Find(&Invoices).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = Invoices

	return log
}

func (i *Invoice) CancelInvoice(invoiceID string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	err := db.Where("invoice_id = ?", invoiceID).Delete(&entities.TBL_InvoiceLineItem{}).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	err = db.Where("invoice_id = ?", invoiceID).Delete(&entities.TBL_Invoice{}).Error
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

type result_invoiceLine struct {
	LessonID    string
	QuantityDay string
}

func (i *Invoice) GetInvoiceLineItemByInvoiceID(invoiceID string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var result []result_invoiceLine
	err := db.Raw(`	SELECT ilt.lesson_id, ilt.quantity_day
					FROM tbl_invoices i
					INNER JOIN tbl_invoice_line_items ilt
					ON i.invoice_id = ilt.invoice_id
					WHERE i.invoice_id = ?`, invoiceID).Scan(&result).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	logs = make(map[string]interface{})
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = result
	return logs

}

func (i *Invoice) UpdateStatusInvoice(invoiceID, status string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	status_bool, err := strconv.ParseBool(status)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	err = db.Model(&entities.TBL_Invoice{InvoiceID: invoiceID}).Updates(entities.TBL_Invoice{
		Status: &status_bool,
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

// AddItemToLineItem Add item to line invoice
func (i *Invoice) AddItemToLineItem(invoiceID, lessonID, quantityDay, amountTotal string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	quantityDay_int, err := strconv.ParseInt(quantityDay, 10, 64)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	amount_float, err := strconv.ParseFloat(amountTotal, 64)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	InvoiceLineItem := entities.TBL_InvoiceLineItem{
		InvoiceID:   invoiceID,
		LessonID:    lessonID,
		QuantityDay: quantityDay_int,
		AmountTotal: amount_float,
	}

	err = db.Create(&InvoiceLineItem).Error
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

// ReadItemLineItem Read item in line item

func (i *Invoice) ReadItemLineItem(invoiceID, lessonID string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var InvoiceLineItem []entities.TBL_InvoiceLineItem
	err := db.Find(&InvoiceLineItem, entities.TBL_InvoiceLineItem{InvoiceID: invoiceID, LessonID: lessonID}).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = InvoiceLineItem

	return log

}

//  UpdateItemLineItem updateitem in line item
func (i *Invoice) UpdateItemLineItem(invoiceID, lessonID, quantityDay, amountTotal string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	quantityDay_int, err := strconv.ParseInt(quantityDay, 10, 64)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	amount_float, err := strconv.ParseFloat(amountTotal, 64)
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	err = db.Model(&entities.TBL_InvoiceLineItem{InvoiceID: invoiceID, LessonID: lessonID}).Updates(entities.TBL_InvoiceLineItem{
		QuantityDay: quantityDay_int,
		AmountTotal: amount_float,
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

// DeleteItemLineItem Delete line item
func (i *Invoice) DeleteItemLineItem(invoiceID, lessonID string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	err := db.Where("invoice_id = ? AND lesson_id = ?", invoiceID, lessonID).Delete(&entities.TBL_InvoiceLineItem{}).Error

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

type result_unpaid struct {
	Invoice_id string
	Total      string
}

func (i *Invoice) GetUnpaidInvoice(uid string) map[string]interface{} {
	db, logs := helper.InitDB()
	if logs["status"] != "1" {
		return logs
	}

	var result []result_unpaid
	err := db.Raw(`	SELECT invoice_id, total
					FROM tbl_invoices
					WHERE user_id = ? AND status = false`, uid).Scan(&result).Error
	if err != nil {
		log := helper.ErrorHandle(err)
		return log
	}

	log := make(map[string]interface{})
	log["status"] = "1"
	log["msg"] = ""
	log["result"] = result

	return log
}

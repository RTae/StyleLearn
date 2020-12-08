package handler

import (
	"backend/helper"
	"backend/invoice"
	"encoding/json"
	"fmt"
	"net/http"
)

//Handler Login Auth CorchroshDB
func Handler(w http.ResponseWriter, r *http.Request) {
	helper.SetupResponse(&w, r)
	i := invoice.Invoice{}
	if (*r).Method == "OPTIONS" {
		return
	}
	if (*r).Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		userID := r.FormValue("userID")
		invoiceID := r.FormValue("invoiceID")
		lessonID := r.FormValue("lessonID")
		mode := r.FormValue("mode")

		if mode == "1" {
			logs := i.Read(invoiceID)
			json.NewEncoder(w).Encode(logs)
		} else if mode == "2" {
			logs := i.ReadItemLineItem(invoiceID, lessonID)
			json.NewEncoder(w).Encode(logs)
		} else if mode == "3" {
			logs := i.GetUnpaidInvoice(userID)
			json.NewEncoder(w).Encode(logs)
		} else if mode == "4" {
			logs := i.GetInvoiceLineItemByInvoiceID(invoiceID)
			json.NewEncoder(w).Encode(logs)
		} else {
			logs := i.GetAllInvoice()
			json.NewEncoder(w).Encode(logs)
		}
	} else if (*r).Method == "POST" {
		w.Header().Set("Content-Type", "application/json")

		// Invoice
		userID := r.FormValue("userID")
		createDate := r.FormValue("createDate")
		total := r.FormValue("total")
		detail := r.FormValue("detail")
		status := r.FormValue("status")

		// Line item
		invoiceID := r.FormValue("invoiceID")
		lessonID := r.FormValue("lessonID")
		quantityDay := r.FormValue("quantityDay")
		amountTotal := r.FormValue("amountTotal")

		mode := r.FormValue("mode")

		if mode == "1" {
			logs := i.AddItemToLineItem(invoiceID, lessonID, quantityDay, amountTotal)
			json.NewEncoder(w).Encode(logs)
		} else {
			logs := i.Create(userID, createDate, total, detail, status)
			json.NewEncoder(w).Encode(logs)
		}

	} else if (*r).Method == "PUT" {
		w.Header().Set("Content-Type", "application/json")

		// Invoice
		invoiceID := r.FormValue("invoiceID")
		userID := r.FormValue("userID")
		createDate := r.FormValue("createDate")
		total := r.FormValue("total")
		detail := r.FormValue("detail")
		status := r.FormValue("status")

		// Line item
		lessonID := r.FormValue("lessonID")
		quantityDay := r.FormValue("quantityDay")
		amountTotal := r.FormValue("amountTotal")

		mode := r.FormValue("mode")

		if mode == "1" {
			logs := i.UpdateItemLineItem(invoiceID, lessonID, quantityDay, amountTotal)
			json.NewEncoder(w).Encode(logs)
		} else if mode == "2" {
			logs := i.UpdateStatusInvoice(invoiceID, status)
			json.NewEncoder(w).Encode(logs)
		} else {
			logs := i.Update(invoiceID, userID, createDate, total, detail, status)
			json.NewEncoder(w).Encode(logs)
		}

	} else if (*r).Method == "DELETE" {
		w.Header().Set("Content-Type", "application/json")

		invoiceID := r.FormValue("invoiceID")
		lessonID := r.FormValue("lessonID")

		mode := r.FormValue("mode")
		if mode == "1" {
			logs := i.DeleteItemLineItem(invoiceID, lessonID)
			json.NewEncoder(w).Encode(logs)
		} else if mode == "2" {
			logs := i.CancelInvoice(invoiceID)
			json.NewEncoder(w).Encode(logs)
		} else {
			logs := i.Delete(invoiceID)
			json.NewEncoder(w).Encode(logs)
		}

	} else {
		fmt.Fprintf(w, "Please use get medthod")
	}
}

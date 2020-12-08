package handler

import (
	"backend/helper"
	"backend/payment"
	"encoding/json"
	"fmt"
	"net/http"
)

//Handler Login Auth CorchroshDB
func Handler(w http.ResponseWriter, r *http.Request) {
	helper.SetupResponse(&w, r)
	p := payment.Payment{}
	if (*r).Method == "OPTIONS" {
		return
	}
	if (*r).Method == "GET" {
		w.Header().Set("Content-Type", "application/json")

		paymentID := r.FormValue("paymentID")

		mode := r.FormValue("mode")

		if mode == "1" {
			logs := p.Read(paymentID)
			json.NewEncoder(w).Encode(logs)
		} else {
			logs := p.GetAllPayment()
			json.NewEncoder(w).Encode(logs)
		}
	} else if (*r).Method == "POST" {
		w.Header().Set("Content-Type", "application/json")

		// Payment
		invoiceID := r.FormValue("invoiceID")
		userID := r.FormValue("userID")
		paymentTypeID := r.FormValue("paymentTypeID")
		status := r.FormValue("status")
		createDate := r.FormValue("createDate")
		tranferDate := r.FormValue("tranferDate")
		amountTranfer := r.FormValue("amountTranfer")
		total := r.FormValue("total")
		transferForm := r.FormValue("tranferForm")
		transferTo := r.FormValue("tranferTo")

		logs := p.Create(invoiceID, userID, paymentTypeID, status, createDate, tranferDate, amountTranfer, total, transferForm, transferTo)
		json.NewEncoder(w).Encode(logs)

	} else if (*r).Method == "PUT" {
		w.Header().Set("Content-Type", "application/json")

		// Payment
		paymentID := r.FormValue("paymentID")
		invoiceID := r.FormValue("invoiceID")
		userID := r.FormValue("userID")
		paymentTypeID := r.FormValue("paymentTypeID")
		status := r.FormValue("status")
		createDate := r.FormValue("createDate")
		tranferDate := r.FormValue("tranferDate")
		amountTranfer := r.FormValue("amountTranfer")
		total := r.FormValue("total")
		transferForm := r.FormValue("tranferForm")
		transferTo := r.FormValue("tranferTo")

		logs := p.Update(paymentID, invoiceID, userID, paymentTypeID, status, createDate, tranferDate, amountTranfer, total, transferForm, transferTo)
		json.NewEncoder(w).Encode(logs)

	} else if (*r).Method == "DELETE" {
		w.Header().Set("Content-Type", "application/json")

		paymentID := r.FormValue("paymentID")

		logs := p.Delete(paymentID)
		json.NewEncoder(w).Encode(logs)
	} else {
		fmt.Fprintf(w, "Please use get medthod")
	}
}

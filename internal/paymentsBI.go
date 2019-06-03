package internal

import (
	"fmt"
	"net/http"
	"../model"
	"../utils"
	"strconv"
)

// Mock Payment
var MakePayment = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MakePayment")

	res_id := utils.GetParam("res_id", w, r)
	if res_id == "" {
		return
	}
	// Check if the payment is already made
	query := "res_id = " + res_id
	payments, err := model.QueryPayment(query)
	if err != nil {
		utils.HTTPResponse(w, http.StatusInternalServerError, utils.JSONMarshal(err))
		return
	}
	if len(payments) == 0 {
		msg := "Payment for reservation (res_id=" + res_id + ") not found"
		utils.HTTPResponse(w, http.StatusBadRequest, utils.FormatErrorMessage(msg))
		return
	}
	if payments[0].Paid {
		msg := "Payment for reservation (res_id=" + res_id + ") already done"
		utils.HTTPResponse(w, http.StatusBadRequest, utils.FormatErrorMessage(msg))
		return
	}

	// Update paid field
	payments[0].Paid = true
	err = model.UpdatePayment(payments[0])
	if err != nil {
		utils.HTTPResponse(w, http.StatusInternalServerError, utils.JSONMarshal(err))
	} else {
		msg := "Payment for reservation (res_id=" + res_id + ") done"
		utils.HTTPResponse(w, http.StatusOK, utils.FormatSuccessMessage(msg))
	}
}

var GetPayments = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetPayments")

	id := 0
	_id := r.URL.Query().Get("id")
	if _id != "" {
		id, _ = strconv.Atoi(_id)
	}

	payments, err := model.GetPayments(uint(id))
	if err != nil {
		utils.HTTPResponse(w, http.StatusInternalServerError, utils.JSONMarshal(err))
	} else {
		utils.HTTPResponse(w, http.StatusOK, utils.JSONMarshal(payments))
	}
}

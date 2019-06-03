package internal

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"../model"
	"../utils"
)

var CreateReservation = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateReservation")
	reservation := &model.Reservation{}

	id := utils.GetParam("id", w, r)
	if id == "" {
		return
	}
	scooterId,_ := strconv.Atoi(id)
	reservation.ScooterID = uint(scooterId)

	// Check if scooter is already reserved
	scooters, err := model.GetScooters(uint(scooterId))
	if err != nil {
		utils.HTTPResponse(w, http.StatusInternalServerError, utils.JSONMarshal(err))
		return
	}
	if len(scooters) == 0 {
		msg := "Scooter (id=" + id + ") not found"
		utils.HTTPResponse(w, http.StatusBadRequest, utils.FormatErrorMessage(msg))
		return
	}

	if scooters[0].Reserved == true {
		msg := "Scooter (id=" + id + ") is already reserved"
		utils.HTTPResponse(w, http.StatusBadRequest, utils.FormatErrorMessage(msg))
		return
	}

	// Create reservation
	err = reservation.Create()
	if err != nil {
		utils.HTTPResponse(w, http.StatusInternalServerError, utils.JSONMarshal(err))
		return
	}

	// Mark the scooter as reserved
	scooters[0].Reserved = true
	err = model.UpdateScooter(scooters[0])
	if err != nil {
		utils.HTTPResponse(w, http.StatusInternalServerError, utils.JSONMarshal(err))
		return
	}
	utils.HTTPResponse(w, http.StatusOK, utils.JSONMarshal(reservation))
}

var EndReservation = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EndReservation")

	res_id := utils.GetParam("res_id", w, r)
	if res_id == "" {
		return
	}
	lat := utils.GetParam("lat", w, r)
	if lat == "" {
		return
	}
	lng := utils.GetParam("lng", w, r)
	if lng == "" {
		return
	}
	resId,_ := strconv.Atoi(res_id)
	latitude, _ := strconv.ParseFloat(lat, 64)
	longitude, _ := strconv.ParseFloat(lng, 64)

	reservations, err := model.GetReservations(uint(resId))
	if err != nil {
		utils.HTTPResponse(w, http.StatusInternalServerError, utils.JSONMarshal(err))
		return
	}
	if len(reservations) == 0 {
		msg := "Reservation (id=" + res_id + ") not found"
		utils.HTTPResponse(w, http.StatusBadRequest, utils.FormatErrorMessage(msg))
		return
	}
	// Check if reservation is already ended
	query := "res_id = " + res_id
	payments, err := model.QueryPayment(query)
	if err != nil {
		utils.HTTPResponse(w, http.StatusInternalServerError, utils.JSONMarshal(err))
		return
	}
	if len(payments) == 1 {
		msg := "Reservation (id=" + res_id + ") is already ended"
		utils.HTTPResponse(w, http.StatusBadRequest, utils.FormatErrorMessage(msg))
		return
	}

	scooterId := reservations[0].ScooterID
	scooters, err := model.GetScooters(scooterId)
	if err != nil {
		utils.HTTPResponse(w, http.StatusInternalServerError, utils.JSONMarshal(err))
		return
	}
	if len(scooters) == 0 {
		msg := "Scooter (id=" + string(scooterId) + ") not found"
		utils.HTTPResponse(w, http.StatusInternalServerError, utils.FormatErrorMessage(msg))
		return
	}
	// Calculate the payment from total distance travelled
	dist := Distance(scooters[0].Latitude, scooters[0].Longitude, latitude, longitude)
	milePerMeter := 0.000621371
	amt := math.Ceil(dist * milePerMeter * model.Rate * 100) / 100
	fmt.Println("Scooter(id=",scooterId,"), distance=",dist,", amount=", amt)

	// Mark the scooter available, update co-ordinates
	scooters[0].Latitude = latitude
	scooters[0].Longitude = longitude
	scooters[0].Reserved = false

	err = model.SaveScooter(scooters[0])
	if err != nil {
		utils.HTTPResponse(w, http.StatusInternalServerError, utils.JSONMarshal(err))
		return
	}
	// Create payment
	payment := &model.Payment{}
	payment.Amount = amt
	payment.ResID = uint(resId)

	err = payment.Create()
	if err != nil {
		utils.HTTPResponse(w, http.StatusInternalServerError, utils.JSONMarshal(err))
	} else {
		utils.HTTPResponse(w, http.StatusOK, utils.JSONMarshal(payment))
	}
}

var GetReservations = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetReservations")

	id := 0
	_id := r.URL.Query().Get("id")
	if _id != "" {
		id, _ = strconv.Atoi(_id)
	}

	reservations, err := model.GetReservations(uint(id))
	if err != nil {
		utils.HTTPResponse(w, http.StatusInternalServerError, utils.JSONMarshal(err))
	} else {
		utils.HTTPResponse(w, http.StatusOK, utils.JSONMarshal(reservations))
	}
}

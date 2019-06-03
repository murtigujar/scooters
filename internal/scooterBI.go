package internal

import (
	"../model"
	"../utils"
	"fmt"
	"net/http"
	"strconv"
)

var AddScooter = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddScooter")
	scooter := &model.Scooter{}

	lat := utils.GetParam("lat", w, r)
	if lat == "" {
		return
	}
	lng := utils.GetParam("lng", w, r)
	if lng == "" {
		return
	}
	scooter.Latitude, _ = strconv.ParseFloat(lat, 64)
	scooter.Longitude, _ = strconv.ParseFloat(lng, 64)

	err := scooter.Create()
	if err != nil {
		utils.HTTPResponse(w, http.StatusInternalServerError, utils.JSONMarshal(err))
	} else {
		utils.HTTPResponse(w, http.StatusOK, utils.JSONMarshal(scooter))
	}
}

var GetAvailableScootersWithinRadius = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAvailableScooters")

	lat := utils.GetParam("lat", w, r)
	if lat == "" {
		return
	}
	lng := utils.GetParam("lng", w, r)
	if lng == "" {
		return
	}
	rad := utils.GetParam("radius", w, r)
	if rad == "" {
		return
	}
	latitude, _ := strconv.ParseFloat(lat, 64)
	longitude, _ := strconv.ParseFloat(lng, 64)
	radius, _ := strconv.ParseFloat(rad, 64)

	availScooters, err := model.GetAvailableScooters()
	if err != nil {
		utils.HTTPResponse(w, http.StatusInternalServerError, utils.JSONMarshal(err))
		return
	}
	// Filter the ones which are in the given radius
	scooters := []*model.Scooter{}
	for _, s := range availScooters {
		if Distance(s.Latitude, s.Longitude, latitude, longitude) <= radius {
			scooters = append(scooters, s)
		}
	}
	utils.HTTPResponse(w, http.StatusOK, utils.JSONMarshal(scooters))
}

var GetScooters = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetScooters")

	id := 0
	_id := r.URL.Query().Get("id")
	if _id != "" {
		id, _ = strconv.Atoi(_id)
	}

	scooters, err := model.GetScooters(uint(id))
	if err != nil {
		utils.HTTPResponse(w, http.StatusInternalServerError, utils.JSONMarshal(err))
	} else {
		utils.HTTPResponse(w, http.StatusOK, utils.JSONMarshal(scooters))
	}
}

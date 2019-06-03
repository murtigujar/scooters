package main

import (
	".glide/cache/src/https-github.com-gorilla-mux"
	"fmt"
	"net/http"
	"os"
	"../internal"
)

func handleRequests() {

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/scooters/reservations", internal.GetReservations).Methods("GET")
	router.HandleFunc("/api/v1/scooters/reserve", internal.CreateReservation).Methods("POST")
	router.HandleFunc("/api/v1/scooters/return", internal.EndReservation).Methods("POST")

	router.HandleFunc("/api/v1/scooters/payments", internal.GetPayments).Methods("GET")
	router.HandleFunc("/api/v1/scooters/pay", internal.MakePayment).Methods("POST")

	router.HandleFunc("/api/v1/scooters/available", internal.GetAvailableScootersWithinRadius).Methods("GET")
	router.HandleFunc("/api/v1/scooters", internal.AddScooter).Methods("POST")
	router.HandleFunc("/api/v1/scooters", internal.GetScooters).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8420" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}


func main()  {
	handleRequests()
}

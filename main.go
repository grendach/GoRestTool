package main

import (
	"app/helper"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// GetCars is a HTTP server function. Inside of it we load our CSV file with the cars.
func GetCars(w http.ResponseWriter, r *http.Request) {
	cars := helper.LoadCarsCsv()
	json.NewEncoder(w).Encode(cars)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/cars", GetCars).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

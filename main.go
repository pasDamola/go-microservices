package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name string `json:"full_name"`
	City string `json:"city"`
	ZipCode string `json:"zip_code"`
}


func main() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer {
		{"Oyincde", "Abuja", "9000001"},
		{"Sandra", "Abia", "9000001"},
	}

	// set the application header
	w.Header().Add("Content-Type", "application/json")

	// encode response as json
	json.NewEncoder(w).Encode(customers)
}
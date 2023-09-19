package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name string `json:"full_name" xml:"name"`
	City string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipcode"`
}


func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer {
		{"Oyincde", "Abuja", "9000001"},
		{"Sandra", "Abia", "9000001"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {

		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)

	} else {

		// set the application header
		w.Header().Add("Content-Type", "application/json")
		// encode response as json
		json.NewEncoder(w).Encode(customers)
	}

}


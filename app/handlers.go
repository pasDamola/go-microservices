package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/pasDamola/go-microservices/service"
)

type Customer struct {
	Name string `json:"full_name" xml:"name"`
	City string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}


func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

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


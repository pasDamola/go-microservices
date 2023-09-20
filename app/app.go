package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pasDamola/go-microservices/domain"
	"github.com/pasDamola/go-microservices/service"
)

func Start() {

	//mux := http.NewServeMux()
	router := mux.NewRouter()

	// wiring the complete application
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
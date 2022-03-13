package app

import (
	"log"
	"net/http"

	"github.com/cspark0610/go-banking-rest-api/domain"
	"github.com/cspark0610/go-banking-rest-api/service"
	"github.com/gorilla/mux"
)

func Start() {
	// DEFINE DEMULTIPLEXER MUX ROUTER
	router := mux.NewRouter()
	// wiring up the routes
	// ch := CustomerHandlers{service.NewCustomerService(domain.CustomerRepositoryStub{})}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// DEFINE ROUTES ENDPOINTS USING DEMULTIPLEXER MUX
	// 1.GET METHODS
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// 2.POST METHODS
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	// start server
	log.Fatal(http.ListenAndServe(":8000",router))
}
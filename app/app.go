package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// DEFINE DEMULTIPLEXER MUX
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	// DEFINE ROUTES ENDPOINTS USING DEMULTIPLEXER MUX
	// GET METHODS
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getAllCustomerById).Methods(http.MethodGet)

	// POST METHODS
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	// start server
	log.Fatal(http.ListenAndServe(":8000",router))
}
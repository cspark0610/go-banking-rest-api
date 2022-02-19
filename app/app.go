package main

import (
	"log"
	"net/http"
)

func Start() {
	// DEFINE ROUTES ENDPOINTS
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// start server
	log.Fatal(http.ListenAndServe(":8000", nil))
}
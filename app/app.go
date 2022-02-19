package app

import (
	"log"
	"net/http"
)

func Start() {
	// DEFINE DEMULTIPLEXER MUX
	mux := http.NewServeMux()

	// DEFINE ROUTES ENDPOINTS USING DEMULTIPLEXER MUX
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	// start server
	log.Fatal(http.ListenAndServe(":8000", mux))
}
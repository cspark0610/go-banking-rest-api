package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// struct are ts classes dtos
type Customer struct {
	// tags are used to define the json key response use backtics
	Name     string  `json:"full_name" xml:"full_name"`
	City     string  `json:"city"      xml:"city"`
	ZipCode  string  `json:"zip_code"  xml:"zip_code"`
}

// 1. ROUTE HANDLERS
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request)  {
	//initalize a slice of customers, a slice is an array objects
	customers := []Customer{
		{ Name: "John", City: "New York", ZipCode: "10001"},
		{ Name: "John", City: "New Jersey", ZipCode: "10002"},
		{ Name: "John", City: "New Delhi", ZipCode: "10003"},
		{ Name: "John", City: "New Hampshire", ZipCode: "10004"},
	}
	// as content-type text/plain; charset=utf-8, debo setear el Header.content-type a 'application/json' o 'application/xml'

	// si seteo desde el cliente el content-type a 'application/xml' se manda en ese formato
	// caso contratio se manda en json
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}


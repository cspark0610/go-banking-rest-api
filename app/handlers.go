package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/cspark0610/go-banking-rest-api/service"
	"github.com/gorilla/mux"
)

// struct are ts classes dtos
type Customer struct {
	// tags are used to define the json key response use backtics
	Name     string  `json:"full_name" xml:"full_name"`
	City     string  `json:"city"      xml:"city"`
	ZipCode  string  `json:"zip_code"  xml:"zip_code"`
}

//CustomerHandlers struct va a tener dependencia del service segun el diagrama
type CustomerHandlers struct {
	service service.CustomerService
}

// 1. ROUTE HANDLERS

func (ch *CustomerHandlers) getAllCustomers(res http.ResponseWriter, req *http.Request)  {

	customers, _ := ch.service.GetAllCustomers()
	
	// as content-type text/plain; charset=utf-8, debo setear el Header.content-type a 'application/json' o 'application/xml'
	// si seteo desde el cliente el content-type a 'application/xml' se manda en ese formato
	// caso contrario se manda en json
	if req.Header.Get("Content-Type") == "application/xml" {
		res.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(res).Encode(customers)
	} else {
		res.Header().Add("Content-Type", "application/json")
		json.NewEncoder(res).Encode(customers)
	}
}

func (ch *CustomerHandlers) getCustomer(res http.ResponseWriter, req *http.Request)  {
	// llamamos a la funcion de mux.Vars para obtener el parametro que se pasa en la url
	vars:=  mux.Vars(req)
	id := vars["customer_id"]
	customer, err :=  ch.service.GetCustomer(id)
	if err != nil {
		res.WriteHeader(err.Code)
		fmt.Fprint(res, err.Message)
	} else {
		// caso exitoso se manda en formato json el customer
		res.Header().Add("Content-Type", "application/json")
		json.NewEncoder(res).Encode(customer)
	}

}

func createCustomer(res http.ResponseWriter, req *http.Request ){
	fmt.Fprintf(res, "Create Customer request received")
}

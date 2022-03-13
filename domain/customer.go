package domain

import "github.com/cspark0610/go-banking-rest-api/errs"

type Customer struct {
	Id string 
	Name string
	City string
	ZipCode string
	DateOfBirth string
	Status string
}

// vamos a definir el SECONDARY PORT(interface): el CustomerRepository Interface
type CustomerRepository interface{
	FindAll() ([]Customer, error)
	// debe retornar un pointer a un customer, y un pointer al struct custom del error
	ById(string) (*Customer, *errs.AppError)
}

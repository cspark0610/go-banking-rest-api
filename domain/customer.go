package domain

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
}

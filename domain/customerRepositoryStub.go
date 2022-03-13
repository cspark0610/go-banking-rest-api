package domain

// definir un struct CustomerRepository para usarlo en el adapter
type CustomerRepositoryStub struct{
	// dummy slice of customers
	customers []Customer 
}

// ADAPTER para el CustomerRepositoryStub, el cual implementa el CustomerRepositoryInterface
func (stub CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return stub.customers, nil
}

// helper function to create a new customer
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	
		customers:= []Customer{
			{Id: "1", Name: "John", City: "New York", ZipCode: "10001", DateOfBirth: "01/01/1990", Status: "1"},
			{Id: "2", Name: "John", City: "New Jersey", ZipCode: "10002", DateOfBirth: "01/01/1990", Status: "1" },
			{Id: "3", Name: "John", City: "New Delhi", ZipCode: "10003" , DateOfBirth: "01/01/1990", Status: "1"},
			{Id: "4", Name: "John", City: "New Hampshire", ZipCode: "10004" , DateOfBirth: "01/01/1990", Status: "1"},
		}
		return CustomerRepositoryStub{customers: customers}
}
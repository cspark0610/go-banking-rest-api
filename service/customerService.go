package service

import "github.com/cspark0610/go-banking-rest-api/domain"

// PRIMARY PORT(interface): el CustomerService Interface
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

// implementacion del CustomerService, la llamamos DefaultCustomerService 
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// ADAPTER para el DefaultCustomerService, el cual implementa el CustomerServiceInterface
func (service DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return service.repo.FindAll()
}

// funcion helper para crear una nueva instancia de CustomerService
// tiene como dependencia el repository "repo" y retorna una nueva instancia de CustomerService
func NewCustomerService(repo domain.CustomerRepository) CustomerService {
	return DefaultCustomerService{repo: repo}
}
package service

import (
	"github.com/cspark0610/go-banking-rest-api/domain"
	"github.com/cspark0610/go-banking-rest-api/errs"
)

// PRIMARY PORT(interface): el CustomerService Interface
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

// implementacion del CustomerService, la llamamos DefaultCustomerService 
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// ADAPTERS para el DefaultCustomerService, el cual implementa el CustomerServiceInterface
func (service DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return service.repo.FindAll()
}

func (service DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return service.repo.ById(id)
}

// funcion helper para crear una nueva instancia de CustomerService
// tiene como dependencia el repository "repo" y retorna una nueva instancia de CustomerService
func NewCustomerService(repo domain.CustomerRepository) CustomerService {
	return DefaultCustomerService{repo: repo}
}
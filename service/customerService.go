package service

import "github.com/Gideon-isa/banking/domian"

type CustomerService interface {
	GetAllCustomer() ([]domian.Customer, error)
}

type DefaultCustomerService struct {
	repo domian.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domian.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domian.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}

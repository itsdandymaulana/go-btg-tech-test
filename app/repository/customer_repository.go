package repository

import "github.com/be-tech-test-btg/app/model"

type CustomerRepository interface {
	GetCustomerByID(cstID uint64) model.Customer
}

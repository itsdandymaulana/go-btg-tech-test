package service

import "github.com/be-tech-test-btg/app/dto"

type CustomerService interface {
	GetCustomerByID(id int) (dto.GetCustomerDTO, error)
}

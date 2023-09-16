package repository

import (
	"github.com/be-tech-test-btg/app/model"
	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &CustomerRepositoryImpl{
		db: db,
	}
}

func (r *CustomerRepositoryImpl) GetCustomerByID(cstID uint64) model.Customer {
	var cust model.Customer
	r.db.Find(&cust, cstID)
	return cust
}

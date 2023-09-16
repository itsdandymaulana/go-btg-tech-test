package repository

import (
	"github.com/be-tech-test-btg/app/model"
	"gorm.io/gorm"
)

type NationalityRepositoryImpl struct {
	db *gorm.DB
}

func NewNationalityRepository(db *gorm.DB) NationalityRepository {
	return &NationalityRepositoryImpl{
		db: db,
	}
}

func (r *NationalityRepositoryImpl) GetNationalityByCstID(natID uint64) model.Nationality {
	var nat model.Nationality
	r.db.Find(&nat, natID)
	return nat
}

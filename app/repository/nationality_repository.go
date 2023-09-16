package repository

import "github.com/be-tech-test-btg/app/model"

type NationalityRepository interface {
	GetNationalityByCstID(natID uint64) model.Nationality
}

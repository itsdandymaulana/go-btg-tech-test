package repository

import (
	"github.com/be-tech-test-btg/app/dto"
)

type FamilyListRepository interface {
	GetFamilyListByCstID(cstID uint64) []dto.FamilyDTO
}

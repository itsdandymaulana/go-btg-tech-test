package repository

import (
	"github.com/be-tech-test-btg/app/dto"
	"gorm.io/gorm"
)

type FamilyListRepositoryImpl struct {
	db *gorm.DB
}

func NewFamilyListRepository(db *gorm.DB) FamilyListRepository {
	return &FamilyListRepositoryImpl{
		db: db,
	}
}

func (r *FamilyListRepositoryImpl) GetFamilyListByCstID(cstID uint64) []dto.FamilyDTO {
	var fl []dto.FamilyDTO
	r.db.Table("family_list").Select([]string{"fl_name", "fl_relation", "fl_dob"}).Where("cst_id = ?", cstID).Scan(&fl)
	return fl
}

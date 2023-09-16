package model

import "time"

type FamilyList struct {
	FlID       uint64 `gorm:"primaryKey:autoIncrement;column:fl_id"`
	CstID      uint64 `gorm:"column:cst_id"`
	FlRelation string `json:"hubungan" gorm:"column:varchar(50)"`
	FlName     string `json:"nama" gorm:"column:varchar(50)"`
	FlDob      string `json:"tanggal_lahir" gorm:"column:date"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Name string `gorm:"tableName:family_list"`
}

package model

import (
	"time"
)

type Customer struct {
	CstID         uint64    `gorm:"primary_key:auto_increment;column:cst_id"`
	CstName       string    `json:"nama" gorm:"type:varchar(50);column:cst_name"`
	CstDob        string    `json:"tanggal_lahir" gorm:"type:date;column:cst_dob"`
	CstPhoneNum   string    `json:"telepon" gorm:"type:varchar(20);column:cst_phoneNum"`
	CstEmail      string    `json:"email" gorm:"type:varchar(50);column:cst_email"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt     time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	NationalityID uint64    `json:"nationality_id" gorm:"column:nationality_id"`
}

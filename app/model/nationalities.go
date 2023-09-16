package model

import (
	"time"
)

type Nationality struct {
	NationalityID   uint64 `gorm:"primaryKey:autoIncrement"`
	NationalityName string `gorm:"column:nationality_name;type:varchar(50)"`
	NationalityCode string `gorm:"column:nationality_code;type:varchar(2)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	Name string `gorm:"tableName:nationality"`
}

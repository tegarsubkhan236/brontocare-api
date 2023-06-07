package model

import (
	"gorm.io/gorm"
	"hospital-api/pkg/api/helper"
	"time"
)

type HspUnit struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UnitCode  string `json:"unit_code"`
	UnitName  string `json:"unit_name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type HspUnits struct {
	HspUnit    []HspUnit                `json:"hsp_unit"`
	Pagination helper.PaginationRequest `json:"pagination"`
}

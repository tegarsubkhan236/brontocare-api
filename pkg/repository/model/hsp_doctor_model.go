package model

import (
	"gorm.io/gorm"
	"hospital-api/pkg/api/helper"
	"time"
)

type HspDoctor struct {
	ID              uint          `gorm:"primaryKey" json:"id"`
	Name            string        `json:"name"`
	HspDisciplineID uint          `json:"hsp_discipline_id"`
	HspDiscipline   HspDiscipline `json:"hsp_discipline"`
	HspUnit         []HspUnit     `gorm:"many2many:hsp_doctors_units" json:"hsp_unit"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type HspDoctors struct {
	HspDoctor  []HspDoctor              `json:"hsp_doctor"`
	Pagination helper.PaginationRequest `json:"pagination"`
}

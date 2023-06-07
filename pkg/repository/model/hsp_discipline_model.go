package model

import (
	"gorm.io/gorm"
	"hospital-api/pkg/api/helper"
	"time"
)

type HspDiscipline struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type HspDisciplines struct {
	HspDiscipline []HspDiscipline          `json:"hsp_discipline"`
	Pagination    helper.PaginationRequest `json:"pagination"`
}

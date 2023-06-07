package model

import "hospital-api/pkg/api/helper"

type CorePermission struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type NewCorePermission struct {
	Name string `json:"name"`
}

type CorePermissions struct {
	Permission []CorePermission         `json:"permissions"`
	Pagination helper.PaginationRequest `json:"pagination"`
}

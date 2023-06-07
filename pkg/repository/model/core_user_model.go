package model

import (
	"gorm.io/gorm"
	"hospital-api/pkg/api/helper"
	"time"
)

type CoreUser struct {
	ID          uint             `gorm:"primaryKey" json:"id"`
	Name        string           `json:"name"`
	Username    string           `json:"username"`
	Sex         string           `json:"sex"`
	Email       string           `json:"email"`
	Password    string           `json:"password"`
	Status      int              `json:"status"`
	Permission  []CorePermission `gorm:"many2many:core_users_permissions" json:"permission"`
	Role        []CoreRole       `gorm:"many2many:core_users_roles" json:"role"`
	HspDoctorID *uint            `json:"hsp_doctor_id"`
	HspDoctor   HspDoctor        `json:"hsp_doctor"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type CoreUsers struct {
	User       []CoreUser               `json:"users"`
	Pagination helper.PaginationRequest `json:"pagination"`
}

type NewCoreUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
}

type UpdateCoreUser struct {
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Sex       string    `json:"sex"`
	Email     string    `json:"email"`
	Status    int       `json:"status"`
}

type UpdateCoreUserPassword struct {
	UpdatedAt   time.Time `json:"updated_at"`
	OldPassword string    `json:"old_password"`
	Password    string    `json:"password"`
}

type AssignPermissionToUser struct {
	Permission []int `json:"permission"`
}

type AssignRoleToUser struct {
	Role []int `json:"role"`
}

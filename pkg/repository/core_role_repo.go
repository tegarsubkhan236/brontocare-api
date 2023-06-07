package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"hospital-api/pkg/api/helper"
	"hospital-api/pkg/repository/model"
	"log"
	"strconv"
	"strings"
)

func (s *storage) GetRoleById(RoleID int) (model.CoreRole, error) {
	var role model.CoreRole
	if err := s.gorm.Preload("Permission").Find(&role, RoleID).Select("*").Error; err != nil {
		log.Printf("this was the error: %v", err)
		return model.CoreRole{}, err
	}
	if role.ID == 0 {
		return model.CoreRole{}, errors.New("record not found error")
	}
	return role, nil
}

func (s *storage) ListRole(page int, perPage int) (model.CoreRoles, error) {
	offset := (page - 1) * perPage
	var roles []model.CoreRole
	s.gorm.Preload("Permission").Find(&roles).Select("*")
	res := model.CoreRoles{
		Roles: roles,
		Pagination: helper.PaginationRequest{
			Page:    page,
			PerPage: perPage,
			Total:   offset,
		},
	}
	return res, nil
}

func (s *storage) CreateRole(request model.NewCoreRole) error {
	var ID int
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("repo err - %v", err)
	}
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
			return
		}
	}(tx)
	err = tx.QueryRow("INSERT INTO core_roles (name) VALUES ($1) RETURNING id;", request.Name).Scan(&ID)
	if err != nil {
		return fmt.Errorf("repo err - %v", err)
	}
	for _, v := range request.Permission {
		_, err = tx.Exec("INSERT INTO core_roles_permissions (core_role_id, core_permission_id) VALUES ($1, $2);", ID, v)
		if err != nil {
			return fmt.Errorf("repo err - %v", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("repo err - %v", err)
	}

	return nil
}

func (s *storage) UpdateRole(RoleID int, request model.CoreRole) error {
	if err := s.gorm.Model(&model.CoreRole{}).Where("id = ?", RoleID).Update("name", request.Name).Error; err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}
	return nil
}

func (s *storage) DeleteRole(RoleID int) error {
	if err := s.gorm.Select("Permission").Delete(&model.CoreRole{ID: uint(RoleID)}).Error; err != nil {
		log.Printf("this was the error: %v", err.Error())
		return err
	}
	return nil
}

func (s *storage) BatchDeleteRole(request model.BatchDeleteRole) error {
	statement := `DELETE FROM core_roles WHERE id = ANY($1::int[])`

	var ids []string
	for _, s := range request.ID {
		ids = append(ids, strconv.Itoa(s))
	}

	param := "{" + strings.Join(ids, ",") + "}"

	err := s.db.QueryRow(statement, param).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) AssignPermissionToRole(RoleID int, permission []model.CorePermission) error {
	var role model.CoreRole
	s.gorm.First(&role, RoleID)
	err := s.gorm.Model(&role).Association("Permission").Replace(permission)
	if err != nil {
		return err
	}
	return nil
}

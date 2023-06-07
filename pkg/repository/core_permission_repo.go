package repository

import (
	"database/sql"
	"fmt"
	"hospital-api/pkg/api/helper"
	"hospital-api/pkg/repository/model"
	"log"
)

func (s *storage) GetPermissionById(PermissionID int) (model.CorePermission, error) {
	var corePermission model.CorePermission

	statement := `SELECT * FROM core_permissions WHERE id = $1`

	err := s.db.QueryRow(statement, PermissionID).Scan(&corePermission.ID, &corePermission.Name)

	if err == sql.ErrNoRows {
		return model.CorePermission{}, fmt.Errorf("unknown value : %d", PermissionID)
	}

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return model.CorePermission{}, err
	}

	return corePermission, nil
}

func (s *storage) ListPermission(page int, perPage int) (model.CorePermissions, error) {
	offset := (page - 1) * perPage

	statement := `SELECT id, "name", count(*) OVER() AS total_count FROM core_permissions ORDER BY id DESC LIMIT $1 OFFSET $2`
	rows, err := s.db.Query(statement, perPage, offset)

	if err != nil {
		log.Printf("this was the error: %v", err)
		return model.CorePermissions{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var corePermissions []model.CorePermission
	var total int
	for rows.Next() {
		var corePermission model.CorePermission
		if err := rows.Scan(&corePermission.ID, &corePermission.Name, &total); err != nil {
			return model.CorePermissions{}, err
		}
		corePermissions = append(corePermissions, corePermission)
	}

	pagination := helper.PaginationRequest{
		Page:    page,
		PerPage: perPage,
		Total:   total,
	}

	res := model.CorePermissions{
		Permission: corePermissions,
		Pagination: pagination,
	}

	return res, nil
}

func (s *storage) CreatePermission(request model.NewCorePermission) error {
	statement := `INSERT INTO core_permissions (name) VALUES ($1);`

	err := s.db.QueryRow(statement, request.Name).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) DeletePermission(PermissionID int) error {
	if err := s.gorm.Delete(&model.CorePermission{ID: uint(PermissionID)}).Error; err != nil {
		log.Printf("this was the error: %v", err.Error())
		return err
	}
	return nil
}

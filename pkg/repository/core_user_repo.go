package repository

import (
	"errors"
	"gorm.io/gorm"
	"hospital-api/pkg/api/helper"
	"hospital-api/pkg/repository/model"
	"log"
)

func (s *storage) AssignPermissionToUser(UserID int, request []model.CorePermission) error {
	var user model.CoreUser
	s.gorm.First(&user, UserID)
	err := s.gorm.Model(&user).Association("Permission").Replace(request)
	if err != nil {
		return err
	}
	return nil
}

func (s *storage) AssignRoleToUser(UserID int, request []model.CoreRole) error {
	var user model.CoreUser
	s.gorm.First(&user, UserID)
	err := s.gorm.Model(&user).Association("Role").Replace(request)
	if err != nil {
		return err
	}
	return nil
}

func (s *storage) CreateUser(request model.NewCoreUser) error {
	statement := `INSERT INTO core_users (name, username, password, sex, email, status) VALUES ($1, $2, $3, $4, $5, $6);`

	err := s.db.QueryRow(statement, request.Name, request.Username, request.Password, request.Sex, request.Email, request.Status).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) GetUserByID(id int) (model.CoreUser, error) {
	var user model.CoreUser
	if err := s.gorm.Preload("Role").
		Preload("Role.Permission").
		Preload("Permission").
		First(&user, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return model.CoreUser{}, err
	}
	return user, nil
}

func (s *storage) GetUserByEmail(email string) (model.CoreUser, error) {
	var user model.CoreUser
	if err := s.gorm.Preload("Role").
		Preload("Role.Permission").
		Preload("Permission").
		Where("email = ?", email).
		First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return model.CoreUser{}, err
	}

	return user, nil
}

func (s *storage) GetUserByUsername(username string) (model.CoreUser, error) {
	var user model.CoreUser
	if err := s.gorm.Preload("Role").
		Preload("Role.Permission").
		Preload("Permission").
		Where("username = ?", username).
		First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return model.CoreUser{}, err
	}

	return user, nil
}

func (s *storage) ListUser(page int, perPage int) (model.CoreUsers, error) {
	offset := (page - 1) * perPage
	var users []model.CoreUser
	s.gorm.Preload("Role").Preload("Role.Permission").Preload("Permission").Find(&users).Select("*")
	res := model.CoreUsers{
		User: users,
		Pagination: helper.PaginationRequest{
			Page:    page,
			PerPage: perPage,
			Total:   offset,
		},
	}
	return res, nil
}

func (s *storage) UpdateUser(UserID int, r model.UpdateCoreUser) error {
	statement := `UPDATE core_users SET name = $1, username = $2, sex = $3, email = $4,  updated_at = $5 WHERE id = $6`

	err := s.db.QueryRow(statement, r.Name, r.Username, r.Sex, r.Email, r.UpdatedAt, UserID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) UpdateUserPassword(UserID int, request model.UpdateCoreUserPassword) error {
	statement := `UPDATE core_users SET password = $1, updated_at = $2 WHERE id = $3`

	err := s.db.QueryRow(statement, request.Password, request.UpdatedAt, UserID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) DeleteUser(UserID int) error {
	if err := s.gorm.Select("Permission").
		Select("Role").
		Delete(&model.CoreRole{ID: uint(UserID)}).Error; err != nil {
		log.Printf("this was the error: %v", err.Error())
		return err
	}
	return nil
}

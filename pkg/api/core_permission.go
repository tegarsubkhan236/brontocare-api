package api

import (
	"errors"
	"hospital-api/pkg/repository/model"
)

// PermissionService contains the methods of the user service
type PermissionService interface {
	List(page int, perPage int) (model.CorePermissions, error)
	Detail(PermissionID int) (model.CorePermission, error)
	New(request model.NewCorePermission) error
	Delete(PermissionID int) error
}

// PermissionRepository is what lets our service do db operations without knowing anything about the implementation
type PermissionRepository interface {
	GetPermissionById(PermissionID int) (model.CorePermission, error)
	ListPermission(page int, perPage int) (model.CorePermissions, error)
	CreatePermission(request model.NewCorePermission) error
	DeletePermission(PermissionID int) error
}

type permissionService struct {
	storage PermissionRepository
}

func (p permissionService) List(page int, perPage int) (model.CorePermissions, error) {
	data, err := p.storage.ListPermission(page, perPage)
	if err != nil {
		return model.CorePermissions{}, err
	}
	return data, nil
}

func (p permissionService) Detail(PermissionID int) (model.CorePermission, error) {
	item, err := p.storage.GetPermissionById(PermissionID)
	if err != nil {
		return model.CorePermission{}, errors.New("permission id not found")
	}
	return item, nil
}

func (p permissionService) New(request model.NewCorePermission) error {
	if request.Name == "" {
		return errors.New("role service - name required")
	}

	err := p.storage.CreatePermission(request)

	if err != nil {
		return err
	}

	return nil
}

func (p permissionService) Delete(PermissionID int) error {
	err := p.storage.DeletePermission(PermissionID)
	if err != nil {
		return err
	}
	return nil
}

func NewPermissionService(permissionRepo PermissionRepository) PermissionService {
	return &permissionService{
		storage: permissionRepo,
	}
}

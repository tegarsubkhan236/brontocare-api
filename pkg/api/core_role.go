package api

import (
	"errors"
	"hospital-api/pkg/repository/model"
	"strings"
)

// RoleService contains the methods of the user service
type RoleService interface {
	Detail(RoleID int) (model.CoreRole, error)
	List(page int, perPage int) (model.CoreRoles, error)
	New(request model.NewCoreRole) error
	Update(RoleID int, role model.CoreRole) error
	Delete(RoleID int) error
	BatchDelete(request model.BatchDeleteRole) error
	AssignPermission(RoleID int, request model.AssignPermissionToRole) error
}

// RoleRepository is what lets our service do db operations without knowing anything about the implementation
type RoleRepository interface {
	GetRoleById(RoleID int) (model.CoreRole, error)
	GetPermissionById(PermissionID int) (model.CorePermission, error)
	ListRole(page int, perPage int) (model.CoreRoles, error)
	CreateRole(request model.NewCoreRole) error
	UpdateRole(RoleID int, role model.CoreRole) error
	DeleteRole(RoleID int) error
	BatchDeleteRole(request model.BatchDeleteRole) error
	AssignPermissionToRole(UserID int, request []model.CorePermission) error
}

type roleService struct {
	storage RoleRepository
}

func (s *roleService) Detail(RoleID int) (model.CoreRole, error) {
	item, err := s.storage.GetRoleById(RoleID)
	if err != nil {
		return model.CoreRole{}, errors.New("role id not found")
	}
	return item, nil
}

func (s *roleService) List(page int, perPage int) (model.CoreRoles, error) {
	roles, err := s.storage.ListRole(page, perPage)
	if err != nil {
		return model.CoreRoles{}, err
	}
	return roles, nil
}

func (s *roleService) New(request model.NewCoreRole) error {
	if request.Name == "" {
		return errors.New("role service - name required")
	}
	if request.Permission == nil {
		return errors.New("role service - Permission required")
	}
	request.Name = strings.ToUpper(request.Name)

	err := s.storage.CreateRole(request)

	if err != nil {
		return err
	}

	return nil
}

func (s *roleService) Update(RoleID int, r model.CoreRole) error {
	if r.Name == "" {
		return errors.New("role service - name required")
	}
	r.Name = strings.ToUpper(r.Name)

	_, err := s.storage.GetRoleById(RoleID)
	if err != nil {
		return err
	}

	err = s.storage.UpdateRole(RoleID, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *roleService) AssignPermission(RoleID int, request model.AssignPermissionToRole) error {
	var permissions []model.CorePermission
	for _, v := range request.Permission {
		res, err := s.storage.GetPermissionById(v)
		if err != nil {
			return err
		}
		permissions = append(permissions, res)
	}
	err := s.storage.AssignPermissionToRole(RoleID, permissions)
	if err != nil {
		return err
	}
	return nil
}

func (s *roleService) Delete(RoleID int) error {
	err := s.storage.DeleteRole(RoleID)
	if err != nil {
		return err
	}
	return nil
}

func (s *roleService) BatchDelete(request model.BatchDeleteRole) error {
	err := s.storage.BatchDeleteRole(request)
	if err != nil {
		return err
	}
	return nil
}

func NewRoleService(roleRepo RoleRepository) RoleService {
	return &roleService{
		storage: roleRepo,
	}
}

package api

import (
	"errors"
	"hospital-api/pkg/repository/model"
	"strings"
)

// HspUnitService contains the methods of the user service
type HspUnitService interface {
	Detail(HspUnitID int) (model.HspUnit, error)
	List(page int, perPage int) (model.HspUnits, error)
	New(r model.HspUnit) error
	Update(HspUnitID int, r model.HspUnit) error
	Delete(HspUnitID int) error
}

// HspUnitRepository is what lets our service do db operations without knowing anything about the implementation
type HspUnitRepository interface {
	GetHspUnitByID(HspUnitID int) (model.HspUnit, error)
	ListHspUnit(page int, perPage int) (model.HspUnits, error)
	CreateHspUnit(r model.HspUnit) error
	UpdateHspUnit(HspUnitID int, r model.HspUnit) error
	DeleteHspUnit(HspUnitID int) error
}

type hspUnitService struct {
	storage HspUnitRepository
}

func (s *hspUnitService) Detail(HspUnitID int) (model.HspUnit, error) {
	item, err := s.storage.GetHspUnitByID(HspUnitID)
	if err != nil {
		return model.HspUnit{}, errors.New("r id not found")
	}
	return item, nil
}

func (s *hspUnitService) List(page int, perPage int) (model.HspUnits, error) {
	roles, err := s.storage.ListHspUnit(page, perPage)
	if err != nil {
		return model.HspUnits{}, err
	}
	return roles, nil
}

func (s *hspUnitService) New(r model.HspUnit) error {
	if r.UnitCode == "" {
		return errors.New("r service - UnitCode required")
	}
	if r.UnitName == "" {
		return errors.New("r service - UnitName required")
	}
	r.UnitCode = strings.ToUpper(r.UnitCode)

	err := s.storage.CreateHspUnit(r)

	if err != nil {
		return err
	}

	return nil
}

func (s *hspUnitService) Update(HspUnitID int, r model.HspUnit) error {
	if r.UnitCode == "" {
		return errors.New("r service - UnitCode required")
	}
	if r.UnitName == "" {
		return errors.New("r service - UnitName required")
	}
	r.UnitCode = strings.ToUpper(r.UnitCode)

	_, err := s.storage.GetHspUnitByID(HspUnitID)
	if err != nil {
		return err
	}

	err = s.storage.UpdateHspUnit(HspUnitID, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *hspUnitService) Delete(HspUnitID int) error {
	err := s.storage.DeleteHspUnit(HspUnitID)
	if err != nil {
		return err
	}
	return nil
}

func NewHspUnitService(s HspUnitRepository) HspUnitService {
	return &hspUnitService{
		storage: s,
	}
}

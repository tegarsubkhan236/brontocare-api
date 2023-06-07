package api

import (
	"errors"
	"hospital-api/pkg/repository/model"
	"strings"
)

// HspDisciplineService contains the methods of the user service
type HspDisciplineService interface {
	Detail(HspDisciplineID int) (model.HspDiscipline, error)
	List(page int, perPage int) (model.HspDisciplines, error)
	New(r model.HspDiscipline) error
	Update(HspDisciplineID int, r model.HspDiscipline) error
	Delete(HspDisciplineID int) error
}

// HspDisciplineRepository is what lets our service do db operations without knowing anything about the implementation
type HspDisciplineRepository interface {
	GetHspDisciplineByID(HspDisciplineID int) (model.HspDiscipline, error)
	ListHspDiscipline(page int, perPage int) (model.HspDisciplines, error)
	CreateHspDiscipline(r model.HspDiscipline) error
	UpdateHspDiscipline(HspDisciplineID int, r model.HspDiscipline) error
	DeleteHspDiscipline(HspDisciplineID int) error
}

type hspDisciplineService struct {
	storage HspDisciplineRepository
}

func (s *hspDisciplineService) Detail(HspDisciplineID int) (model.HspDiscipline, error) {
	item, err := s.storage.GetHspDisciplineByID(HspDisciplineID)
	if err != nil {
		return model.HspDiscipline{}, errors.New("r id not found")
	}
	return item, nil
}

func (s *hspDisciplineService) List(page int, perPage int) (model.HspDisciplines, error) {
	roles, err := s.storage.ListHspDiscipline(page, perPage)
	if err != nil {
		return model.HspDisciplines{}, err
	}
	return roles, nil
}

func (s *hspDisciplineService) New(r model.HspDiscipline) error {
	if r.Name == "" {
		return errors.New("r service - name required")
	}
	r.Name = strings.ToUpper(r.Name)

	err := s.storage.CreateHspDiscipline(r)

	if err != nil {
		return err
	}

	return nil
}

func (s *hspDisciplineService) Update(HspDisciplineID int, r model.HspDiscipline) error {
	if r.Name == "" {
		return errors.New("r service - name required")
	}
	r.Name = strings.ToUpper(r.Name)

	_, err := s.storage.GetHspDisciplineByID(HspDisciplineID)
	if err != nil {
		return err
	}

	err = s.storage.UpdateHspDiscipline(HspDisciplineID, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *hspDisciplineService) Delete(HspDisciplineID int) error {
	err := s.storage.DeleteHspDiscipline(HspDisciplineID)
	if err != nil {
		return err
	}
	return nil
}

func NewHspDisciplineService(s HspDisciplineRepository) HspDisciplineService {
	return &hspDisciplineService{
		storage: s,
	}
}

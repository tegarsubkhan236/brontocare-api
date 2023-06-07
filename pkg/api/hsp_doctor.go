package api

import (
	"errors"
	"hospital-api/pkg/repository/model"
	"strings"
)

// HspDoctorService contains the methods of the user service
type HspDoctorService interface {
	Detail(HspDoctorID int) (model.HspDoctor, error)
	List(page int, perPage int) (model.HspDoctors, error)
	New(r model.HspDoctor) error
	Update(HspDoctorID int, r model.HspDoctor) error
	Delete(HspDoctorID int) error
}

// HspDoctorRepository is what lets our service do db operations without knowing anything about the implementation
type HspDoctorRepository interface {
	GetHspDoctorByID(HspDoctorID int) (model.HspDoctor, error)
	ListHspDoctor(page int, perPage int) (model.HspDoctors, error)
	CreateHspDoctor(r model.HspDoctor) error
	UpdateHspDoctor(HspDoctorID int, r model.HspDoctor) error
	DeleteHspDoctor(HspDoctorID int) error
}

type hspDoctorService struct {
	storage HspDoctorRepository
}

func (s *hspDoctorService) Detail(HspDoctorID int) (model.HspDoctor, error) {
	item, err := s.storage.GetHspDoctorByID(HspDoctorID)
	if err != nil {
		return model.HspDoctor{}, errors.New("r id not found")
	}
	return item, nil
}

func (s *hspDoctorService) List(page int, perPage int) (model.HspDoctors, error) {
	roles, err := s.storage.ListHspDoctor(page, perPage)
	if err != nil {
		return model.HspDoctors{}, err
	}
	return roles, nil
}

func (s *hspDoctorService) New(r model.HspDoctor) error {
	if r.Name == "" {
		return errors.New("r service - name required")
	}
	r.Name = strings.ToUpper(r.Name)

	err := s.storage.CreateHspDoctor(r)

	if err != nil {
		return err
	}

	return nil
}

func (s *hspDoctorService) Update(HspDoctorID int, r model.HspDoctor) error {
	if r.Name == "" {
		return errors.New("r service - name required")
	}
	r.Name = strings.ToUpper(r.Name)

	_, err := s.storage.GetHspDoctorByID(HspDoctorID)
	if err != nil {
		return err
	}

	err = s.storage.UpdateHspDoctor(HspDoctorID, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *hspDoctorService) Delete(HspDoctorID int) error {
	err := s.storage.DeleteHspDoctor(HspDoctorID)
	if err != nil {
		return err
	}
	return nil
}

func NewHspDoctorService(s HspDoctorRepository) HspDoctorService {
	return &hspDoctorService{
		storage: s,
	}
}

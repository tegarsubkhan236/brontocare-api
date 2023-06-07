package repository

import "hospital-api/pkg/repository/model"

func (s *storage) GetHspDoctorByID(HspDoctorID int) (model.HspDoctor, error) {
	//TODO implement me
	panic("implement me")
}

func (s *storage) ListHspDoctor(page int, perPage int) (model.HspDoctors, error) {
	//TODO implement me
	panic("implement me")
}

func (s *storage) CreateHspDoctor(r model.HspDoctor) error {
	//TODO implement me
	panic("implement me")
}

func (s *storage) UpdateHspDoctor(HspDoctorID int, r model.HspDoctor) error {
	//TODO implement me
	panic("implement me")
}

func (s *storage) DeleteHspDoctor(HspDoctorID int) error {
	//TODO implement me
	panic("implement me")
}

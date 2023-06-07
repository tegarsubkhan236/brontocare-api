package repository

import "hospital-api/pkg/repository/model"

func (s *storage) GetHspDisciplineByID(HspDisciplineID int) (model.HspDiscipline, error) {
	//TODO implement me
	panic("implement me")
}

func (s *storage) ListHspDiscipline(page int, perPage int) (model.HspDisciplines, error) {
	//TODO implement me
	panic("implement me")
}

func (s *storage) CreateHspDiscipline(r model.HspDiscipline) error {
	//TODO implement me
	panic("implement me")
}

func (s *storage) UpdateHspDiscipline(HspDisciplineID int, r model.HspDiscipline) error {
	//TODO implement me
	panic("implement me")
}

func (s *storage) DeleteHspDiscipline(HspDisciplineID int) error {
	//TODO implement me
	panic("implement me")
}

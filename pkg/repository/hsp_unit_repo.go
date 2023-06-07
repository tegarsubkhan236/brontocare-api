package repository

import "hospital-api/pkg/repository/model"

func (s *storage) GetHspUnitByID(HspUnitID int) (model.HspUnit, error) {
	//TODO implement me
	panic("implement me")
}

func (s *storage) ListHspUnit(page int, perPage int) (model.HspUnits, error) {
	//TODO implement me
	panic("implement me")
}

func (s *storage) CreateHspUnit(r model.HspUnit) error {
	//TODO implement me
	panic("implement me")
}

func (s *storage) UpdateHspUnit(HspUnitID int, r model.HspUnit) error {
	//TODO implement me
	panic("implement me")
}

func (s *storage) DeleteHspUnit(HspUnitID int) error {
	//TODO implement me
	panic("implement me")
}

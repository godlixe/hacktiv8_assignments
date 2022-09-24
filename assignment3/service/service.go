package service

import (
	"assignment3/entity"
	"assignment3/helpers"
	"assignment3/repository"
	"time"
)

type Service interface {
	GetData() (entity.Response, error)
	UpdateData() error
}

type service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetData() (entity.Response, error) {
	return s.repository.GetData()
}

func (s *service) UpdateData() error {
	var response entity.Response

	response.Data.Water = helpers.GenerateRandomNumber(1, 100)
	response.Data.Wind = helpers.GenerateRandomNumber(1, 100)

	response.Status = entity.DataStatus{
		WaterStatus: helpers.CheckWaterStatus(response.Data.Water),
		WindStatus:  helpers.CheckWindStatus(response.Data.Wind),
	}

	response.UpdatedAt = time.Now()

	return s.repository.UpdateData(response)
}

package repository

import (
	"assignment3/entity"
	"encoding/json"
	"io/ioutil"
)

type Repository interface {
	GetData() (entity.Response, error)
	UpdateData(entity.Response) error
}

type repository struct {
	filePath string
}

func NewDataRepository(filePath string) Repository {
	return &repository{
		filePath: filePath,
	}
}

func (r *repository) GetData() (entity.Response, error) {
	byte, err := ioutil.ReadFile(r.filePath)
	if err != nil {
		return entity.Response{}, err
	}

	var response entity.Response
	err = json.Unmarshal(byte, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (r *repository) UpdateData(response entity.Response) error {
	byte, err := json.Marshal(response)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.filePath, byte, 0777)
	if err != nil {
		return err
	}

	return nil
}

package repository

import (
	"encoding/json"
	"net/http"
	"order_service/dto"
	"order_service/entity"

	"github.com/mashingan/smapping"
)

type PersonRepository interface {
	GetPerson() (entity.Person, error)
}

type personRepository struct {
	host string
}

func NewPersonRepository(hostURL string) PersonRepository {
	return &personRepository{
		host: hostURL,
	}
}

func (r *personRepository) GetPerson() (entity.Person, error) {
	var person entity.Person
	var data dto.GidhanAPIResponseDTO
	var client = &http.Client{}

	request, err := http.NewRequest(http.MethodGet, r.host+"/data.php?qty=1&apikey=7f8fc96e-de1f-4aab-9c62-3dd1de365e66", nil)
	if err != nil {
		return person, err
	}

	response, err := client.Do(request)
	if err != nil {
		return person, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return person, err
	}

	err = smapping.FillStruct(&person, smapping.MapFields(&data.Result[0]))
	if err != nil {
		return person, err
	}

	return person, nil
}

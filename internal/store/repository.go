package store

import "github.com/RustamRR/cli-wether-app/internal/model"

type CityRepository interface {
	Create(*model.City) error
	FindByName(string) (*model.City, error)
}

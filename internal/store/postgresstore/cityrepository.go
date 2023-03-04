package postgresstore

import (
	"github.com/RustamRR/cli-wether-app/internal/model"
)

type CityRepository struct {
	store *Store
}

func (r *CityRepository) Create(c *model.City) error {
	if err := c.Validate(); err != nil {
		return err
	}

	if err := r.store.db.Create(c).Error; err != nil {
		return err
	}

	return nil
}

func (r CityRepository) FindByName(name string) (*model.City, error) {
	var city model.City
	if err := r.store.db.First(&city, "name = ?", name).Error; err != nil {
		return nil, err
	}

	return &city, nil
}

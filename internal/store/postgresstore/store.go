package postgresstore

import (
	"github.com/RustamRR/cli-wether-app/internal/model"
	"github.com/RustamRR/cli-wether-app/internal/store"
	"gorm.io/gorm"
)

type Store struct {
	db             *gorm.DB
	cityRepository *CityRepository
}

func New(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) City() store.CityRepository {
	if s.cityRepository != nil {
		return s.cityRepository
	}

	s.cityRepository = &CityRepository{
		store: s,
	}

	return s.cityRepository
}

func (s *Store) Migrate() error {
	err := s.db.Migrator().AutoMigrate(&model.City{})
	if err != nil {
		return err
	}

	return nil
}

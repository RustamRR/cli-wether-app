package app

import (
	"errors"
	"github.com/RustamRR/cli-wether-app/internal/model"
	"github.com/RustamRR/cli-wether-app/internal/store"
	"github.com/RustamRR/cli-wether-app/internal/store/postgresstore"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	config *viper.Viper
	logger *zap.Logger
	store  store.Store
}

func (s *Server) Migrate() {
	if err := s.store.Migrate(); err != nil {
		s.logger.Fatal(err.Error())
		panic(err)
	}
}

func New(config *viper.Viper) *Server {
	db, err := gorm.Open(postgres.Open(config.GetString("dsn")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	pgStore := postgresstore.New(db)

	return &Server{
		config: config,
		logger: logger,
		store:  pgStore,
	}
}

func (s *Server) GetWeatherForCity(name string) (CurrentWeatherResult, error) {
	city, err := s.getCity(name)
	if err != nil {
		s.logger.Warn(err.Error())
		return CurrentWeatherResult{}, err
	}

	weatherData, err := GetWeatherByCoordinates(city.Latitude, city.Longitude)
	if err != nil {
		s.logger.Warn(err.Error())
		return CurrentWeatherResult{}, err
	}

	weatherData.Name = city.Title

	return weatherData, nil
}

func (s *Server) getCity(name string) (*model.City, error) {
	city, err := s.store.City().FindByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			city, err := s.addCity(name)
			if err == nil {
				return city, nil
			}
		}

		return nil, err
	}

	return city, nil
}

func (s *Server) addCity(name string) (*model.City, error) {
	cityData, err := GetGeographicalCoordinates(name)
	if err != nil {
		return nil, err
	}

	city := &model.City{
		Name:      name,
		Title:     cityData.Name,
		Longitude: cityData.Longitude,
		Latitude:  cityData.Latitude,
	}

	if err := s.store.City().Create(city); err != nil {
		return nil, err
	}

	return city, nil
}

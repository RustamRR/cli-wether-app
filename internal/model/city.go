package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Title     string  `json:"title"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (c *City) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.Title, validation.Required),
		validation.Field(&c.Latitude, validation.Required),
		validation.Field(&c.Longitude, validation.Required),
	)
}

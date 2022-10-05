package model

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	FirstName string `json:"first_name" example:"David"`
	LastName  string `json:"last_name" example:"Louis"`
}

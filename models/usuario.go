package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Name           string `json:"name"`
	Lastname       string `json:"lastname"`
	Age            int8   `json:"age"`
	Address        string `json:"address"`
	Phone          int    `json:"phone"`
	Identification string `json:"identification" gorm:"not null; unique"`
}

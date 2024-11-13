package models

import "gorm.io/gorm"

type libro struct {
	gorm.Model
	Titulo              string `json:"titulo"`
	ISBN                string `json:"isbn"`
	Autor               string `json:"autor"`
	Cantidad_Ejemplares int    `json:"cantidad_ejemplares"`
	Codigo_Interno      string `json:"codigo_interno"`
}

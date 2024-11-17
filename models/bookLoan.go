package models

import (
	"time"

	"gorm.io/gorm"
)

type Prestamo struct {
	gorm.Model
	UsuarioID        uint
	LibroID          uint
	Fecha_Prestamo   time.Time `json:"fecha_prestamo"`
	Fecha_Devolucion time.Time `json:"fecha_devolucion"`
	Estado_Prestamo  bool      `json:"estado_prestamo"`
}

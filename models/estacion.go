package models

import (
	"github.com/LauraBarriosg/PRUEBA/libs"
	"github.com/jinzhu/gorm"
	"time"
)

// Estructura de Estaciones
type Estaciones struct {
	gorm.Model
	Nombre		  	string		`gorm:"size:255;not null;unique" json:"nombre"`
	Ubicacion	  	string		`gorm:"size:100;not null;unique" json:"ubicacion"`
	Estado        	bool		`gorm:"size:100;not null;default:True" json:"estado"`
	Fecha         	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"fecha"`
	DireccionMac	string		`gorm:"size:200;not null;unique" json:"mac"`
}

// Crear tabla de Estaciones
func TableEstaciones(){
	libs.DB.AutoMigrate(&Estaciones{})
}
 
// Crear Estaciones
func CreateEstacion(){
	estacion := Estaciones{Nombre:"Estacion 3", Ubicacion: "Unimet 1.2", DireccionMac: "prueba3313"}
	libs.DB.Create(&estacion)
}


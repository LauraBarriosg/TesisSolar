package models

import (
	"github.com/LauraBarriosg/PRUEBA/libs"
	"github.com/jinzhu/gorm"
	//"fmt"
)

// Estructura de Dispositivos
type Dispositivos struct {
	gorm.Model
	EstacionMac		string	`gorm:"not null" json:"estacion_mac"`
	Modelo			string	`gorm:"size:100;not null" json:"modelo"`
	Variable		string	`gorm:"size:100;not null" json:"variable"`
	LimiteInferior	string	`gorm:"size:100;not null;" json:"limite_inferior"`
	LimiteSuperior	string	`gorm:"size:100;not null;" json:"limite_superior"`
	Estado			bool	`gorm:"size:100;not null;default:True" json:"estado"`
}

// Crear tabla de Dispositivos
func TableDispositivos(){
	libs.DB.AutoMigrate(&Dispositivos{})
	libs.DB.Model(&Dispositivos{}).AddForeignKey("estacion_mac", "estaciones(direccion_mac)", "CASCADE", "NO ACTION")
}

// Crear Dispositivos
func CreateDispositivo(){
	libs.DB.Create(&Dispositivos{
		Modelo:"2", 
		Variable: "Temperatura", 
		LimiteInferior: "30",
		LimiteSuperior: "100",
		EstacionMac: "prueba3313",
	})
}







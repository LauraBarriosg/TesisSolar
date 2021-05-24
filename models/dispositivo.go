package models

import (
	"github.com/LauraBarriosg/PRUEBA/libs"
	"github.com/jinzhu/gorm"
	"fmt"
	//"math/rand"
	//"time"
	//"strconv"
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
	libs.DB.Model(&Dispositivos{}).AddForeignKey("estacion_mac", "estaciones(estacion_mac)", "CASCADE", "NO ACTION")
}
var i int = 0
// Crear Dispositivos
func CreateDispositivo(estacion_mac string){
	X := fmt.Sprintf("Modelo %d", i)
	libs.DB.Create(&Dispositivos{
		Modelo:X, 
		Variable: "Temperatura", 
		LimiteInferior: "30",
		LimiteSuperior: "100",
		EstacionMac: estacion_mac,
	})
	i++
}







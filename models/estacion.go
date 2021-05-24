package models

import (
	"github.com/LauraBarriosg/PRUEBA/libs"
	"github.com/jinzhu/gorm"
	"time"
	"fmt"
)

// Estructura de Estaciones
type Estaciones struct {
	gorm.Model
	Nombre		  	string		`gorm:"size:255;not null;unique" json:"nombre"`
	Ubicacion	  	string		`gorm:"size:100;not null" json:"ubicacion"`
	Estado        	bool		`gorm:"size:100;not null;default:True" json:"estado"`
	Fecha         	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"fecha"`
	EstacionMac		string		`gorm:"size:200;not null;unique" json:"estacion_mac"`
}

// Crear tabla de Estaciones
func TableEstaciones(){
	libs.DB.AutoMigrate(&Estaciones{})
}
 
// Crear Estaciones
func CreateEstacion(){
	for i := 1 ; i < 1001; i++{
		fmt.Printf("%d", i)
		Y := fmt.Sprintf("Codigo %d", i)
		name:= fmt.Sprintf("Estacion %d",i)
		e := &Estaciones{
			Nombre:name, 
			Ubicacion: "Unimet", 
			EstacionMac: Y,
		}
		libs.DB.Create(&e)
		CreateDispositivo(e.EstacionMac)
		fmt.Printf("Pasa por aqui %v \n", i)
	}
}


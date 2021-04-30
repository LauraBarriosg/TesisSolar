package models

import (
	"time"
	"fmt"
	"github.com/LauraBarriosg/PRUEBA/libs"
	"github.com/jinzhu/gorm"
)

// Estructura de Alertas
type Alertas struct {
	gorm.Model
	DispositivoId	int           `gorm:"not null" json:"dispositivo_id"`
	Fecha			time.Time     `gorm:"not null" json:"created_at" json:"fecha"`
	Contador		int           `gorm:"size:150; not null; default:'0'" json:"contador"`
	Reporte			string        `gorm:"size:300; not null" json:"reporte"`
}

// Crear tabla de Alertas
func TableAlertas(){	
	libs.DB.AutoMigrate(&Alertas{})
	libs.DB.Model(&Alertas{}).AddForeignKey("dispositivo_id", "dispositivos(id)", "CASCADE", "NO ACTION")
	//libs.DB.Model(&Dispositivos{}).AddForeignKey("estacion_mac", "estaciones(direccion_mac)", "CASCADE", "NO ACTION")
}

// Se usará para definir las alertas
func Alerta(mac string){
	var dispo Dispositivos
	libs.DB.First(&dispo, "estacion_mac = ? AND veriable = ? ", mac, "Temperatura")
	X := dispo.ID
	fmt.Printf("%d\n", X)
}
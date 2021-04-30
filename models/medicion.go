package models

import (
	"time"
	"github.com/LauraBarriosg/PRUEBA/libs"
	"github.com/jinzhu/gorm"
	"encoding/json"
    "log"
	"fmt"
	//"strconv"
)

// Esructura de Mediciones
type Mediciones struct {
	gorm.Model
	DispositivoId	uint		`gorm:"not null" json:"dispositivo_id"` 
	Fecha			time.Time	`gorm:"not null" json:"fecha"`
	ValorMedido		float64		`gorm:"not null" json:"valor_medido"`
}

// Estrucrura del Json
type DatosJson struct {
	EstacionMac		string		
	Fecha			time.Time
	Datos			[]DataPoint
}

// Estructura para los datos de Json
type DataPoint struct {
    Variable	string
    Valor		float64
}

// Se crea tabla de Mediciones
func TableMediciones(){
	libs.DB.AutoMigrate(&Mediciones{})
	libs.DB.Model(&Mediciones{}).AddForeignKey("dispositivo_id", "dispositivos(id)", "CASCADE", "NO ACTION")
}

// Se envian las mediciones a la Base de Datos
func CreateMedicion(Mac string, Fecha time.Time, Valor_medido float64, Variable string){
	var dispo Dispositivos
	libs.DB.First(&dispo, "estacion_mac = ? AND variable = ? ", Mac, Variable)
	X := dispo.ID
	libs.DB.Create(&Mediciones{
		DispositivoId: X,
		Fecha: Fecha,
		ValorMedido: Valor_medido, 
	})
}
//---------------------------------------------------------------------------
// Solo para contar los mensajes que llegan
var X int = 0
var Y int = 0
var Z int = 0
//---------------------------------------------------------------------------

// Extraer medicion del JSON
func ExtraerDatos(msg string){
	var err error
	var medicion DatosJson
	err = json.Unmarshal([]byte(msg), &medicion)
	if err != nil {
		log.Printf("error  %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
	}
	for i,_ := range medicion.Datos {
		variable := medicion.Datos[i].Variable
		valor := medicion.Datos[i].Valor
		//fmt.Printf("%s\n", variable)
		//fmt.Printf("%v\n", valor)
		CreateMedicion(medicion.EstacionMac, medicion.Fecha, valor, variable)
	}
	// Para contar los mensajes que llegan
	if medicion.EstacionMac == "prueba3312"{
		X++
		fmt.Printf("Publicador 1 ha enviado --> %d mensajes\n", X)
	}
	if medicion.EstacionMac == "prueba33" {
		Y++
		fmt.Printf("Publicador 2 ha enviado --> %d mensajes\n", Y)
	}
	if medicion.EstacionMac == "prueba3313" {
		Z++
		fmt.Printf("Publicador 3 ha enviado --> %d mensajes\n", Z)
	}
}	
			
	



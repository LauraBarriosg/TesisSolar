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
	EstacionMac			string		`gorm:"not null" json:"estacion_mac"` 
	Fecha				time.Time	`gorm:"not null" json:"fecha"`
	GHI					float64		`gorm:"default: null" json:"ghi"`
	DHI					float64		`gorm:"default: null" json:"dhi"`
	DNI					float64		`gorm:"default: null" json:"dni"`
	Temperatura			float64		`gorm:"default: null"" json:"temperatura"`
	Humedad				float64		`gorm:"default: null" json:"humedad"`
	Lluvia				float64		`gorm:"default: null" json:"lluvia"`
	VelocidadViento		float64		`gorm:"default: null" json:"velocidad_viento"`
	DireccionViento		float64		`gorm:"default: null" json:"direccion_viento"`
	PresionAtm			float64		`gorm:"default: null" json:"presion_atm"`
}

// Se crea tabla de Mediciones
func TableMediciones(){
	libs.DB.AutoMigrate(&Mediciones{})
	libs.DB.Model(&Dispositivos{}).AddForeignKey("estacion_mac", "estaciones(estacion_mac)", "CASCADE", "NO ACTION")
}

// Se envian las mediciones a la Base de Datos
func CrearMedicion(Mac string, Fecha time.Time, GHI float64, DHI float64, DNI float64, Temperatura float64, Humedad float64, Lluvia float64, DireccionViento float64, VelocidadViento float64, PresionAtm float64){
	
	libs.DB.Create(&Mediciones{
		EstacionMac: Mac,
		Fecha: Fecha,
		GHI: GHI,
		DHI: DHI,
		DNI: DNI,
		Temperatura: Temperatura,
		Humedad: Humedad,
		Lluvia: Lluvia,
		VelocidadViento: VelocidadViento,
		DireccionViento: DireccionViento,
		PresionAtm: PresionAtm, 
	})
	
}

// Estrucrura del Json
type DatosJson struct {
	EstacionMac			string		
	Fecha				time.Time	
	GHI					float64
	DHI					float64
	DNI					float64
	Temperatura			float64
	Humedad				float64
	Lluvia				float64
	VelocidadViento		float64
	DireccionViento		float64
	PresionAtm			float64
}

//---------------------------------------------------------------------------
// Solo para contar los mensajes que llegan
var X int = 0
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
	/*defer libs.DB.Close()
    dbConfig := libs.Configure("./", "postgres")
	libs.DB = dbConfig.InitPostgresDB()
	
	go*/ CrearMedicion(
		medicion.EstacionMac,
		medicion.Fecha,
		medicion.GHI, 
		medicion.DNI, 
		medicion.DHI,
		medicion.Temperatura,
		medicion.Humedad,
		medicion.Lluvia,
		medicion.VelocidadViento,
		medicion.DireccionViento,
		medicion.PresionAtm,
	)
	
	// Para contar los mensajes que llegan
	X++
	fmt.Printf("se han recibido --> %d mensajes\n", X)
}	
			
	



package main

import (
	"github.com/LauraBarriosg/PRUEBA/models"
	"github.com/LauraBarriosg/PRUEBA/libs"
	"github.com/LauraBarriosg/PRUEBA/mqtt"
)

func main() {
	//se abre la conección con postgres
	dbConfig := libs.Configure("./", "postgres")
	libs.DB = dbConfig.InitPostgresDB()

	//se crean las tablas
	models.TableUser()
	models.TableEstaciones()
	models.TableDispositivos()
	models.TableMediciones()
	models.TableAlertas()

	//Se crean algunos datos 
	//models.CreateEstacion()
	//models.CreateDispositivo()

	//Se conecta con mqtt
    mqtt.Mqtt()
}

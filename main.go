package main

import (
	"log"
	"github.com/LauraBarriosg/PRUEBA/models"
	"github.com/LauraBarriosg/PRUEBA/libs"
	"github.com/LauraBarriosg/PRUEBA/mqtt"
	//"os"
	//"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//"github.com/jinzhu/gorm"
)

func main() {
	//se abre la conección con postgres
	dbConfig, err := libs.LoadConfig("../../config","app")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	libs.DB = dbConfig.InitPostgresDB()

	/*c, err := libs.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}


	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s  sslmode=disable ", 
	c.User, c.Password, c.Host, c.Port, c.Database)
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	
	}*/













	//se crean las tablas
	models.TableUser()
	models.TableEstaciones()
	models.TableDispositivos()
	models.TableMediciones()
	models.TableAlertas()

	//Se crean algunos datos 
	//models.CreateEstacion()

	//libs.DB.Close()
	//Se conecta con mqtt

    mqtt.Mqtt()
}

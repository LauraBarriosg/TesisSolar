package models

import (
	"github.com/LauraBarriosg/PRUEBA/libs"
	"github.com/jinzhu/gorm"
)

// Estructura de Usuarios
type Users struct {
	gorm.Model
	Nombre		string	`gorm:"size:255;not null;unique" json:"nombre"`
	Email		string	`gorm:"size:100;not null;unique" json:"email"`
	Password	string	`gorm:"size:100;not null;" json:"password"`
	Admin		bool	`gorm:"size:100;not null;default:True" json:"admin"`
}

// Crear tabla de Usuarios
func TableUser(){	
	libs.DB.AutoMigrate(&Users{})
}

// Crear usuarios
func CreateUser(){
	user := Users{Nombre:"Laura", Email: "LauraBarriosg", Password:"lllllll"}
	libs.DB.Create(&user)
}

// Actualizar Usuarios
func UpdateUser(){

}
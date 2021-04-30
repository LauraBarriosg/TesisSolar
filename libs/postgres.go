package libs

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

func (c *DbConfig) InitPostgresDB() *gorm.DB {
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s  sslmode=disable ", c.User, c.Password, c.Host, c.Port, c.Database)
	//db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}

	return db
}
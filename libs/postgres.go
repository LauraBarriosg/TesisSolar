package libs

import (
	"fmt"
	"log"
	//"os"
	"time"

	//"gorm.io/driver/postgres"
	"github.com/jinzhu/gorm"
)



func (c *DbConfig) InitPostgresDB() *gorm.DB {
    connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ",
    c.Host, c.User, c.Password, c.Database, c.Port)
	
	db, err := gorm.Open("postgres", connString)
	for err != nil {
		log.Panic("No se puede conectar a postgres", err)
		time.Sleep(time.Second*10)
		db, err = gorm.Open("postgres", connString)
	}

	return db
}




/*func (c *DbConfig) InitPostgresDB() *gorm.DB {
	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ", 
	c.Host, c.User, c.Password, c.Database, c.Port)
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}

	return db




	conn, err := sql.Open(c.User, c.Password, c.Host, c.Port, c.Database)
    if err != nil {
        log.Fatal("cannot connect to db:", err)
    }

    store := db.NewStore(conn)
    server := api.NewServer(store)

    err = server.Start(config.ServerAddress)
    if err != nil {
        log.Fatal("cannot start server:", err)
	}
}*/
package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBConfig struct {
	DBName, DbHost, DbPassword, DbUser, DbPort string
}

type DBConnectionString struct {
	Dialect string
	Dsn     string
}

var DB *gorm.DB

func ConnectDatabase(config DBConnectionString) {

	database, err := gorm.Open(config.Dialect, config.Dsn)

	if err != nil {
		log.Fatal(err)
		panic("Failed to connect to DB")
	}

	database.AutoMigrate(&Book{}, &Author{})

	DB = database

}

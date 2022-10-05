package database

import (
	"fga/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	host     = "fga-db"
	port     = 5432
	user     = "user1"
	password = "password"
	dbname   = "fga"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("cant open database ", err)
	}

	db.Debug().AutoMigrate(&model.Person{})
}

func GetDB() *gorm.DB {
	return db
}

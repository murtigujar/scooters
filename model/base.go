package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"fmt"
	"log"
	"os"
)

var db *gorm.DB

// Rental rate: $1 / mile
var Rate = 1.0

func init() {

	e := godotenv.Load()
	if e != nil {
		log.Fatal(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")


	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		log.Fatal(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Scooter{}, &Reservation{}, &Payment{})
}


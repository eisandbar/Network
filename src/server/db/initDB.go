package db

import (
	"fmt"
	"log"
	"os"
	"time"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
)

var (
	HOST = os.Getenv("POSTGRES_HOST")
	USER = os.Getenv("POSTGRES_USER")
	DB_NAME = os.Getenv("POSTGRES_DB")
	DB_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
)

func InitDB () {
	connStr := "host=db port=5432 user=pguser dbname=pgdb password=secret sslmode=disable"
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to postgres: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})
  	db.AutoMigrate(&Message{})

	fmt.Println("Postgres DB initialized")
}

type User struct {
	Username string
	Email string
	Password string
	Id int64
	ProfileId int64
}

type Message struct {
	Text string
	SenderId int64
	Date time.Time
}
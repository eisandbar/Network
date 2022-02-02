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
	LB_HOST = os.Getenv("LB_HOST")
	HOST = os.Getenv("POSTGRES_HOST")
	USER = os.Getenv("POSTGRES_USER")
	DB_NAME = os.Getenv("POSTGRES_DB")
	DB_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
)

var WriteConnStr = fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable", 
	HOST,
	USER,
	DB_NAME,
	DB_PASSWORD,
)

var ReadConnStr = fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable", 
LB_HOST,
USER,
DB_NAME,
DB_PASSWORD,
)

func InitDB () {
	db, err := gorm.Open("postgres", WriteConnStr)
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
	Id int64 `gorm:"primaryKey"`
	ProfileId int64
}

type Message struct {
	Id int64 `gorm:"primaryKey"`
	Text string
	SenderId int64
	Date time.Time
}
package endpoints

import (
	"fmt"
	"log"
	"os"
	"io"
	"time"
	"encoding/json"
	"net/http"
	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"context"
	"github.com/go-redis/redis/v8"
	dbp "polarion/network/src/server/db"
)

var (
	HOST = os.Getenv("POSTGRES_HOST")
	USER = os.Getenv("POSTGRES_USER")
	DB_NAME = os.Getenv("POSTGRES_DB")
	DB_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
)

var connStr = fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable", 
	HOST,
	USER,
	DB_NAME,
	DB_PASSWORD,
)

var ropt = redis.Options{
	Addr: "redis:6379",
	Password: "",
	DB: 0,
}
var ctx = context.Background()

func UserPost (w http.ResponseWriter, r *http.Request) {
	fmt.Println("User signup request")

	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to postgres: %v", err)
	}
	defer db.Close()
  
	var user dbp.User

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	db.Create(&user)

	io.WriteString(w, "User created\n")
	fmt.Println("User created", user)
}

func UserDel (w http.ResponseWriter, r *http.Request) {
	fmt.Println("User termination request")
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to postgres: %v", err)
	}
	defer db.Close()

	params := mux.Vars(r)
  
	db.Delete(&dbp.User{}, params["id"])
	
	io.WriteString(w, "User deleted\n")
	fmt.Println("User deleted", params["id"])
}

func UserGet (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Finding user")
  
	var user dbp.User
	params := mux.Vars(r)
	
	// check redis cache
	rdb := redis.NewClient(&ropt)
	val, err := rdb.Get(ctx, "id" + params["id"]).Result()
	if err == redis.Nil {  // Key not found

		// Search postgres db
		db, err := gorm.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("Error connecting to postgres: %v", err)
		}
		defer db.Close()
  
		db.First(&user, params["id"])
		json, err := json.Marshal(user)
		if err != nil {
			log.Fatalf("Error converting to json")
		}

		// Add to rdb
		err = rdb.Set(ctx, "id" + params["id"], json, 120*time.Second).Err()
		if err != nil {
			log.Fatalf("Error saving to cache: %v", err)
		}

		// sending response
		fmt.Fprintf(w, string(json))

	} else if err != nil {
		log.Fatalf("Error getting from cache: %v", err)

	} else {
		fmt.Fprintf(w, val)
	}
}

func MessageGet (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Finding user messages")

	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to postgres: %v", err)
	}
	defer db.Close()

	params := mux.Vars(r)
	type Result struct {
		Username string
		Text string
		Date time.Time
	}
	var results []Result
	db.Model(&dbp.User{}).
	Select("users.Username, messages.Text, messages.Date").
	Joins("right join messages on messages.Sender_Id = users.Id").
	Where("users.Id = ?", params["id"]).
	Scan(&results)

	
	json.NewEncoder(w).Encode(&results)
	fmt.Println("Found messages for user", results)
}
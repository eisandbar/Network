package endpoints

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	dbp "polarion/network/src/server/db"
)

func PGQuery (w http.ResponseWriter, r *http.Request) {
	connStr := "host=db port=5432 user=pguser dbname=pgdb password=secret sslmode=disable"
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to postgres: %v", err)
	}
	defer db.Close()

	params := mux.Vars(r)
  
	var user dbp.User
  
	db.First(&user, params["id"])
	
	json.NewEncoder(w).Encode(&user)
	fmt.Println("Received db query")
}

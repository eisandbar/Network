package main

import (
    "fmt"
    "net/http"
    "io"
    "log"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
    sep "polarion/network/src/server/endpoints"
    db "polarion/network/src/server/db"
)

const (
    CONN_HOST = "0.0.0.0"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

func main() {
    // init db
    db.InitDB()

    router := mux.NewRouter()

    router.HandleFunc("/ping", sep.Ping).Methods("GET")
    router.HandleFunc("/longQ", sep.LongQuery).Methods("GET")
    router.HandleFunc("/loadQ/{num}", sep.LoadQueryRPC).Methods("GET")

    // RESTful
    router.HandleFunc("/users", sep.UserPost).Methods("POST") // signup
    router.HandleFunc("/users/{id}", sep.UserDel).Methods("DELETE") // acc termination
    router.HandleFunc("/users/{id}", sep.UserGet).Methods("GET") // get user data
    router.HandleFunc("/users/{id}/messages", sep.MessageGet).Methods("GET") // get user messages
    
    handler := cors.Default().Handler(router)
    
    fmt.Println("Listening on port:", CONN_PORT)

    log.Fatal(http.ListenAndServe(":" + CONN_PORT, handler))
}

// Handles incoming requests.
func simpleHandler(w http.ResponseWriter, _ *http.Request) {
    io.WriteString(w, "Hello from a HandleFunc\n")
    fmt.Println("Received message")
}
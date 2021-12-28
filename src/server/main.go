package main

import (
    "fmt"
    "net/http"
    "io"
    "log"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
    sep "polarion/network/src/server/endpoints"
)

const (
    CONN_HOST = "0.0.0.0"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/ping", sep.Ping).Methods("GET")
    router.HandleFunc("/longQ", sep.LongQuery).Methods("GET")
    router.HandleFunc("/pgQ/{id}", sep.PGQuery).Methods("GET")
    
    handler := cors.Default().Handler(router)
    
    fmt.Println("Listening on port:", CONN_PORT)

    log.Fatal(http.ListenAndServe(":" + CONN_PORT, handler))
}

// Handles incoming requests.
func simpleHandler(w http.ResponseWriter, _ *http.Request) {
    io.WriteString(w, "Hello from a HandleFunc\n")
    fmt.Println("Received message")
}
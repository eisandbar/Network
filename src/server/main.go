package main

import (
    "fmt"
    "net/http"
    "io"
    "log"
    sep "polarion/network/src/server/endpoints"
)

const (
    CONN_HOST = "server"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

func main() {
    http.HandleFunc("/ping", sep.Ping)
    http.HandleFunc("/longQ", sep.LongQuery)
    
    fmt.Println("Listening on port:", CONN_PORT)
    log.Fatal(http.ListenAndServe(CONN_HOST + ":" + CONN_PORT, nil))
}

// Handles incoming requests.
func simpleHandler(w http.ResponseWriter, _ *http.Request) {
    io.WriteString(w, "Hello from a HandleFunc\n")
    fmt.Println("Received message")
}
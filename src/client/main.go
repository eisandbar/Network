package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"io"
	"math/rand"
)

const (
	CONN_HOST = "api"
	CONN_PORT = "5100"
)

func main() {
	for {
		time.Sleep(2 * time.Second)
		if rand.Float64() > 0.2 {
			go simpleGet("/ping")
		} else {
			go simpleGet("/longQ")
		}
	}
}

func simpleGet(ep string) {
	// Get request
	resp, err := http.Get("http://" + CONN_HOST + ":" + CONN_PORT + ep)
	if err != nil {
		log.Fatalf("Error Connecting: %v", err)
	}
	// Close body when the application closes.
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}

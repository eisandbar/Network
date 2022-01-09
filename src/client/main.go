package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"io"
	"math/rand"
	conn "polarion/network/src/client/util"
	req "polarion/network/src/client/requests"
)

func main() {
	for {
		time.Sleep(200 * time.Millisecond)

		// Simple queries
		if rand.Float64() > 0.2 {
			go simpleGet("/ping")
		}

		if rand.Float64() > 0.5 {
			go simpleGet("/longQ")
		}

		go req.UserPost()
		for i:=0; i<10; i++ {
			go req.UserGet(rand.Intn(10000))
		}
		for i:=0; i<30; i++ {
			go req.UserGet(rand.Intn(400)) // super users
		}
	}
}

func simpleGet(ep string) {
	// Get request
	resp, err := http.Get(conn.ConnStr + ep)
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

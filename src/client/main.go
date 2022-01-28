package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"io"
	"strconv"
	"math/rand"
	conn "polarion/network/src/client/util"
	req "polarion/network/src/client/requests"
)

func main() {
	for true {
		time.Sleep(400 * time.Millisecond)

		// Simple queries
		go simpleGet("/ping")
		go simpleGet("/longQ")

		// // MQ queries
		num := rand.Intn(7) + 31
		go simpleGet("/loadQ/" + strconv.Itoa(num))

		// REST API queries
		go req.UserPost()
		go req.UserGet(rand.Intn(10000))
		go req.UserGet(rand.Intn(10000))
		for i:=0; i<4; i++ {
			go req.UserGet(rand.Intn(400)) // super users
		}
	}
}

func simpleGet(ep string) {
	// Get request
	resp, err := http.Get(conn.ConnStr + ep)
	if err != nil {
		fmt.Printf("Error Connecting: %v", err)
		return
	}
	// Close body when the application closes.
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}

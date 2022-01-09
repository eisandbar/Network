package requests

import (
	"fmt"
	"log"
	"net/http"
	"io"
	"bytes"
	"strconv"
	"encoding/json"
	"math/rand"
	conn "polarion/network/src/client/util"
)

type User struct {
	Username string
	Email string
	Password string
	Id int64
	ProfileId int64
}

func UserPost () {
	newUser := Constructor()
	b, err := json.Marshal(newUser)
	// Post request
	resp, err := http.Post(conn.ConnStr + "/users", "application/json", bytes.NewBuffer(b))
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

func UserGet (id int) {
	// Post request
	resp, err := http.Get(conn.ConnStr + "/users/" + strconv.Itoa(id))
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

func Constructor () User {
	var u User
	l := 5 + rand.Intn(12)
	username := make([]byte, l)
	for i:=0; i<l; i++ {
		username[i] = byte(rand.Intn(26)) + byte('a')
	}
	u.Username = string(username)
	u.Email = u.Username + "@mail.com"
	u.Password = "secretPass"
	return u
}
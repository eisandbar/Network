package endpoints

import (
	"fmt"
	"io"
	"net/http"
)

func Ping (w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Pinged server\n")
	fmt.Println("Received message")
}
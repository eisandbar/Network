package endpoints

import (
	"fmt"
	"io"
	"net/http"
)

func Ping (w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello from a HandleFunc\n")
	fmt.Println("Received message")
}
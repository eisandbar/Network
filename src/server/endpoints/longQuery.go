package endpoints

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func LongQuery (w http.ResponseWriter, _ *http.Request) {
	time.Sleep(20 * time.Second)
	io.WriteString(w, "This query takes 20s\n")
	fmt.Println("Received long query")
}
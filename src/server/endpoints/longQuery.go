package endpoints

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func LongQuery (w http.ResponseWriter, _ *http.Request) {
	
	fmt.Println("Received long query")
	time.Sleep(20 * time.Second)
	io.WriteString(w, "This query takes 20s\n")
	fmt.Println("Waited 20s")
}
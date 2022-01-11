package endpoints

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func LoadQuery (w http.ResponseWriter, _ *http.Request) {
	
	fmt.Println("Received a load heavy query")
	res := 1
	for i:=1; i<(1<<20); i++ {
		for j:=1; j<(1<<30); j++ {
			res = ((res * (i % 133711)) + j) % 133711
		}
	}
	io.WriteString(w, strconv.Itoa(res))
	fmt.Println("Result of arduous calculations", res)
}
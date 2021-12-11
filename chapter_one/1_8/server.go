// 1.12: modify the server to read parameter values from the url
package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("cycles") {
		Lissajous(w, 5)
	}

	cycles, err := strconv.Atoi(r.URL.Query().Get("cycles"))

	if err != nil {
		panic(err)
	}

	Lissajous(w, cycles)
}

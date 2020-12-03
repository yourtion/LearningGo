// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler1) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler1 echoes the Path component of the request URL r.
func handler1(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

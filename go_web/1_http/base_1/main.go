package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Set router handler function
	http.HandlerFunc("/", indexHandler)
	http.HandlerFunc("/hello", helloHandler)

	// Start HTTP server and listen port 9999
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

package main

import (
	"api"
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/test", testserve)
	http.HandleFunc("/helloworld/read", api.ReadHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func testserve(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

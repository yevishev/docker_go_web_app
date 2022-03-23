package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["name"]

	name := "guest"

	if ok {

		name = keys[0]
	}

	fmt.Fprintf(w, "Hello %s!\n", name)
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Your solution goes here. Good luck!
	http.HandleFunc("/hello", helloFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, "Hello, ", name)
}

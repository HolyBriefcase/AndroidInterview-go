package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func hello2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World222")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/hello", hello2)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}

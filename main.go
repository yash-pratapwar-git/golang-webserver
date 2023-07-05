package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World!!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there !!")
	})

	http.HandleFunc("/second", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there second !!")
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}

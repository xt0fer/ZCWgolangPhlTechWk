package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain") // optional
		w.WriteHeader(http.StatusOK)                 // optional
		fmt.Fprintln(w, "Hello, 世界")
	}

	helloHandler = logg(helloHandler) // wraps the orig handler

	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func logg(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before")
		h.ServeHTTP(w, r) // call original handler
		log.Println("After")
	})
}

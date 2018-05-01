package main

import (
    "fmt"
    "net/http"
)

func main() {
helloHandler := func(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", “text/plain”) // optional
w.WriteHeader(http.StatusOK) // optional
fmt.Fprintln(w, "Hello, 世界”)
        }

   http.HandleFunc(“/", helloHandler)
   http.ListenAndServe(":8080", nil)
}

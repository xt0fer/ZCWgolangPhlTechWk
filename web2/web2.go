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

helloHandler = log(helloHandler) // wraps the orig handler

   http.HandleFunc(“/", helloHandler)
   http.ListenAndServe(":8080", nil)
}

func log(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)
  {
    log.Println("Before")
    h.ServeHTTP(w, r) // call original handler
    log.Println("After")
  })
}

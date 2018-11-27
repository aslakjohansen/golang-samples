package main

import (
    "fmt"
    "net/http"
)

func index_handler (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "You have reached the root!")
}

func about_handler (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "About what?")
}

func main () {
    http.HandleFunc("/", index_handler)
    http.HandleFunc("/about", about_handler)
    http.ListenAndServe(":8000", nil)
}


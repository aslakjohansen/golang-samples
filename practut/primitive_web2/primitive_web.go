package main

import (
    "fmt"
    "net/http"
)

func index_handler (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Some %s</h1>", "header")
    
    fmt.Fprintf(w,`<p>line 1</p>
    <p>line 2</p>
    <p>line 3</p>`)
}

func main () {
    http.HandleFunc("/", index_handler)
    http.ListenAndServe(":8000", nil)
}

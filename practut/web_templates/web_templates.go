package main

import (
    "fmt"
    "net/http"
    "html/template"
)

type NewsAggPage struct {
    Title string
    News string
}

func index_handler (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Some header</h1>")
}

func newsagg_handler (w http.ResponseWriter, r *http.Request) {
    p := NewsAggPage {Title: "Amazing News Aggregator", News: "Some news"}
    t, _ := template.ParseFiles("basictemplating.html")
    t.Execute(w, p)
}

func main () {
    http.HandleFunc("/", index_handler)
    http.HandleFunc("/agg", newsagg_handler)
    http.ListenAndServe(":8000", nil)
}


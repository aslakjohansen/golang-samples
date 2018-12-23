package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main () {
    resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
    bytes, _ := ioutil.ReadAll(resp.Body)
    contents := string(bytes)
    resp.Body.Close()
    
    fmt.Println(contents)
}


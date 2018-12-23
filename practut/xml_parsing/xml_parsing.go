package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/xml"
)

//<?xml version="1.0" encoding="UTF-8"?>
//<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
//   </sitemap>
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
//   </sitemap>
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
//   </sitemap>
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-blogs-opinions-sitemap.xml</loc>
//   </sitemap>
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-local-sitemap.xml</loc>
//   </sitemap>
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-blogs-local-sitemap.xml</loc>
//   </sitemap>   
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-sports-sitemap.xml</loc>
//   </sitemap>
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-blogs-sports-sitemap.xml</loc>
//   </sitemap>   
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-national-sitemap.xml</loc>
//   </sitemap>
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-blogs-national-sitemap.xml</loc>
//   </sitemap>   
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-world-sitemap.xml</loc>
//   </sitemap>
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-blogs-world-sitemap.xml</loc>
//   </sitemap>   
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-business-sitemap.xml</loc>
//   </sitemap>
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-blogs-business-sitemap.xml</loc>
//   </sitemap>   
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-technology-sitemap.xml</loc>
//   </sitemap>
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-blogs-technology-sitemap.xml</loc>
//   </sitemap>   
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-lifestyle-sitemap.xml</loc>
//   </sitemap>
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-blogs-lifestyle-sitemap.xml</loc>
//   </sitemap>   
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-entertainment-sitemap.xml</loc>
//   </sitemap>
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-blogs-entertainment-sitemap.xml</loc>
//   </sitemap>  
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-blogs-goingoutguide-sitemap.xml</loc>
//   </sitemap>   
//   <sitemap>
//      <loc>http://www.washingtonpost.com/news-goingoutguide-sitemap.xml</loc>
//   </sitemap> 
//</sitemapindex>

type SitemapIndex struct {
    Locations []Location `xml:"sitemap"`
}

type Location struct {
    Loc string `xml:"loc"`
}

func (l Location) String () string {
    return fmt.Sprintf(l.Loc)
}

func main () {
    resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
    bytes, _ := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    
    var s SitemapIndex
    xml.Unmarshal(bytes, &s)
    fmt.Println(s.Locations)
}


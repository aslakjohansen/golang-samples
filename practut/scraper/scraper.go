package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/xml"
)

var washPostXML = []byte(`
<?xml version="1.0" encoding="UTF-8"?>
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
   <sitemap>
      <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-opinions-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-local-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-local-sitemap.xml</loc>
   </sitemap>   
   <sitemap>
      <loc>http://www.washingtonpost.com/news-sports-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-sports-sitemap.xml</loc>
   </sitemap>   
   <sitemap>
      <loc>http://www.washingtonpost.com/news-national-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-national-sitemap.xml</loc>
   </sitemap>   
   <sitemap>
      <loc>http://www.washingtonpost.com/news-world-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-world-sitemap.xml</loc>
   </sitemap>   
   <sitemap>
      <loc>http://www.washingtonpost.com/news-business-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-business-sitemap.xml</loc>
   </sitemap>   
   <sitemap>
      <loc>http://www.washingtonpost.com/news-technology-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-technology-sitemap.xml</loc>
   </sitemap>   
   <sitemap>
      <loc>http://www.washingtonpost.com/news-lifestyle-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-lifestyle-sitemap.xml</loc>
   </sitemap>   
   <sitemap>
      <loc>http://www.washingtonpost.com/news-entertainment-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-entertainment-sitemap.xml</loc>
   </sitemap>  
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-goingoutguide-sitemap.xml</loc>
   </sitemap>   
   <sitemap>
      <loc>http://www.washingtonpost.com/news-goingoutguide-sitemap.xml</loc>
   </sitemap> 
</sitemapindex>
`)

type SitemapIndex struct {
    Locations []string `xml:"sitemap>loc"`
}

type News struct {
    Titles    []string `xml:"url>news>title"`
    Keywords  []string `xml:"url>news>Keywords"`
    Locations []string `xml:"url>loc"`
}

func main () {
    var s SitemapIndex
    var n News
    
//    resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
//    bytes, _ := ioutil.ReadAll(resp.Body)
//    resp.Body.Close()
    bytes := washPostXML
    fmt.Println(string(bytes))
    xml.Unmarshal(bytes, &s)
    fmt.Println(s)
    
    for _, Location := range s.Locations {
        resp, _ := http.Get(Location)
        bytes, _ := ioutil.ReadAll(resp.Body)
        resp.Body.Close()
        xml.Unmarshal(bytes, &n)
        fmt.Println(n)
    }
}

package main

import (
    "fmt"
)

func main () {
//    var grades map[string]float32
//    grades := map[string]float32
    grades := make(map[string]float32)
    
    // insert individual mappings
    grades["Timmy"] = 42
    grades["Jess"]  = 92
    grades["Sam"]   = 67
    
    // contents
    fmt.Println(grades)
    
    // lookup
    TimsGrade := grades["Timmy"]
    fmt.Println(TimsGrade)
    
    // delete element
    delete(grades, "Timmy")
    
    // contents
    fmt.Println(grades)
    
    for k, v := range grades {
        fmt.Println(k, ": ", v)
    }
    
    // check whether key is in map
    _, exists := grades["Jenn"]
    fmt.Println("Jenn in map:", exists)
    if _, exists := grades["Jess"]; exists {
        fmt.Println("Jess is in map")
    }
    
}

package main

import (
    "fmt"
)

func main () {
    var a [3]int   // 1d array
    var b [2][2]int // 2d array
    var c []int // slice
    
    c = a[1:2]
    
    a[0] = 1
    b[0][1] = 2
    c[0] = 3
    
    fmt.Println(a[0], b[0][1], c[0])
    fmt.Println("End", c)
}


package main

import "fmt"

func add (x float64, y float64) float64 {
    return x+y
}

func main () {
    const num1 float64 = 5.6
    var   num2 float64 = 9.5
    var   a    int     = 12
    var   b    float64 = float64(a)
    
    fmt.Println(add(num1, num2)+b)
}


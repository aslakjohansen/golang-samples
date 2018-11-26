package main

import "fmt"

func add (x float64, y float64) float64 {
    return x+y
}

func main () {
    var num1 float64 = 5.6
    var num2 float64 = 9.5
    
    fmt.Println(add(num1, num2))
}


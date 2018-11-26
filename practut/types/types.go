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
    
    x := b // infers the type of x from b
    
    fmt.Println(add(num1, num2)+b+x)
}


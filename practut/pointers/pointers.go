package main

import "fmt"

func main () {
    var a int = 12
    b := &a // address in memory of a
    
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(*b)
    
    *b = 5
    fmt.Println(" a =", a)
    
    *b = *b**b
    fmt.Println("*b =", *b)
    
    var c *int = b
    fmt.Println("*c =", *c)
}


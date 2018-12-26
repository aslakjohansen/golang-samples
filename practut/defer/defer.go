package main

import (
    "fmt"
)

func foo() {
    defer fmt.Println("Done!")
    defer fmt.Println("Are we done?")
    
    for i:=0 ; i<5 ; i++ {
        defer fmt.Println(i)
    }
    
    fmt.Println("Doing some stuff ..")
}

func main () {
    foo()
}


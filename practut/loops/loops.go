package main

import (
    "fmt"
)

func main () {
    // for
    for i:=0 ; i<10 ; i++ {
        fmt.Println(i)
    }
    
    // while
    j := 0
    for j<10 {
        fmt.Println(j)
        j += 3
    }
    
    // do (almost) forever
    k := 1
    for {
        fmt.Println(k, ": Busy!")
        k += 4
        if k>13 {
            break
        }
    }
}


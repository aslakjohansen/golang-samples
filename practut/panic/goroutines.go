package main

import (
    "fmt"
    "time"
    "sync"
)

var wg sync.WaitGroup // wait group

func cleanup () {
    defer wg.Done()
    if r:=recover() ; r!=nil {
        fmt.Println("Recovered in cleanup:", r)
    }
}

func say (s string) {
    defer cleanup()
    
    for i:=0 ; i<3 ; i++ {
        if i==2 {
            panic("Oh dear, a 2")
        }
        fmt.Println(s)
        time.Sleep(time.Millisecond*100)
    }
}

func main () {
    wg.Add(1)
    go say("Hey")
    
    wg.Add(1)
    go say("There")
    
    wg.Add(1)
    go say("Hi")
    
    wg.Wait()
}


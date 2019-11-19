package main

import (
    "fmt"
    "time"
)

func producer (channel chan int, value int, sleeptime int) {
    time.Sleep(time.Duration(sleeptime)*time.Second)
    channel <- value
}

func main () {
    for i:=0 ; i<10 ; i++ {
        channel := make(chan int)
        go producer(channel, i, i%5)
        
        select {
        case v := <- channel:
            fmt.Println(v)
        case <- time.After(3*time.Second):
            fmt.Println("timeout")
        }
    }
}


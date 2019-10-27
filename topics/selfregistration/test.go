package main

import (
    "fmt"
    
    "golang-samples/topics/selfregistration/modules"
    "golang-samples/topics/selfregistration/modules/a"
)

func main () {
    fmt.Println("Done loading ...")
    modules.Print()
    a.Print()
}

package a

import (
    "fmt"
    
    "golang-samples/topics/selfregistration/modules"
)

func init () {
    fmt.Println("Registering module ...")
    var entry modules.Module = modules.Module{Name: "a"}
    modules.Register(entry)
}

func Print () {
    fmt.Println("Done(a)")
}

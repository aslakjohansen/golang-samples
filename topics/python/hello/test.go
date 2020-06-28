package main

import (
    "runtime"
    "fmt"
    python "github.com/sbinet/go-python"
)

func main () {
    runtime.LockOSThread() // stick go routine to thread
    
    err := python.Initialize()
    if err != nil {
        fmt.Println("Unable to initialize python", err)
    }
    
    module := python.PyImport_ImportModule("module")
    if module == nil {
        fmt.Println("Unable to import module")
    }
    
    hello := module.GetAttrString("hello")
    if module == nil {
        fmt.Println("Unable to name function in module")
    }
    
    state := python.PyEval_SaveThread()
    gstate := python.PyGILState_Ensure()
    
    hello.Call(python.PyTuple_New(0), python.PyDict_New())
    
    python.PyGILState_Release(gstate)
    python.PyEval_RestoreThread(state)
    
    err = python.Finalize()
    if err != nil {
        fmt.Println("Unable to finalize python", err)
    }
}


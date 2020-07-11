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
    
    fun := module.GetAttrString("add")
    if fun == nil {
        fmt.Println("Unable to name function in module")
    }
    
    state := python.PyEval_SaveThread()
    gstate := python.PyGILState_Ensure()
    
    // construct arguments
    args := python.PyTuple_New(2)
    var a1 float64 = 42.5
    var a2 float64 = 13.5
    python.PyTuple_SET_ITEM(args, 0, python.PyFloat_FromDouble(a1))
    python.PyTuple_SET_ITEM(args, 1, python.PyFloat_FromDouble(a2))
    
    resPython := fun.Call(args, python.PyDict_New())
    
    // decode result
    res := python.PyFloat_AsDouble(resPython)
    if python.PyErr_Occurred()!=nil {
        fmt.Println("Decoding of python return value failed")
    }
    
    fmt.Println("Result:", res)
    
    python.PyGILState_Release(gstate)
    python.PyEval_RestoreThread(state)
    
    err = python.Finalize()
    if err != nil {
        fmt.Println("Unable to finalize python", err)
    }
}


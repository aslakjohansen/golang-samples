package main

import (
    "runtime"
    "fmt"
    python "github.com/sbinet/go-python"
)

var (
    fun *python.PyObject
)

func call (par float64) float64 {
    state := python.PyEval_SaveThread()
    gstate := python.PyGILState_Ensure()
    
    // construct arguments
    args := python.PyTuple_New(1)
    python.PyTuple_SET_ITEM(args, 0, python.PyFloat_FromDouble(par))
    
    resPython := fun.Call(args, python.PyDict_New())
    
    // decode result
    res := python.PyFloat_AsDouble(resPython)
    if python.PyErr_Occurred()!=nil {
        fmt.Println("Decoding of python return value failed")
    }
    
    python.PyGILState_Release(gstate)
    python.PyEval_RestoreThread(state)
    
    return res;
}

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
    
    fun = module.GetAttrString("add")
    if fun == nil {
        fmt.Println("Unable to name function in module")
    }
    
    r1 := call(42.5)
    fmt.Println("res: ", r1);
    fmt.Println("res: ", call(13.5));
    fmt.Println("res: ", call(6.35));
    
    err = python.Finalize()
    if err != nil {
        fmt.Println("Unable to finalize python", err)
    }
}


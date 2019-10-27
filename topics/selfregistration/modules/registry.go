package modules

import (
    "fmt"
)

type Module struct {
    Name string
}

var Entries []Module = make([]Module, 0, 0)

func Register (module Module) {
    fmt.Println("Registering '", module.Name, "'");
    Entries = append(Entries, module)
}

func Print () {
    fmt.Println("Modules loaded:");
    for _, entry := range Entries {
        fmt.Println(" -", entry.Name);
    }
}

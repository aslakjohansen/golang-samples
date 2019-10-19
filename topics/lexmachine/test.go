package main

import (
    "fmt"
    "os"
    
    "golang-samples/topics/lexmachine/lang"
)

func main () {
    var lexer = lang.NewLexer(true)
    
    var input string = "3.14+7"
    
    line, err := lang.Parse(lexer, input)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    
    fmt.Println(line)
}

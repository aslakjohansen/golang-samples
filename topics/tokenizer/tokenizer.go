package main

import "fmt"

type TokenClass int
const (
    NUM TokenClass = iota
    VAR
    PLUS
    MINUS
    MUL
    DIV
    POW
    LPAR
    RPAR
)

type token struct {
    class      TokenClass
    begin_line int
    begin_col  int
    end_line   int
    end_col    int
}

func tokenize (input string) {
//    var start int = 0
    var begin_line  int = 1
    var begin_col   int = 1
    var end_line  int = 1
    var end_col   int = 1
    
    for i := range input {
        fmt.Printf("input[%d] = '%d'\n", i, input[i])
        
        begin_line = end_line
        begin_col  = end_col
    }
}

func main () {
    fmt.Println("Tokenizer")
    
    var input string = "1.2+ (13-\npi/2)*e ^2";
    
    fmt.Println("input: "+input)
    tokenize(input)
}

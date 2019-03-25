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
type token_constructor func(c chan int, begin_line int, begin_col int, end_line int, end_col int)

type state struct {
    constructor token_constructor
}

type tokenizer struct {
    states []state
    fill int
}

func (t *tokenizer) init () {
    t.states = make([]state, 0, 0)
    t.fill = 0
}

type token struct {
    class      TokenClass
    begin_line int
    begin_col  int
    end_line   int
    end_col    int
}

func (t *tokenizer) tokenize (input string) {
//    var start int = 0
    var begin_line  int = 1
    var begin_col   int = 1
    var end_line    int = 1
    var end_col     int = 1
    
    for i := range input {
        fmt.Printf("input[%d] = '%d'\n", i, input[i])
        
        // temp fix to keep golang happy
        if true {
            begin_line = begin_col
            begin_col = begin_line
        }
        
        begin_line = end_line
        begin_col  = end_col
    }
}

func main () {
    fmt.Println("Tokenizer")
    
    var t tokenizer = tokenizer{}
    t.init()
    
    var input string = "1.2+ (13-\npi/2)*e ^2";
    
    fmt.Println("input: "+input)
    t.tokenize(input)
}

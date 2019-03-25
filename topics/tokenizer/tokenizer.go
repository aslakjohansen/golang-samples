package main

import "fmt"

type TokenClass int
const (
    INVALID TokenClass = iota
    NUM
    VAR
    PLUS
    MINUS
    MUL
    DIV
    POW
    LPAR
    RPAR
    EOF
)
type token_constructor func(tc TokenClass, c chan int, begin_line int, begin_col int, end_line int, end_col int)

type state struct {
    constructor token_constructor
    tokenclass  TokenClass // TODO: prettyfy
    next [256]int
}

type tokenizer struct {
    states []state
    fill int
}

func (t *tokenizer) init () {
    t.states = make([]state, 0, 0)
    t.fill = 0
}
func (t *tokenizer) populate_static (key string, constructor token_constructor, tc TokenClass) {
    var s int = 0
    
    // find correct state
    for i := range key {
        var char byte = key[i]
        
        // guard: we are starting a new branch
        if t.states[s].next[char] == 0 {
            t.states = append(t.states, state{nil, INVALID, [256]int{0}})
            t.states[s].next[char] = len(t.states)-1
        }
        
        s = t.states[s].next[char]
    }
    
    t.states[s].constructor = constructor
    t.states[s].tokenclass = tc
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

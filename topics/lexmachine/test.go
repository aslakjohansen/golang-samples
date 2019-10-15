package main

import (
    "fmt"
    
    "github.com/timtadh/lexmachine"
    "github.com/timtadh/lexmachine/machines"
)

var NUMBER int = 1
var PLUS   int = 2
var SPACE  int = 3

func token (id int) lexmachine.Action {
    return func (s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
        return s.Token(id, string(m.Bytes), m), nil
    }
}

func newLexer (dfa bool) *lexmachine.Lexer {
    var lexer = lexmachine.NewLexer()
    lexer.Add([]byte(`\+`), token(PLUS))
    lexer.Add([]byte(`[0-9]*\.?[0-9]+`), token(NUMBER))
    lexer.Add([]byte(`\s+`), token(SPACE))
    
    var err error
    if dfa {
        err = lexer.CompileDFA()
    } else {
        err = lexer.CompileNFA()
    }
    if err != nil {
        panic(err)
    }
    return lexer
}

func scan (lexer *lexmachine.Lexer, input string) error {
    scanner, err := lexer.Scanner([]byte(input))
    if err != nil {
        panic(err)
    }
    
    for tok, err, eos := scanner.Next(); !eos; tok, err, eos = scanner.Next() {
        if ui, is := err.(*machines.UnconsumedInput); is {
            // skip the error via:
            // scanner.TC = ui.FailTC
            //
            fmt.Println("scan: ui=", ui)
            fmt.Println("scan: is=", is)
            return err
        } else if err != nil {
            return err
        }
        fmt.Println(tok)
    }
    return nil
}

func main () {
    var lexer = newLexer(true)
    
    var input string = "3.14 + 7"
    var err = scan(lexer, input)
    fmt.Println("main: error=", err)
    
//    fmt.Println("Salary:", salary)
//    fmt.Printf("Sibling[%d] = %s\n", i, s.(string))
}

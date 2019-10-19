package lang

import (
    "fmt"
    
    "github.com/timtadh/lexmachine"
    "github.com/timtadh/lexmachine/machines"
)

//var NUMBER int = 1
//var PLUS   int = 2
//var SPACE  int = 3

type golex struct {
    *lexmachine.Scanner
    line *Node
}

func newGoLex (lexer *lexmachine.Lexer, text []byte) (*golex, error) {
    scan, err := lexer.Scanner(text)
    if err != nil {
        return nil, err
    }
    return &golex{Scanner: scan}, nil
}

func (g *golex) Lex (lval *yySymType) (tokenType int) {
    s := g.Scanner
    tok, err, eof := s.Next()
    if err != nil {
        g.Error(err.Error())
    } else if eof {
        return -1 // Note: signals EOF to yyParse
    }
    
    lval.token = tok.(*lexmachine.Token)
    return lval.token.Type
}

func token (id int) lexmachine.Action {
    return func (s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
        return s.Token(id, string(m.Bytes), m), nil
    }
}

func (l *golex) Error (message string) {
    panic(fmt.Errorf(message))
}

// TODO: is this dead code?
func scan (lexer *lexmachine.Lexer, input string) error {
    fmt.Println("!!!!!!!!!!!!!!!!! UNDEAD CODE !!!!!!!!!!!!!!!!!")
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

func NewLexer (dfa bool) *lexmachine.Lexer {
    var lexer = lexmachine.NewLexer()
    lexer.Add([]byte(`\+`), token(PLUS))
    lexer.Add([]byte(`[0-9]*\.?[0-9]+`), token(NUMBER))
//    lexer.Add([]byte(`\s+`), token(SPACE))
    
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

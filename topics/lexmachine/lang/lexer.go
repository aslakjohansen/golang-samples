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
    fmt.Println("lexer.go:newGoLex> ", lexer, text, text[0])
    scan, err := lexer.Scanner(text)
    if err != nil {
        return nil, err
    }
    return &golex{Scanner: scan}, nil
}

func (g *golex) Lex (lval *yySymType) (tokenType int) {
    fmt.Println("lexer.go:Lex> lval=", lval)
    s := g.Scanner
    tok, err, eof := s.Next()
    fmt.Println("lexer.go:Lex> next token: ", tok);
    if err != nil {
        g.Error(err.Error())
    } else if eof {
        fmt.Println("lexer.go:Lex> EOF reached");
        return -1 // Note: signals EOF to yyParse
    }
    
    lval.token = tok.(*lexmachine.Token)
    fmt.Println("lexer.go:Lex> lval.token: ", lval.token);
    
    fmt.Println("lexer.go:Lex> return value: ", lval.token.Type + (yyPrivate-1), " or ", lval.token.Type);
//    return lval.token.Type + (yyPrivate-1)
    return lval.token.Type
}

func token (id int) lexmachine.Action {
    fmt.Println("Creating token", id)
    return func (s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
        fmt.Println("lexer.go:token/anon> ", id, string(m.Bytes), m)
        return s.Token(id, string(m.Bytes), m), nil
    }
}

func (l *golex) Error (message string) {
    fmt.Println("lexer.go:Error> '", message, "'", l.line);
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
    fmt.Println("token(NUMBER) ", token(NUMBER), " | NUMBER ", NUMBER)
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

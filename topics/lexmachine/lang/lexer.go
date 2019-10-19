package lang

import (
    "fmt"
    
    "github.com/timtadh/lexmachine"
    "github.com/timtadh/lexmachine/machines"
)

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

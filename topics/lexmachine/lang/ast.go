package lang

import (
    "fmt"
    "strings"
    
    "github.com/timtadh/lexmachine"
)

type Node struct {
    Name     string
    Token    *lexmachine.Token
    Children []*Node
}

func NewNode (name string, token *lexmachine.Token) *Node {
    return &Node{
        Name:  name,
        Token: token,
    }
}

func (n *Node) AddKid (kid *Node) *Node {
    n.Children = append(n.Children, kid)
    return n
}

func (n *Node) PrependKid (kid *Node) *Node {
    kids := append(make([]*Node, 0, cap(n.Children)+1), kid)
    n.Children = append(kids, n.Children...)
    return n
}

func (n *Node) String () string {
    parts := make([]string, 0, len(n.Children))
    parts = append(parts, n.Name)
    if n.Token != nil && string(n.Token.Lexeme) != n.Name {
        parts = append(parts, fmt.Sprintf("%q", string(n.Token.Lexeme)))
    }
    for _, k := range n.Children {
        parts = append(parts, k.String())
    }
    if len(parts) > 1 {
        return fmt.Sprintf("(%v)", strings.Join(parts, " "))
    }
    return strings.Join(parts, " ")
}


%{

package lang

import (
    "fmt"
    "github.com/timtadh/lexmachine"
)

%}

%union{
    token *lexmachine.Token
    ast   *Node
}

%token	PLUS
%token	NUMBER

%left PLUS

%% /* The grammar follows.  */

Line : Expr { 
                fmt.Println("----------------------")
                yylex.(*golex).line = $1.ast
            }
     ;

Expr : Expr PLUS Expr { $$.ast = NewNode("+", $2.token).AddKid($1.ast).AddKid($3.ast) }
       | NUMBER       { 
                        fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!")
                        fmt.Println($$.ast)
                        $$.ast = NewNode("number", $1.token)
        }
     ;

;
%%

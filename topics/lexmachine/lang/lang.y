%{

package lang

import (
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

Line : Expr { yylex.(*golex).line = $1.ast }
     ;

Expr : Expr PLUS Expr           { $$.ast = NewNode("+", $2.token).AddKid($1.ast).AddKid($3.ast) }
       | NUMBER                 { $$.ast = NewNode("number", $1.token) }
     ;

;
%%

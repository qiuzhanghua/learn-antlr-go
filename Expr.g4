grammar Expr;

prog: stat+ ;

stat: expr NEWLINE               # printExpr
    | ID '=' expr NEWLINE        # assign
    | NEWLINE                    # blank
    ;

expr: left=expr MULT right=expr  # Mult
    | left=expr DIV  right=expr  # Div
    | left=expr ADD  right=expr  # Add
    | left=expr MINUS right=expr # Minus
    | value=INT                  # int
    | id=ID                      # id
    | '(' expr ')'               # parentheses
    ;

MULT: '*' ;
DIV : '/' ;
ADD : '+' ;
MINUS: '-' ;

ID  : [a-zA-Z]+ ;
INT : [0-9]+ ;
NEWLINE: '\r'? '\n' ;
WS  : [ \t]+ -> skip ;

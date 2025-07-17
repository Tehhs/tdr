todoComment
    : HASH? TODO tagBlock? COLON message
    ;

tagBlock
    : LPAREN tag (COMMA tag)* RPAREN
    ;

tag
    : IDENTIFIER (EQUALS IDENTIFIER)?
    ;

message
    : TEXT_LINE+ // can handle multi-line messages
    ;

// Tokens
HASH        : '#' ;
TODO        : 'todo' ;
LPAREN      : '(' ;
RPAREN      : ')' ;
COMMA       : ',' ;
COLON       : ':' ;
EQUALS      : '=' ;
IDENTIFIER  : [a-zA-Z_][a-zA-Z0-9_-]* ;
TEXT_LINE   : ~[\r\n]+ -> skip ;
NEWLINE     : [\r\n]+ ;
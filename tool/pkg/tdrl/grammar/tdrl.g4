grammar tdrl;

todo: TODO COLON message EOF;

message
    : ~NL* // match any token other than a line break zero or more times
    ;



TODO: 'todo' ;
COLON: ':' ;
WS: [ \t\r\n]+ -> skip ;
NL      : '\r'? '\n' | '\r';
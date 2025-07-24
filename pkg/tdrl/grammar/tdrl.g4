grammar tdrl;

main: todoRule? ;
todoRule: TODO tagRule? COLON messageContent;
tagRule: LPARAN tagList RPARAN; 
messageContent: (WS | .)+?;
tagList: TAG_ID (',' TAG_ID)* ;

TODO: 'todo' ;
TAG_ID: [0-9a-zA-Z_]+ ;
COMMA: ',' ;
LPARAN: '(' ;
RPARAN: ')' ;
COLON: ':' ;

WS : [ \t\r\n]+ -> channel(HIDDEN) ;

//todo(tag):back to school we go

